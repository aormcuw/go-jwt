package initializer

import (
	"github.com/aormcuw/go-jwt/models"
)

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
