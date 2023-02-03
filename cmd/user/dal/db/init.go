package db

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {
	var err error
	addr := viper.GetString("db.addr")
	port := viper.GetString("db.port")
	db := viper.GetString("db.database")
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
	if DB.Migrator().HasTable(User{}) {
		return
	}
	err = DB.Migrator().CreateTable(&User{})
	if err != nil {
		panic("failed to create the user table")
	}
}
