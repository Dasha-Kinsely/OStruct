package config

import (
	"os"
	"strconv"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var SqlDB *gorm.DB

func SetupSqlDBConnection() {
	// get from .env files, assuming dotenv has been found and loaded from initialize.go
	dsn := os.Getenv("MYSQL_DSN")
	disableDatetimePrecision, parseErr := strconv.ParseBool(os.Getenv("MYSQL_DISABLE_DATETIME_PRECISION"))
	dontSupportRenameIndex, parseErr := strconv.ParseBool(os.Getenv("MYSQL_DONT_SUPPORT_RENAME_INDEX"))
	dontSupportRenameColumn, parseErr := strconv.ParseBool(os.Getenv("MYSQL_DONT_SUPPORT_RENAME_COLUMN"))
	skipInitializeWithVersion, parseErr := strconv.ParseBool(os.Getenv("MYSQL_SKIP_INITIALIZE_WITH_VERSION"))
	if parseErr != nil {
		panic("Error parsing boolean variables from .env files, please check your .env files...")
	}
	// Database connection configuration, does not require returning things
	db, dbErr := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  disableDatetimePrecision,
		DontSupportRenameIndex:    dontSupportRenameIndex,
		DontSupportRenameColumn:   dontSupportRenameColumn,
		SkipInitializeWithVersion: skipInitializeWithVersion,
	}), &gorm.Config{})
	if dbErr != nil {
		panic("Error initializing MYSQL_DB...")
	}
	// point the initialized sql db to DB variable
	SqlDB = db
}

func GetSqlDB() *gorm.DB {
	return SqlDB
}

func CloseSqlDBConnection(db *gorm.DB) {
	database, err := db.DB()
	if err != nil {
		panic("Failed to close database...")
	}
	database.Close()
}