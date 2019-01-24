package model

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

func init() {
	var err error
	db, err = gorm.Open("mysql", "root:123456@(localhost:3305)/test_1?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	} else {
		fmt.Println("Database connect!")
	}
	// 自动迁移
	db.AutoMigrate(&User{})
}

// func CloseDb() {
// 	defer Db.Close()
// }
