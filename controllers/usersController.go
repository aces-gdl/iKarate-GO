package controllers

import (
	"iKarate-GO/initializers"
	"iKarate-GO/models"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetUsers(c *gin.Context) {
	//var page = c.DefaultQuery("page", "1")
	//var limit = c.DefaultQuery("limit", "10")
	var CategoryID = c.DefaultQuery("CategoryID", "")
	var ScheduleID = c.DefaultQuery("ScheduleID", "")
	var Filter = c.DefaultQuery("filter", "")
	var ID = c.DefaultQuery("ID", "")
	//intPage, _ := strconv.Atoi(page)
	//intLimit, _ := strconv.Atoi(limit)
	//offset := (intPage - 1) * intLimit

	type userExtended struct {
		ID                    uuid.UUID
		CategoryID            uuid.UUID
		PermissionIDID        uuid.UUID
		Name                  string
		FamilyName            string
		GivenName             string
		Email                 string
		CategoryDescription   string
		Color1                string
		Color2                string
		Level                 int
		PermissionID          uuid.UUID
		StartDate             time.Time
		Observations          string
		ContactName           string
		ContactPhone          string
		ScheduleID            uuid.UUID
		HasPicture            int
		SeleccionadoID        uuid.UUID
		Birthday              time.Time
		PermissionDescription string
	}

	var usersExtended []userExtended

	var whereClause = " where 1 = 1 "
	if ID != "" {
		whereClause = whereClause + " AND ID = '" + ID + "' "
	}
	if CategoryID != "" {
		whereClause = whereClause + " AND category_id = '" + CategoryID + "' "
	}
	if ScheduleID != "" {
		whereClause = whereClause + " AND schedule_id = '" + ScheduleID + "' "
	}

	if len(Filter) > 0 {
		whereClause = whereClause + " AND p.description = '" + Filter + "'"
	}
	queryString := `select 
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
						u.schedule_id,
						u.Seleccionado_id,
						u.birthday,
						u.has_picture,
						p.description as permission_description
					from users u
						inner join categories c on u.category_id = c.id
						right outer  join permissions p on u.permission_id = p.id ` + whereClause

	results := initializers.DB.Debug().Raw(queryString).Where(whereClause).Scan(&usersExtended)

	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(usersExtended), "data": usersExtended})
}

func PutUser(c *gin.Context) {
	var body struct {
		ID             string
		Email          string
		Name           string
		FamilyName     string
		GivenName      string
		StartDate      string
		Observations   string
		ContactName    string
		ContactPhone   string
		ScheduleID     string
		SeleccionadoID string
		Birthday       string
		HasPicture     string
		PermissionID   string
		CategoryID     string
	}
	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	startDate, _ := time.ParseInLocation(time.RFC3339, body.StartDate, initializers.DB.NowFunc().Location())
	birthday, _ := time.ParseInLocation(time.RFC3339, body.Birthday, initializers.DB.NowFunc().Location())
	scheduleID, _ := uuid.Parse(body.ScheduleID)
	permissionID, _ := uuid.Parse(body.PermissionID)
	seleccionadoID, _ := uuid.Parse(body.SeleccionadoID)
	categoryID, _ := uuid.Parse(body.CategoryID)
	hasPicture, _ := strconv.Atoi(body.HasPicture)
	id, _ := uuid.Parse(body.ID)
	user := models.User{
		ID:             id,
		Email:          body.Email,
		Name:           body.Name,
		FamilyName:     body.FamilyName,
		GivenName:      body.GivenName,
		StartDate:      startDate,
		Observations:   body.Observations,
		ContactName:    body.ContactName,
		ContactPhone:   body.ContactPhone,
		ScheduleID:     scheduleID,
		SeleccionadoID: seleccionadoID,
		Birthday:       birthday,
		HasPicture:     hasPicture,
		PermissionID:   permissionID,
		CategoryID:     categoryID,
	}

	results := initializers.DB.Save(&user)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": results.RowsAffected, "data": user})

}

func PostUser(c *gin.Context) {

	var body struct {
		Email          string
		Name           string
		FamilyName     string
		GivenName      string
		StartDate      string
		Observations   string
		ContactName    string
		ContactPhone   string
		ScheduleID     string
		SeleccionadoID string
		Birthday       string
		HasPicture     string
		PermissionID   string
		CategoryID     string
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	startDate, _ := time.ParseInLocation(time.RFC3339, body.StartDate, initializers.DB.NowFunc().Location())
	birthday, _ := time.ParseInLocation(time.RFC3339, body.Birthday, initializers.DB.NowFunc().Location())
	scheduleID, _ := uuid.Parse(body.ScheduleID)
	permissionID, _ := uuid.Parse(body.PermissionID)
	seleccionadoID, _ := uuid.Parse(body.SeleccionadoID)
	categoryID, _ := uuid.Parse(body.CategoryID)
	hasPicture, _ := strconv.Atoi(body.HasPicture)
	user := models.User{
		Email:          body.Email,
		Name:           body.Name,
		FamilyName:     body.FamilyName,
		GivenName:      body.GivenName,
		StartDate:      startDate,
		Observations:   body.Observations,
		ContactName:    body.ContactName,
		ContactPhone:   body.ContactPhone,
		ScheduleID:     scheduleID,
		SeleccionadoID: seleccionadoID,
		Birthday:       birthday,
		HasPicture:     hasPicture,
		PermissionID:   permissionID,
		CategoryID:     categoryID,
	}

	result := initializers.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al crear usuario... ",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"ID": user.ID,
	})
}
