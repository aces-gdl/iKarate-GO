package controllers

import (
	"iPadel-GO/initializers"
	"iPadel-GO/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func GetTournaments(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var tournaments []models.Tournament
	results := initializers.DB.Preload(clause.Associations).Limit(intLimit).Offset(offset).Find(&tournaments)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(tournaments), "data": tournaments})
}

func PostTournaments(c *gin.Context) {
	//var body models.Tournament

	var body struct {
		Description    string
		StartDate      string
		EndDate        string
		HostClubID     string
		GameDuration   int
		NumberOfCourts int
		Active         bool
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}
	startDate, _ := time.Parse("2006-01-02", body.StartDate)
	endDate, _ := time.Parse("2006-01-02", body.EndDate)
	clubID, _ := uuid.Parse(body.HostClubID)
	tournament := models.Tournament{
		Description:    body.Description,
		StartDate:      startDate,
		EndDate:        endDate,
		HostClubID:     clubID,
		NumberOfCourts: body.NumberOfCourts,
		GameDuration:   body.GameDuration,
		Active:         body.Active,
	}
	//fmt.Println(tournament)
	result := initializers.DB.Create(&tournament)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear torneo... ",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{})
}
