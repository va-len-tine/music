package global

import (
	"fmt"
	"github.com/spf13/viper"
	"goServer/error"
	"goServer/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"path/filepath"
)

var (
	CF model.ConfigServer
	MyDB     *gorm.DB
)

func InitViper() {
	fmt.Println(filepath.Abs("."))
	v := viper.New()
	v.SetConfigFile("config.yaml")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	error.HandleErr(err, "配置文件初始化失败")
    err = v.Unmarshal(&CF)
	error.HandleErr(err, "配置文件初始化失败")
}

func InitLog(){
	if CF.Log.Outputfile{
		logFile, _ := os.OpenFile("./log/" + CF.Log.Logname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		log.SetOutput(logFile)
	}
	log.SetPrefix(CF.Log.Logprefix)
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime)
}

func InitDB() *gorm.DB{
	m := CF.Mysql
	dsn := m.Username + ":" + m.Password + "@tcp(" + m.Hostname + ":" + m.Port + ")/" + m.Dbname + "?" + "charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN: dsn, // data source name
		DefaultStringSize: 256, // default size for string fields
		DisableDatetimePrecision: true, // disable datetime precision, which not supported before MySQL 5.6
		DontSupportRenameIndex: true, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
		DontSupportRenameColumn: true, // `change` when rename column, rename column not supported before MySQL 8, MariaDB
		SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
	}), &gorm.Config{})
	error.HandleErr(err, "数据库初始化失败")
	return db
}
