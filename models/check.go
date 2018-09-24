package models

import (
   "fmt"
   "log"
   "net/http"
   "github.com/gin-gonic/gin"
   "github.com/jinzhu/gorm"
   _ "github.com/jinzhu/gorm/dialects/mysql"
)

import database "avto98/kkm_collector/database"

import utils "avto98/kkm_collector/utils"

type Check struct {
	gorm.Model
	Number       		string  `gorm:"number; not null; unique"`
	KKM 				int64  	`gorm:"kkm; not null; unique"`
	Department     		int64   `gorm:"department"`
	Nalog     			int64   `gorm:"nalog"`
	CheckType     		int64   `gorm:"CheckType"`
	CheckTotal     		float64 `gorm:"CheckTotal"`
	TotalByCasche     	float64 `gorm:"TotalByCasche"`
	TotalByElectron     float64 `gorm:"TotalByElectron"`
	CheckDate     		int64   `gorm:"check_date"`
	Sync     			bool    `gorm:"sync; default false"`

 }
 
 
type JsonCheck struct {
	ID         		string   `json:"id"`
	Number       	string   `json:"number"`
	KKM 			string   `json:"kkm"`
	Department    	string   `json:"department"`
	Nalog     		string   `json:"nalog"`
	CheckType     	string   `json:"CheckType"`
	CheckTotal     	string   `json:"CheckTotal"`
	TotalByCasche   string   `json:"TotalByCasche"`
	TotalByElectron string   `json:"TotalByElectron"`
	CheckDate     	string   `json:"check_date"`	
}


//CreateCheck добавляет чек 
func CreateCheck(c *gin.Context) {
	fmt.Printf("number='%s'",c.PostForm("number"))
	var err error
	var json JsonCheck
	err = c.BindJSON(&json)
	if err != nil {
		log.Panic(err)
	}

	//fmt.Printf("number='%s'\n",json.Number)
	
	check := Check{
		Number: json.Number,
		KKM: utils.ParseInt(json.KKM),
		Department: utils.ParseInt(json.Department),
		Nalog: utils.ParseInt(json.Nalog),
		CheckType: utils.ParseInt(json.CheckType),
		CheckTotal: utils.ParseFloat(json.CheckTotal),
		TotalByCasche: utils.ParseFloat(json.TotalByCasche),
		TotalByElectron: utils.ParseFloat(json.TotalByElectron),
		CheckDate: utils.ParseInt(json.CheckDate),	
		}
	
	var message string;
	var status int

	db := database.Database()
	err = db.Save(&check).Error
	fmt.Println(err)
	switch(err){
		case nil :
			status =  http.StatusCreated
			message = "Item created successfully!"
		default:
			status =  http.StatusForbidden
			message = fmt.Sprint(err)
	}	
	c.JSON(http.StatusCreated, gin.H{"status": status, "message": message, "resourceId": json.ID})	
	db.Close()
 }

 //FetchAllCheck  возвращает все чеки
 func FetchAllCheck(c *gin.Context)  {
	var items []Check
	var _items []JsonCheck
	db := database.Database()
	db.Find(&items)
 
	if len(items) <= 0 {
	   c.JSON(http.StatusNotFound, gin.H{"status": http.StatusNotFound, "message": "No todo found!"})
	   return
	}
 
	//transforms the todos for building a good response
	for _, item := range items {
	   _items = append(
		  _items, JsonCheck{
			 ID: string(item.ID),
			 Number: string(item.Number),
			 KKM: string(item.KKM),
			 Department: string(item.Department),
			 Nalog:  string(item.Nalog),
			 CheckType: string(item.CheckType),
			 CheckTotal: utils.ParseFloat2String(item.CheckTotal),
			 TotalByCasche: utils.ParseFloat2String(item.TotalByCasche),
			 TotalByElectron: utils.ParseFloat2String(item.TotalByElectron),
			 CheckDate: string(item.CheckDate),
			 //Sync: item.Sync,
			 })
	}
	c.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": _items})
	db.Close()
 }