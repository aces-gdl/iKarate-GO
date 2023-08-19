package controllers

import (
	"fmt"
	"iKarate-GO/initializers"
	"iKarate-GO/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostCreateTimeSlots(c *gin.Context) {

	var body struct {
		TournamentID uuid.UUID
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	var tournament models.Tournament

	result := initializers.DB.Where("ID = ?", body.TournamentID).First(&tournament)

	if result.Error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Torneo no encontrado...",
		})
		return

	}
	fmt.Println(tournament)
}
