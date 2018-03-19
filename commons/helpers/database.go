package helpers

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func InitConnection(dialect string, url string) (*gorm.DB, error) {
	var err error
	db, err = gorm.Open(dialect, url)
	return db, err
}

func GetConnection() *gorm.DB {
	return db
}
