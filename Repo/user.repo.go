package repo

import (
	domain "github.com/vandenbill/brand-commerce/Model/Domain"
	"gorm.io/gorm"
)

type UserRepo interface {
	Save(user domain.User) (domain.User, error)
	Update(user domain.User) (domain.User, error)
	Delete(id uint) error
	FindAll() ([]domain.User, error)
	FindByID(id uint) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
}

type userRepoImpl struct {
	dbConn *gorm.DB
}

func NewUserRepoImpl(dbConn *gorm.DB) UserRepo {
	return &userRepoImpl{dbConn: dbConn}
}

func (u *userRepoImpl) Save(user domain.User) (domain.User, error) {
	tx := u.dbConn.Create(&user)
	return user, tx.Error
}

func (u *userRepoImpl) Update(user domain.User) (domain.User, error) {
	tx := u.dbConn.Save(&user)
	return user, tx.Error
}

func (u *userRepoImpl) Delete(id uint) error {
	tx := u.dbConn.Unscoped().Delete(&domain.User{}, id)
	return tx.Error
}

func (u *userRepoImpl) FindAll() ([]domain.User, error) {
	var user []domain.User
	tx := u.dbConn.Find(&user)
	return user, tx.Error
}

func (u *userRepoImpl) FindByID(id uint) (domain.User, error) {
	var user domain.User
	tx := u.dbConn.First(&user, id)
	return user, tx.Error
}

func (u *userRepoImpl) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	tx := u.dbConn.Where("email = ?", email).First(&user)
	return user, tx.Error
}
