package config

import (
	"github.com/isaacgoncod/go-project-example/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySQL() (*gorm.DB, error) {
	logger := GetLogger("mysql")



	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		logger.Errf("MySQL Opening Error: %v", err)

		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errf("MySQL AutoMigration Error: %v", err)

		return nil, err
	}

	return db, nil
}
