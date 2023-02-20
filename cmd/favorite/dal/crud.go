package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	// 连接数据库
	db, err := gorm.Open(mysql.Open("root:12345678@tcp(47.92.171.253:23306)/tiktok?charset=utf8"))
	if err != nil {
		panic(fmt.Errorf("cannot establish db connection: %w", err))
	}
	g := gen.NewGenerator(gen.Config{
		OutPath: "/house/liuzhiwei/code/go/test/dal/query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	g.UseDB(db)
	// Generate basic type-safe DAO API
	g.ApplyBasic(g.GenerateAllTable()...)
	g.Execute()
}
