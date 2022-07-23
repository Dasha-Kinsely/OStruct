package services

import (
	"errors"

	"github.com/dasha-kinsely/ostruct/controllers/repos"
	"github.com/dasha-kinsely/ostruct/models/dto"
	"github.com/dasha-kinsely/ostruct/models/entities"
	"gorm.io/gorm"
)

type UserService interface {
	CreateUser(signupRequest dto.SignupRequest) (entities.User, error)
}

type userService struct {
	userRepo repos.UserRepo
}

func NewUserService(userRepo repos.UserRepo) UserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (repo *userService) CreateUser(signupRequest dto.SignupRequest) (entities.User, error) {
	_, err := repo.userRepo.FindByEmail(signupRequest.Email)
	// If gorm does not find any existing users with specified email address, proceed
	if errors.Is(err, gorm.ErrRecordNotFound) {
		newUser := entities.User{}
		newUser.Username = signupRequest.Username
		newUser.Email = signupRequest.Email
		newUser.Password = signupRequest.Password
		newUser, err := repo.userRepo.InsertUser(newUser)
		return newUser, err
	}
	user := entities.User{}
	return user, errors.New("Problem occurred at data accessing layer: (user already exists!!!) failed to create user...")
}