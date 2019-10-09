package conf

import (
	"btube/model"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	//
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//MySQLConnect isa singleton pool of mysql connection.
var MySQLConnect *gorm.DB

// Database connet to mysql.
func Database(connString string) {
	db, err := gorm.Open("mysql", connString)
	db.LogMode(true)
	// Error
	if err != nil {
		panic(err)
	}
	if gin.Mode() == "release" {
		db.LogMode(false)
	}
	db.DB().SetMaxIdleConns(20)
	db.DB().SetMaxOpenConns(100)
	db.DB().SetConnMaxLifetime(time.Second * 30)

	MySQLConnect = db

	migration()
}

func migration() {
	err := MySQLConnect.AutoMigrate(&model.User{}).AutoMigrate(&model.Video{}).Error
	if err != nil {
		fmt.Print("can't not create table.")
		panic(err)
	}
}
