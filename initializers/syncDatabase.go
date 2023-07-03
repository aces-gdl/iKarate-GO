package initializers

import "iKarate-GO/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{}, &models.Club{}, &models.Court{})

}
