package services

import (
	"log"

	"github.com/dasha-kinsely/ostruct/controllers/repos"
	"github.com/dasha-kinsely/ostruct/models/dto"
	"github.com/dasha-kinsely/ostruct/models/entities"
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
	user := entities.User{}
	user.Email = signupRequest.Email
	user.Username = signupRequest.Username
	user.PasswordHash = signupRequest.Password
	log.Println(user)
	return user, nil
	// This step should not return any users
	/*_, err := repo.userRepo.FindByEmail(signupRequest.Email)
	// If gorm does not find any existing users with specified email address, proceed
	if err == gorm.ErrRecordNotFound {
		//user, err := repo.userRepo.InsertUser(user)
		return user, err
	} else {
		return entities.User{}, errors.New("problem occurred at data accessing layer, failed to create user...")
	}*/
}