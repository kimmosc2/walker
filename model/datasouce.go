package model

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"school-walker/walker/pkg/setting"
)

var DB *gorm.DB


func init() {
	var (
		err                          error
		dbName, user, password, host string
	)

	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}

	// dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	// tablePrefix = sec.Key("TABLE_PREFIX").String()

	DB, err = gorm.Open(mysql.Open(fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName)), &gorm.Config{})

	if err != nil {
		log.Println(err)
	}

	// gorm.DefaultTableNameHandler = func(DB *gorm.DB, defaultTableName string) string {
	// 	return tablePrefix + defaultTableName
	// }
	//
	// DB.SingularTable(true)
	// DB.LogMode(true)
	// DB.DB().SetMaxIdleConns(10)
	// DB.DB().SetMaxOpenConns(100)

}

func CloseDB() {
	s, _ := DB.DB()
	s.Close()
}
