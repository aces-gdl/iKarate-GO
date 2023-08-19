package controllers

import (
	"fmt"
	"iKarate-GO/initializers"
	"iKarate-GO/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm/clause"
)

func GetSuscribed(c *gin.Context) {
}

func PostSimulateEnrollment(c *gin.Context) {
	var body struct {
		CategoryID   uuid.UUID
		UserCount    int
		TournamentID uuid.UUID
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}
	// Buscar Usuarios
	var users []models.User
	results := initializers.DB.Preload(clause.Associations).Limit(body.UserCount).Find(&users)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	var foundUsers int = int(results.RowsAffected)
	fmt.Println(foundUsers)

	// insertar en teams los 2 usuarios
	counter := len(users) / 2

	for i := 0; i < counter; i++ {
		var team models.TournamentTeam
		var teamMember1 models.User
		var teamMember2 models.User

		team.Member1ID = users[i+1].ID
		initializers.DB.First(&teamMember1, "id= ?", team.Member1ID)
		if teamMember1.ID.String() == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Usuario 1 no encontrado... ",
			})
			return
		}
		team.Member2ID = users[i+counter].ID
		initializers.DB.First(&teamMember2, "id= ?", team.Member2ID)
		if teamMember2.ID.String() == "" {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Usuario 2 no encontrado... ",
			})
			return
		}

		team.Name = fmt.Sprintf("Pareja  - %d", i+1)
		team.TournamentID = body.TournamentID
		team.CategoryID = body.CategoryID

		team.Name1 = teamMember1.Name
		team.Ranking1 = teamMember1.Ranking

		team.Name2 = teamMember2.Name
		team.Ranking2 = teamMember2.Ranking

		fmt.Println(team)
		result := initializers.DB.Create(&team)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Fallo al crear Equipo... ",
			})
			return
		}

	}
	c.JSON(http.StatusOK, gin.H{})
}

func PostCreateGroups(c *gin.Context) {

	// Round Robin Formula
	// Games = Teams (Teams -1) / 2

	const groupSize int = 3

	var body struct {
		CategoryID   uuid.UUID
		TournamentID uuid.UUID
	}

	if c.Bind(&body) != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Fallo al leer body...",
		})
		return
	}

	var tournament models.Tournament

	initializers.DB.First(&tournament, "id= ?", body.TournamentID)
	if tournament.ID.String() == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Torneo no existe... ",
		})
		return
	}

	var teams []models.TournamentTeam

	results := initializers.DB.Order("ranking1 + ranking2 DESC").Find(&teams)
	if results.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Usuarios no existen... ",
		})
		return
	}

	teamCounter := int(results.RowsAffected)

	groupsCounter := teamCounter / groupSize

	groupsCounterFinal := int(groupsCounter)
	if (teamCounter % groupSize) != 0 {
		groupsCounterFinal = int(groupsCounter) + 1
	}
	groups := make([]struct {
		ID      uuid.UUID
		counter int
	}, groupsCounterFinal)

	for i := 0; i < groupsCounterFinal; i++ {
		var group models.TournamentGroup
		fmt.Println("group ", i+1)
		group.Name = fmt.Sprintf("Grupo - %d", i+1)
		group.TournamentID = body.TournamentID
		group.CategoryID = body.CategoryID
		group.GroupNumber = i + 1
		result := initializers.DB.Create(&group)
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Fallo al crear grupo... ",
			})
			return
		}
		groups[i].ID = group.ID
		groups[i].counter = 1

	}
	goingUp := true
	teamSelector := 0
	groupSelector := 0
	teamCounterByGroup := 1
	for {

		fmt.Println("Group : ", groupSelector, "  --> Team : ", teamSelector+1)
		var teamByGroup models.TournamentTeamByGroup
		teamByGroup.TournamentID = body.TournamentID
		teamByGroup.CategoryID = body.CategoryID
		teamByGroup.GroupID = groups[groupSelector].ID
		teamByGroup.GroupNumber = groupSelector + 1
		teamByGroup.Name = fmt.Sprintf("Equipo : %d", teamCounterByGroup)
		teamByGroup.TeamID = teams[teamSelector].ID
		initializers.DB.Create(&teamByGroup)
		if goingUp {
			groupSelector++
		} else {
			groupSelector--
		}
		teamSelector++
		if groupSelector < 0 || groupSelector > groupsCounter-1 {
			goingUp = !goingUp
			if groupSelector < 0 {
				groupSelector = 0
			}
			if groupSelector > groupsCounter-1 {
				groupSelector = groupsCounter - 1
			}
			teamCounterByGroup++
			if teamCounterByGroup > groupSize {
				teamCounterByGroup = 1
			}
		}
		if teamSelector >= int(groupsCounter*groupSize) {
			break
		}
	}
	if teamSelector < int(teamCounter) {
		teamCounterByGroup = 1
		for i := teamSelector; i < int(teamCounter); i++ {
			fmt.Println(" Extra Group : ", groupsCounterFinal, "  --> Team : ", i+1)
			var teamByGroup models.TournamentTeamByGroup
			teamByGroup.TournamentID = body.TournamentID
			teamByGroup.CategoryID = body.CategoryID
			teamByGroup.GroupID = groups[groupsCounterFinal-1].ID
			teamByGroup.GroupNumber = groupsCounterFinal
			teamByGroup.Name = fmt.Sprintf("Equipo : %d", teamCounterByGroup)
			teamByGroup.TeamID = teams[i].ID
			initializers.DB.Create(&teamByGroup)
			teamCounterByGroup++
		}
	}
	// asignar equipos a grupos

}

