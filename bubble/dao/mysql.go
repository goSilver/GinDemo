package dao

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

/**
 * @Description 定义全局的数据库对象
 **/
var (
	DB *gorm.DB
)

/**
 * @Description 初始化mysql
 * @Param
 * @return err
 **/
func InitMySQL() (err error) {
	DB, err = gorm.Open("mysql", "root:@(localhost)/go?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		fmt.Printf("connect msyql failed, err: %v \n", err)
		return
	}
	// 测试是否能够连通
	return DB.DB().Ping()
}
