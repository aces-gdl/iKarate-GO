package controllers

import (
	"bufio"
	"fmt"
	"iKarate-GO/initializers"
	"iKarate-GO/models"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

func PostLoadUsers(c *gin.Context) {
	fileHeader, err := c.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Archivo no esta presente...",
		})
		return
	}

	PermissionID, _ := uuid.Parse(c.PostForm("PermissionID"))

	var Permission models.Permission
	result := initializers.DB.Where("ID = ?", PermissionID).First(&Permission)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error buscando Permisos...",
		})
		return
	}

	fileToImport, err := fileHeader.Open()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Archivo no esta presente...",
		})
		return
	}
	defer fileToImport.Close()

	var categories []models.Category

	initializers.DB.Find(&categories)

	fileScanner := bufio.NewScanner(fileToImport)

	fileScanner.Split(bufio.ScanLines)

	// Create password default

	hash, err := bcrypt.GenerateFromPassword([]byte("Fortin123"), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al convertir password a hash...",
		})
		return
	}
	catCounter := 0

	var pwdDefault = string(hash)
	for fileScanner.Scan() {
		var user models.User
		arrayUser := strings.Split(fileScanner.Text(), ",")
		user.GivenName = arrayUser[0]
		user.FamilyName = arrayUser[1]
		user.Email = arrayUser[2]
		user.Name = fmt.Sprintf("%s, %s", user.GivenName, user.FamilyName)
		user.PermissionID = Permission.ID
		user.Password = pwdDefault
		user.CategoryID = categories[catCounter].ID
		if catCounter > len(categories)-2 {
			catCounter = 0
		} else {
			catCounter = catCounter + 1
		}
		initializers.DB.Create(&user)
	}

	c.JSON(http.StatusOK, gin.H{})
}
