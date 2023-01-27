package connect

import (
	_ "Web-chat/config"
	"Web-chat/pkg/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func init() {
	var err error
	dsn := os.Getenv("DB_CONNECTION")
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{DisableAutomaticPing: false})
	if err != nil {
		panic(err)
	}
	autoMigrate(DB)
}

func autoMigrate(DB *gorm.DB) {
	DB.AutoMigrate(&models.Chat{}, &models.Message{}, &models.Social{}, &models.User{})
}
