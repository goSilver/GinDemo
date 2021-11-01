package main

import (
	"GinDemo/bubble/dao"
	"GinDemo/bubble/entity"
	"GinDemo/bubble/routers"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("connect mysql success")
	}
	// 延迟关闭数据库
	defer dao.DB.Close()
	// 模型关闭 自动迁移【把结构体和数据表进行对应】
	dao.DB.AutoMigrate(&entity.Todo{})
	// 注册路由
	r := routers.SetRouter()
	r.Run(":8090")
}
