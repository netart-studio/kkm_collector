package models

import (
   "github.com/jinzhu/gorm"
   _"github.com/jinzhu/gorm/dialects/mysql"
)

import "avto98/kkm_collector/config"

var conf = config.Get_config() 

//Database инициализация базы данных
func Database() *gorm.DB {
	//open a db connection
	//db, err := gorm.Open("mysql", "root:mysql@tcp(127.0.0.1)/kkm_collector?parseTime=true")
	
	db, err := gorm.Open(conf.DataBase, conf.ConnectString)

	if err != nil {
	   panic("failed to connect database")
	}
	return db
 }