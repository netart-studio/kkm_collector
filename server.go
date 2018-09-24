package main

import (
	
   "net/http"
   "github.com/gin-gonic/gin"
   _ "github.com/jinzhu/gorm/dialects/mysql"
)

import models "avto98/kkm_collector/models"
import database "avto98/kkm_collector/database"
import config "avto98/kkm_collector/config"


var conf = config.Get_config() 

func Init(){
	db := database.Database()
	db.AutoMigrate(&models.Check{})
}

func main() {	
	Init()	
	router := gin.Default()
	router.GET("/", startPage)
	router.LoadHTMLGlob("templates/*")
	v1 := router.Group("/api/v1/")
	{
	   v1.POST("check_kkm/", models.CreateCheck)
	   v1.GET("check_kkm/",  models.FetchAllCheck)
	   //v1.GET("check/:id", FetchSingleCheck)
	   //v1.PUT("check/:id", UpdateCheck)
	   //v1.DELETE("check/:id", DeleteCheck)
	}
	router.Run(":"+conf.Port)

 }

 
 //стартовая страница 
 func startPage(c *gin.Context)  {
	c.HTML(http.StatusOK, "index.tmpl", gin.H{
	   "title": "simple api gin",
	})
 }