// TODO :
//   - Persistir grupos
//   - Crear partidos
//   - Dar horarios a cada partido
//        - Crear Horarios disponibles como slots
//        - tomar un partido por grupo para cada slot ordenandolos por categoria

func GetEnrolledTeams(c *gin.Context) {
	var page = c.DefaultQuery("page", "1")
	var limit = c.DefaultQuery("limit", "10")
	var tournamentID = c.DefaultQuery("TournamentID", "")
	var categoryID = c.DefaultQuery("CategoryID", "")

	intPage, _ := strconv.Atoi(page)
	intLimit, _ := strconv.Atoi(limit)
	offset := (intPage - 1) * intLimit

	var teams []models.TournamentTeam
	results := initializers.DB.Where("tournament_id = ? AND category_id = ?", tournamentID, categoryID).Limit(intLimit).Offset(offset).Find(&teams)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(teams), "data": teams})

}

func GetGroups(c *gin.Context) {
	/*
		var page = c.DefaultQuery("page", "1")
		var limit = c.DefaultQuery("limit", "10")


			intPage, _ := strconv.Atoi(page)
			intLimit, _ := strconv.Atoi(limit)
			//offset := (intPage - 1) * intLimit
	*/
	var tournamentID = c.DefaultQuery("TournamentID", "")
	var categoryID = c.DefaultQuery("CategoryID", "")

	type groupExtended struct {
		TournamentID uuid.UUID
		CategoryID   uuid.UUID
		GroupID      uuid.UUID
		GroupNumber  int
		Name         string
		Name1        string
		Ranking1     int
		Name2        string
		Ranking2     int
		TeamRanking  int
	}
	var groupsExtended []groupExtended

	groupsQuery := `select 
					ttbg.tournament_id , 
					ttbg.category_id , 
					ttbg.group_id, 
					ttbg.group_number , 
					tt.name, 
					tt.name1, 
					tt.ranking1 , 
					tt.name2,
					tt.ranking2,
					tt.ranking1 + tt.ranking2 as team_ranking
					from tournament_team_by_groups ttbg inner join tournament_teams tt on ttbg.team_id  = tt.id
					where ttbg.category_id = '` + categoryID + `' and ttbg.tournament_id = '` + tournamentID + `'
					order by ttbg.group_number ASC, tt.ranking1 + ranking2 DESC`

	results := initializers.DB.Raw(groupsQuery).Where("tournament_id = ? AND category_id = ?", tournamentID, categoryID).Scan(&groupsExtended)
	if results.Error != nil {
		c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(groupsExtended), "data": groupsExtended})

}
