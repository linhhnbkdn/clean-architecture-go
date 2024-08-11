package main

import (
	ginItem "cleanArchitechureGo/module/item/transport/gin"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

func main() {
	dsnDB := "user=postgres password=postgres dbname=postgres port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(mysql.Open(dsnDB), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to database", db)
	r := gin.Default()

	v1 := r.Group("/v1")
	{
		item := v1.Group("/item")
		{
			item.POST("", ginItem.CreateNewItem(db))
		}
	}
}
