package db

import (
	"fmt"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBInitialize() {
	//数据库连接
	viper.SetConfigFile("application.yml")
	viper.AddConfigPath("../config/application.yml")
	err := viper.ReadInConfig()
	if err != nil {
		return
	}

	user := viper.GetString("database.user")
	fmt.Println(user)
	// 参考 https://github.com/go-sql-driver/mysql#dsn-data-source-name 获取详情
	dsn := "root:root@tcp(127.0.0.1:3306)/xiaozheai?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	fmt.Println(db)
}
