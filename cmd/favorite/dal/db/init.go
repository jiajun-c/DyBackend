package db

import (
	"favorite/dal/query"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Q *query.Query

func Init() {
	var err error
	db, err := gorm.Open(mysql.Open("root:12345678@tcp(47.92.171.253:23306)/tiktok?charset=utf8"))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	query.SetDefault(db)
	Q = query.Q
	if err != nil {
		panic(err)
	}
}
