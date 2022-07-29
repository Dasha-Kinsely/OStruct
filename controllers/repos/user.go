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

func (db *MySQLClient) InsertUser(user entities.User) (entities.User, error) {
	user.Password = utils.EncryptPassword(user.Password)
	if err := db.database.Save(&user).Error; err != nil {
		return user, err
	}
	linkedProfile := entities.UserExtras{User: user}
	if err := db.database.Table("user_extras").Create(&linkedProfile).Error; err != nil {
		return user, err
	}
	return user, nil
}

func (db *MySQLClient) FindByEmail(email string) (entities.User, error) {
	var user entities.User
	// This will either return an error or a fully bound user entity obj.
	err := db.database.Where("email = ?", email).Take(&user)
	if err.Error != nil {
		return user, err.Error
	} else {
		return user, nil
	}
}

func (db *MySQLClient) FindByID(uid string) (entities.User, error) {
	var user entities.User
	err := db.database.Where("id = ?", uid).Take(&user)
	if err.Error != nil {
		return user, err.Error
	}
	return user, nil
}

func (db *MySQLClient) DeleteUser(uid string) (error) {
	var user entities.User
	result := db.database.Preload("User").Where("id= ?", uid).Take(&user)
	if result.Error != nil {
		return result.Error
	}
	db.database.Delete(&user)
	return nil
}