package services

import (
	"LoveDiary/model"
	"LoveDiary/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

var err error
var db *gorm.DB

func InitDb() {
	// 拼接dsn参数
	dsn := fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		utils.DbUser,
		utils.DbPassWord,
		utils.DbHost,
		utils.DbPort,
		utils.DbName)

	// 连接mySql,获得Db类型实例
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		fmt.Printf("连接数据库失败，请检查参数：", err)
	}

	// 禁止默认表名的复数形式, 目前文档找不到该设置
	// db.SingularTable(true)
	//db = db.Debug().Session(&gorm.Session{SingularTable:true})

	// 自动迁移
	_ = db.AutoMigrate(&model.User{})

	// 获取通用数据库对象 sql.DB，然后使用其提供的功能
	sqlDB, _ := db.DB()
	//defer sqlDB.Close()
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(10)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(100)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	//defer sqlDB.Close()
}

func CloseDb() {
	sqlDB, err := db.DB()
	if err != nil {
		fmt.Println("获取数据库连接对象失败, err:", err)
		return
	}
	err = sqlDB.Close()
	if err != nil {
		fmt.Println("关闭数据库连接失败, err:", err)
		return
	}
}
