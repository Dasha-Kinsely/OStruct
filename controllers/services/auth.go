package services

import (
	"github.com/dasha-kinsely/ostruct/controllers/repos"
	"github.com/dasha-kinsely/ostruct/utils"
)

type AuthService interface {
	VerifyCredentials(email string, password string) error
}

type authService struct {
	userRepo repos.UserRepo
}

func NewAuthService(userRepo repos.UserRepo) AuthService{
	return &authService{
		userRepo: userRepo,
	}
}

func (repo *authService) VerifyCredentials(email string, password string) error {
	// check if user exists
	user, err := repo.userRepo.FindByEmail(email)
	if err != nil {
		return err
	}
	// check whether the password is correct
	if err := utils.DecryptPassword(user.PasswordHash, password); err != nil {
		return err
	}
	return nil
}

