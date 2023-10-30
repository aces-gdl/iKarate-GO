package controllers

import (
	"iKarate-GO/initializers"
	"iKarate-GO/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetDojos(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var dojos []models.Dojo
	results := initializers.DB.Preload("Manager").Order("description asc").Limit(intLimit).Offset(offset).Find(&dojos)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(dojos), "data": dojos})
}

func PostDojos(c *gin.Context) {

	var body struct {
		Name        string
		Description string
		ShortName   string
		Address     string
		Phone       string
		ManagerID   uuid.UUID
		Active      bool
	}

	resultTest := c.Bind(&body)
	if resultTest != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	dojo := models.Dojo{
		Name:        body.Name,
		Description: body.Description,
		ShortName:   body.ShortName,
		Address:     body.Address,
		Phone:       body.Phone,
		ManagerID:   body.ManagerID,
		Active:      body.Active,
	}

	//fmt.Println(tournament)
	result := initializers.DB.Debug().Create(&dojo)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": result.Error.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}
