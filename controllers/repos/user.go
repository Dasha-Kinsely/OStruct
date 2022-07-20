package repos

import (
	"github.com/dasha-kinsely/ostruct/models/entities"
	"github.com/dasha-kinsely/ostruct/utils"
	"gorm.io/gorm"
)

type UserRepo interface {
	InsertUser(user entities.User) (entities.User, error)
	//UpdateUser(user entities.User) (entities.User, error): this should only be accessed by someone with admin privillege.
	FindByEmail(email string) (entities.User, error)
	FindByID(uid string) (entities.User, error)
	DeleteUser(uid string) (error)
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return &MySQLClient{
		database: db,
	}
}

func (repo *MySQLClient) InsertUser(user entities.User) (entities.User, error) {
	user.PasswordHash = utils.EncryptPassword(user.PasswordHash)
	if err := repo.database.Save(&user).Error; err != nil {
		return user, err
	}
	linkedProfile := entities.UserExtras{User: user}
	if err := repo.database.Table("userextras").Create(&linkedProfile).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (repo *MySQLClient) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	result := repo.database.Where("email = ?", email).Take(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (repo *MySQLClient) FindByID(uid string) (entities.User, error) {
	var user entities.User
	result := repo.database.Where("id = ?", uid).Take(&user)
	if result.Error != nil {
		return user, result.Error
	}
	return user, nil
}

func (repo *MySQLClient) DeleteUser(uid string) (error) {
	var user entities.User
	result := repo.database.Preload("User").Where("id= ?", uid).Take(&user)
	if result.Error != nil {
		return result.Error
	}
	repo.database.Delete(&user)
	return nil
}