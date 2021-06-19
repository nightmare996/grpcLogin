package model

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nightmare996/grpcLogin/config"
)

var db *gorm.DB

func InitDB() (err error) {
	dsn := config.DataSourseConfig.UserName + ":" + config.DataSourseConfig.PassWord + "@tcp(" + config.DataSourseConfig.Host + ":" + config.DataSourseConfig.Port + ")/" + config.DataSourseConfig.Database + "?" + config.DataSourseConfig.Options
	db, err = gorm.Open("mysql", dsn)
	if err != nil {
		return
	}
	return nil
}

func Close() {
	db.Close()
}
