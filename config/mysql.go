package config

import (
	"fmt"

	"github.com/labstack/gommon/log"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	MySQL *gorm.DB
)

func MysqlConnect() {
	var err error
	dsn := fmt.Sprintf(`%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local`, MySQLUser, MySQLPassword, MySQLHost, MySQLPort, MySQLDBName)

	MySQL, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error("failed connect mysql: " + err.Error())
	}
}
