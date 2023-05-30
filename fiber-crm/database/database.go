package database

import (
	"github.com/jinzhu/gorm"
	_ "modernc.org/sqlite"
)

var (
	DBConn *gorm.DB
)
