package controllers

import (
	"iKarate-GO/initializers"
	"iKarate-GO/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostClub(c *gin.Context) {
	var body models.Club

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	club := models.Club{
		ID:          body.ID,
		Name:        body.Name,
		Description: body.Description,
		Contact:     body.Contact,
		Phone:       body.Phone,
		Address:     body.Address,
	}

	result := initializers.DB.Create(&club)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear usuario... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
