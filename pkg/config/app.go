package config

import (
	"log"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func Connect_Mysql_server_and_ReturnDB() *gorm.DB{
	//var db *gorm.DB
	db, err := gorm.Open("mysql", "root:@tcp(127.0.0.1:3306)/blog_post_2?charset=utf8&parseTime=True&loc=Local")

	if err!=nil{
		log.Fatal("From app.go",err)
	}
	return db
}