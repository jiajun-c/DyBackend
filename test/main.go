package main

import (
	"context"
	"fmt"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

type Message struct {
	FromUserID int64     `json:"from_user_id"`
	ToUserID   int64     `json:"to_user_id"`
	Content    string    `json:"content"`
	Status     int       `json:"status"`
	CreatedAt  time.Time `json:"created_at"`
}

func (m *Message) TableName() string {
	return "message"
}

func Init() {
	var err error
	addr := "47.92.171.253"
	port := "23306"
	db := "tiktok"
	MysqlDefaultDSN := fmt.Sprintf("root:12345678@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", addr, port, db)
	logrus.Info(MysqlDefaultDSN)
	DB, err = gorm.Open(mysql.Open(MysqlDefaultDSN),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}
	if DB.Migrator().HasTable(Message{}) {
		return
	}
	err = DB.Migrator().CreateTable(&Message{})
	if err != nil {
		panic("failed to create the user table")
	}
}

func main() {
	Init()
	if err := DB.WithContext(context.Background()).Create(&Message{
		FromUserID: 1,
		ToUserID:   2,
		Content:    "",
		Status:     0,
		CreatedAt:  time.Now().Add(time.Hour),
	}).Error; err != nil {
		fmt.Println(err)
	}
}
