package config

import (
	"github.com/isaacgoncod/go-project-example/schemas"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitializeMySQL() (*gorm.DB, error) {
	logger := GetLogger("mysql")

	dsn := "root:@tcp(127.0.0.1:3306)/go_project_example"

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
