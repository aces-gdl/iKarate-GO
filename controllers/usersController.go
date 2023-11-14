package controllers

import (
	"iKarate-GO/initializers"
	"iKarate-GO/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUsers(c *gin.Context) {
	//var page = c.DefaultQuery("page", "1")
	//var limit = c.DefaultQuery("limit", "10")
	var categoryID = c.DefaultQuery("CategoryID", "")
	var PermissionID = c.DefaultQuery("PermissionID", "")
	var ID = c.DefaultQuery("ID", "")
	//intPage, _ := strconv.Atoi(page)
	//intLimit, _ := strconv.Atoi(limit)
	//offset := (intPage - 1) * intLimit

	var whereClause = " 1 = 1 "
	if ID != "" {
		whereClause = whereClause + " AND ID = '" + ID + "' "
	}
	if categoryID != "" {
		whereClause = whereClause + " AND category_id = '" + categoryID + "' "
	}
	if PermissionID != "" {
		whereClause = whereClause + " AND permission_id = '" + PermissionID + "'"
	}

	type userExtended struct {
		ID                  uuid.UUID
		CategoryID          uuid.UUID
		PermissionIDID      uuid.UUID
		Name                string
		FamilyName          string
		GivenName           string
		Email               string
		CategoryDescription string
		Color1              string
		Color2              string
		Level               int
		PermissionID        uuid.UUID
		StartDate           time.Time
		Observations        string
		ContactName         string
		ContactPhone        string
		AttendingClassID    uuid.UUID
		HasPicture          int
		SeleccionadoID      uuid.UUID
		Birthday            time.Time
	}
	var usersExtended []userExtended

	const queryString = `select 
							u.id,
							u.name, 
							u.family_name , 
							u.given_name , 
							u.email,
							u.category_id, 
							u.birthday,
							u.category_id ,
							c.description as category_description,
							c.color1,
							c.color2,
							c.level,
							u.permission_id,
							u.start_date,
							u.observations,   
							u.contact_name,
							u.contact_phone,
							u.attending_class_id,
							u.Seleccionado_id,
							u.birthday
						from users u
							inner join categories c on u.category_id = c.id`

	results := initializers.DB.Debug().Raw(queryString).Where(whereClause).Scan(&usersExtended)

	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(usersExtended), "data": usersExtended})
}

func UpdateUsers(c *gin.Context) {
	var body models.User

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	user := models.User{
		ID:           body.ID,
		Email:        body.Email,
		GivenName:    body.GivenName,
		FamilyName:   body.FamilyName,
		GoogleID:     body.GoogleID,
		ImageURL:     body.ImageURL,
		Name:         body.Name,
		PermissionID: body.PermissionID,
		CategoryID:   body.CategoryID,
	}

	results := initializers.DB.Save(&user)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": results.RowsAffected, "data": user})

}
