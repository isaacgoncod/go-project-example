package handler

import (
	"github.com/isaacgoncod/go-project-example/config"
	"gorm.io/gorm"
)

var (
	logger *config.Logger
	db     *gorm.DB
)

func InitializeHandler() {
	logger = config.GetLogger("handler")
	db = config.GetMySQL()
}
