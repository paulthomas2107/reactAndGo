package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connect ...
func Connect() {

	_, err := gorm.Open(mysql.Open("root:Caitlin1966@/yt_go_auth"), &gorm.Config{})

	if err != nil {
		panic("*** Could not connect to the database")
	}
}
