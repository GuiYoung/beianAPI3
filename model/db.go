package model

import (
	"beianAPI/utils"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"time"
)

var Db *gorm.DB

func InitDb() (err error) {
	// init config
	if err = utils.Init("asserts/config.ini"); err != nil {
		return err
	}

	// create connection
	err = initMysql(&utils.Conf.MySQL)
	if err != nil {
		return err
	}

	// create table
	_ = Db.AutoMigrate(&Beian{}, &User{})

	sqlDB, _ := Db.DB()

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(10 * time.Second)

	return
}

func initMysql(config *utils.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User, config.Password, config.IP, config.Port, config.Database)
	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger.Default.LogMode(logger.Silent),
		SkipDefaultTransaction: true,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		return err
	}
	return
}
