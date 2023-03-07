package initializers

import "karate-backend/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
