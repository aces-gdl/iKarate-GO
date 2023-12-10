package controllers

import (
	"iKarate-GO/initializers"
	"iKarate-GO/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func PostProduct(c *gin.Context) {

	var body struct {
		Description  string
		RegularPrice string
		Type         string
		Active       string
	}

	x := c.Bind(&body)
	if x != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}
	regularPrice, _ := strconv.ParseFloat(strings.TrimSpace(body.RegularPrice), 64)
	productType, _ := strconv.Atoi(body.Type)
	active, _ := strconv.ParseBool(body.Active)
	Product := models.Product{
		Description:  body.Description,
		RegularPrice: regularPrice,
		Type:         productType,
		Active:       active,
	}

	result := initializers.DB.Create(&Product)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear usuario... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{})
}

func GetProducts(c *gin.Context) {
	//var page = c.DefaultQuery("page", "1")
	//var limit = c.DefaultQuery("limit", "10")

	//intPage, _ := strconv.Atoi(page)
	//intLimit, _ := strconv.Atoi(limit)
	//offset := (intPage - 1) * intLimit

	var products []models.Product
	results := initializers.DB.Order("description asc").Find(&products)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(products), "data": products})
}
