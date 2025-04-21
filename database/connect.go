package database

import (
	"os"

	"github.com/aPPClanK/gotodolist/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dsn := os.Getenv("DB_PARAMS")
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	DB.AutoMigrate(&model.User{}, &model.Task{})
	return nil
}
