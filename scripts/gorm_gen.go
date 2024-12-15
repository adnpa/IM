package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// 连接需要的db
	//dsn := "host=localhost user=srbbs password=123456 dbname=srbbs port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	//if gormdb, err := gorm.Open(postgres.Open(dsn)); err != nil {
	//	log.Println("postgresql connecting error")
	//	panic(err)
	//}
	gormdb, _ := gorm.Open(mysql.Open("root:root@(192.168.1.129:3306)/goim?charset=utf8mb4&parseTime=True&loc=Local"))

	////////////////////////////////////////////
	g.UseDB(gormdb) // reuse your gorm db

	// Generate basic type-safe DAO API for struct `model.User` following conventions

	g.ApplyBasic(
		g.GenerateModelAs("user", "User"),
		g.GenerateModelAs("friend", "Friend"),
		g.GenerateModelAs("friend_request", "FriendRequest"),
	)
	g.ApplyBasic(
		// Generate structs from all tables of current database
		g.GenerateAllTable()...,
	)
	// Generate the code
	g.Execute()
}
