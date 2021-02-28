package database

import (
	"github.com/paulthomas2107/reactandgo/github.com/paulthomas2107/reactandgo/main/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// DB ...
var DB *gorm.DB

// Connect ...
func Connect() {

	connection, err := gorm.Open(mysql.Open("root:Caitlin1966@/yt_go_auth"), &gorm.Config{})

	if err != nil {
		panic("*** Could not connect to the database")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})

}
