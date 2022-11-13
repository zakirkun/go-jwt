package service

import (
	"errors"
	helper "github.com/vandenbill/brand-commerce/Helper"
	web "github.com/vandenbill/brand-commerce/Model/Web"
	repo "github.com/vandenbill/brand-commerce/Repo"
)

type AuthService interface {
	VerifyCredential(loginDto web.LoginDto) error
}

type authServiceImpl struct {
	userRepo repo.UserRepo
}

func NewAuthServiceImpl(userRepo repo.UserRepo) AuthService {
	return &authServiceImpl{userRepo: userRepo}
}

func (a *authServiceImpl) VerifyCredential(loginDto web.LoginDto) error {
	user, err := a.userRepo.FindByEmail(loginDto.Email)
	if err != nil {
		return errors.New("Wrong email")
	}

	isValidPassword := helper.ComparePassword(user.Password, []byte(loginDto.Password))
	if !isValidPassword {
		return errors.New("Wrong password")
	}
	return nil
}
