package db

import (
	"os"
	"fmt"
	"sync"

	"github.com/jinzhu/gorm"
 	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbInstance *gorm.DB
	once sync.Once
)

func DB() *gorm.DB {
	once.Do(func() {
		username := os.Getenv("MYSQL_USER")
		password := os.Getenv("MYSQL_PASSWORD")
		dbName := os.Getenv("MYSQL_DATABASE")
		link := fmt.Sprintf("%s:%s@/%s?charset=utf8&parseTime=True&loc=Local", username, password, dbName)

		var err error
		dbInstance, err = gorm.Open("mysql", link)
		if err != nil {
			panic(err)
		}
	})
	return dbInstance
}