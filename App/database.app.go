package app

import (
	domain "github.com/vandenbill/brand-commerce/Model/Domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func NewDB() *gorm.DB {
	LoadEnv()

	dsn := os.Getenv("DB_URL")
	L.Info("Load DB_URL from env var")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		L.Panic(err)
	}
	L.Info("Load connection to DB")

	DBAutoMigrate(db)
	return db
}

func DBAutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&domain.User{})
	if err != nil {
		L.Panic(err)
	}
	L.Info("Migrate model to DB")
}
