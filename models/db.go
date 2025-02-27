package models

import (
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // remember to import mysql driver
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

// Init the database connection.
func init() {
	// dbType := conf.ServerCfg.Get(constant.DB_TYPE)
	// dbHost := conf.ServerCfg.Get(constant.DB_HOST)
	// dbPort := conf.ServerCfg.Get(constant.DB_PORT)
	// dbUser := conf.ServerCfg.Get(constant.DB_USER)
	// dbPwd := conf.ServerCfg.Get(constant.DB_PWD)
	// dbName := conf.ServerCfg.Get(constant.DB_NAME)

	// dbType := constant.DB_TYPE
	// dbHost := constant.DB_HOST
	// dbPort := constant.DB_PORT
	// dbUser := constant.DB_USER
	// dbPwd := constant.DB_PWD
	// dbName := constant.DB_NAME

	var err error

	Db, err = gorm.Open("mysql", "root:root@(localhost:3306)/scratch_maker?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalln("Fail to connect database!")
	}
	fmt.Printf("database connected!!!")

	Db.SingularTable(true)
	if !Db.HasTable(&User{}) {
		Db.CreateTable(&User{})
	}
	if !Db.HasTable(&Game{}) {
		Db.CreateTable(&Game{})
	}
	if !Db.HasTable(&Photo{}) {
		Db.CreateTable(&Photo{})
	}
}

// The base model of all models, including ID & CreatedAt & UpdatedAt.
type BaseModel struct {
	ID          uint      `json:"id" gorm:"primary_key;AUTO_INCREMENT" form:"id"`
	CreatedTime time.Time `json:"createdTime" gorm:"default: CURRENT_TIMESTAMP" form:"created_at"`
	UpdatedTime time.Time `json:"updatedTime" gorm:"default: CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP" form:"updated_at"`
}
