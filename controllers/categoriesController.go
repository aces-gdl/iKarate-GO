package controllers

import (
	"iKarate-GO/initializers"
	"iKarate-GO/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func PostCategory(c *gin.Context) {

	var body struct {
		Description string
		Color1      string
		Color2      string
		Level       string
		Active      string
	}

	x := c.Bind(&body)
	if x != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}
	level, _ := strconv.Atoi(body.Level)
	active, _ := strconv.ParseBool(body.Active)
	category := models.Category{
		Description: body.Description,
		Color1:      body.Color1,
		Color2:      body.Color2,
		Level:       level,
		Active:      active,
	}

	result := initializers.DB.Create(&category)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear usuario... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetCatgories(c *gin.Context) {
	//var page = c.DefaultQuery("page", "1")
	//var limit = c.DefaultQuery("limit", "10")

	//intPage, _ := strconv.Atoi(page)
	//intLimit, _ := strconv.Atoi(limit)
	//offset := (intPage - 1) * intLimit

	var categories []models.Category
	results := initializers.DB.Order("level asc, description asc").Find(&categories)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(categories), "data": categories})
}
