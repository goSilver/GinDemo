package main

/*
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	//Default返回一个默认的路由引擎
	r := gin.Default()
	r.GET("/user/search", func(c *gin.Context) {
		// 可以添加默认值
		username := c.DefaultQuery("username", "小王子")
		//username := c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})
	r.Run()
}
*/

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type User struct {
	ID   int64
	Name string `gorm:"default:'小王子'"`
	Age  int64
}

func main() {
	db, err := gorm.Open("mysql", "root:@(127.0.0.1:3306)/go?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库链接ok")
	user := User{Name: "", Age: 18}
	//db.Create(&user) // 创建user

	//db.Find(&user)

	// Get first matched record
	first := db.Where("name = ?", "q1mi").First(&user)
	//// SELECT * FROM users WHERE name = 'jinzhu' limit 1;

	fmt.Println(first)

	defer db.Close()

	// db.Xx
}
