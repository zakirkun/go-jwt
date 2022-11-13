package app

import (
	helper "github.com/vandenbill/brand-commerce/Helper"
	domain "github.com/vandenbill/brand-commerce/Model/Domain"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

func NewDB() *gorm.DB {
	LoadEnv()
	dsn := os.Getenv("DB_URL")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	helper.PanicIfError(err)
	DBAutoMigrate(db)
	return db
}

func DBAutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&domain.User{})
	helper.PanicIfError(err)
}
