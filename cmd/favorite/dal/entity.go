package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Favorite struct {
	ID      uint `gorm:"primaryKey"`
	UserId  uint
	VideoId uint
}

func main() {
	dsn := "root:12345678@tcp(47.92.171.253:23306)/tiktok?charset=utf8"
	println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&Favorite{})
}
