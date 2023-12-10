package initializers

import "iKarate-GO/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{},
		&models.Permission{},
		&models.Dojo{},
		&models.Category{},
		&models.Schedule{},
		&models.Payment{},
		&models.Product{},
		&models.Ticket{},
		&models.TicketDetail{})
}
