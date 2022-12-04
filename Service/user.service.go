package service

import (
	"github.com/mashingan/smapping"
	helper "github.com/vandenbill/brand-commerce/Helper"
	domain "github.com/vandenbill/brand-commerce/Model/Domain"
	web "github.com/vandenbill/brand-commerce/Model/Web"
	repo "github.com/vandenbill/brand-commerce/Repo"
)

type UserService interface {
	Create(userDto web.UserDto) (domain.User, error)
	Update(sub uint, userDto web.UserDto) (domain.User, error)
	Delete(id uint) error
	FindAll() ([]domain.User, error)
	FindByID(id uint) (domain.User, error)
	FindByEmail(email string) (domain.User, error)
}

// TODO implement logging here

type userServicImpl struct {
	userRepo repo.UserRepo
}

func NewUserServicImpl(userRepo repo.UserRepo) UserService {
	return &userServicImpl{userRepo: userRepo}
}

func (u userServicImpl) Create(userDto web.UserDto) (domain.User, error) {
	user := domain.User{}
	err := smapping.FillStruct(&user, smapping.MapFields(&userDto))

	user.Password = helper.HashAndSalt([]byte(user.Password))

	user, err = u.userRepo.Save(user)

	return user, err
}

func (u userServicImpl) Update(id uint, userDto web.UserDto) (domain.User, error) {
	user, err := u.userRepo.FindByID(id)
	if err != nil {
		return domain.User{}, err
	}
	user.Name = userDto.Name
	user.Email = userDto.Email
	user.Password = helper.HashAndSalt([]byte(userDto.Password))
	user, err = u.userRepo.Update(user)
	return user, err
}

func (u userServicImpl) Delete(id uint) error {
	err := u.userRepo.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (u userServicImpl) FindAll() ([]domain.User, error) {
	users, err := u.userRepo.FindAll()
	return users, err
}

func (u userServicImpl) FindByID(id uint) (domain.User, error) {
	user, err := u.userRepo.FindByID(id)
	return user, err
}

func (u userServicImpl) FindByEmail(email string) (domain.User, error) {
	var user domain.User
	user, err := u.userRepo.FindByEmail(email)
	return user, err
}
