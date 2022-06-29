package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// UserInfo 用户信息
type User struct {
	gorm.Model
	Name   string `gorm:"type:varchar(20);column:name"`
	Gender string `gorm:"type:varchar(20);column:gender"`
	Hobby  string `gorm:"type:varchar(20);column:hobby"`
}

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/test1?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// 自动迁移
	db.AutoMigrate(&User{})
	//
	// u1 := UserInfo{
	// 	Name:   "test1",
	// 	Gender: "男",
	// 	Hobby:  "篮球",
	// }
	// u2 := UserInfo{
	// 	Name:   "topgoer.com",
	// 	Gender: "女",
	// 	Hobby:  "足球",
	// }
	// // 创建记录
	// db.Create(&u1)
	// db.Create(&u2)

}
