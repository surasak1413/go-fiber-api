package database

import (
	"loly/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	dsn := "host=localhost user=loly password=P@ssw0rd dbname=lolydatabase port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = connection.AutoMigrate(&models.User{})
	if err != nil {
		return nil, err
	}

	return connection, nil
}
