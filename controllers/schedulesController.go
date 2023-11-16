package controllers

import (
	"iKarate-GO/initializers"
	"iKarate-GO/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PostSchedule(c *gin.Context) {

	var body struct {
		Description string
		DojoID      uuid.UUID
		UserID      uuid.UUID
		Active      string
	}

	x := c.Bind(&body)
	if x != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}
	active, _ := strconv.ParseBool(body.Active)
	schedule := models.Schedule{
		Description: body.Description,
		DojoID:      body.DojoID,
		UserID:      body.UserID,
		Active:      active,
	}

	result := initializers.DB.Create(&schedule)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear usuario... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetSchedules(c *gin.Context) {
	//var page = c.DefaultQuery("page", "1")
	//var limit = c.DefaultQuery("limit", "10")

	//intPage, _ := strconv.Atoi(page)
	//intLimit, _ := strconv.Atoi(limit)
	//offset := (intPage - 1) * intLimit

	var schedules []models.Schedule
	results := initializers.DB.Order(" description asc").Find(&schedules)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(schedules), "data": schedules})
}
