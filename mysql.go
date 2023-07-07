package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name         string
	Age          sql.NullInt64
	Birthday     *time.Time
	Email        string  `gorm:"type:varchar(100);unique_index"`
	Role         string  `gorm:"size:255"`        //设置字段的大小为255个字节
	MemberNumber *string `gorm:"unique;not null"` // 设置 memberNumber 字段唯一且不为空
	Num          int     `gorm:"AUTO_INCREMENT"`  // 设置 Num字段自增
	Address      string  `gorm:"index:addr"`      // 给Address 创建一个名字是  `addr`的索引
	IgnoreMe     int     `gorm:"-"`               //忽略这个字段
}

type UserInfo struct {
	ID    uint
	Name  string
	Hobby string
}

type Test struct {
	ID   uint
	Name string
}

//项目初始化
func main() {
	var userList UserInfo

	db := InitDB()

	//defer db.Close() //延时关闭

	// 获取第一条匹配的记录
	//db.Where("name = ?", "jinzhu").First(&user)

	//user := UserInfo{Name: "Jinzhu"}

	//result := db.Create(&user) // 通过数据的指针来创建

	//fmt.Println(result)  // 在终端打印 Hello World!

	// 自动迁移
	//db.AutoMigrate(&UserInfo{})

	router := gin.Default()

	router.GET("/", func(c *gin.Context) {

		//list := db.Find(&userList)
		//list := db.Where("id > ?", "0").Scan(&userList)
		//list :=  db.Table("user_infos").Select("id, name, hobby").Where("id > ?", 1).Scan(&userList)
		list := db.Where("id <> ?", 0).Find(&userList)

		//fmt.Printf("%v",list)
		currentTime := time.Now()

		c.JSON(200, gin.H{
			"status":      0,
			"message":     "message",
			"data":        list.Value,
			"currentTime": currentTime,
		})
	})

	router.GET("/user/:name/:action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action

		user := UserInfo{Name: name, Hobby: action}
		db.Create(&user)

		c.String(http.StatusOK, message)
	})

	router.Run(":8000")
}

//数据库连接
func InitDB() *gorm.DB {

	driverName := "mysql"
	host := "127.0.0.1"
	port := "3306"
	database := "test"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset,
	)

	db, err := gorm.Open(driverName, args)
	//db, err := gorm.Open("mysql", "user:islot@/blog?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic("failed to connect database,err:" + err.Error())
	}

	//自动创建数据表
	db.AutoMigrate(&User{})

	return db

}
