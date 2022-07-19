package repos

import "gorm.io/gorm"

type MySQLClient struct {
	database *gorm.DB
}