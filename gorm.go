package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// UserInfo --> 数据表
type UserInfo struct {
	ID     uint
	Name   string
	Gender string
	Hobby  string
}

func (u *UserInfo) TableName() string {
	return "user_info"
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/mygin?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("database open failed ,err=", err)
		return
	}

	//创建表 自动迁移 （把结构体和数据表进行对应 ）
	db.AutoMigrate(&UserInfo{})

	user := UserInfo{Name: "Jinzhuss", Gender: "asdfsdf", Hobby: "sdf"}

	db.Create(&user) // 通过数据的指针来创建

	// 查询，将查询出来的数据赋值到u，要更改u所以要加指针
	var u UserInfo

	db.First(&u) //查询表中的第一条数据保存到u中
	fmt.Printf("u:%#v\n", u)

	// 更新
	db.Model(&u).Update("hobby", "money")

	// 删除
	// db.Delete(&u)

	var U UserInfo
	db.First(&U)
	fmt.Printf("U:%#v\n", U)
}
