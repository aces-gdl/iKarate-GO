package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetGames(c *gin.Context) {
	/*
		var tournamentID = c.DefaultQuery("TournamentID", "")
		var categoryID = c.DefaultQuery("CategoryID", "")

		 results := initializers.DB.Raw(groupsQuery).Where("tournament_id = ? AND category_id = ?", tournamentID, categoryID).Scan(&groupsExtended)
		if results.Error != nil {
			c.JSON(http.StatusBadGateway, gin.H{"status": "error", "message": results.Error})
			return
		}

		c.JSON(http.StatusOK, gin.H{"status": "success", "results": len(groupsExtended), "data": groupsExtended})
	*/

}

type Game struct {
	Team1 int
	Team2 int
}

var rolOfGames []Game

func AddNewGame(team1 int, team2 int) {
	var gameNotFound = true
	for i := 0; i < len(rolOfGames); i++ {
		if (rolOfGames[i].Team1 == team1 && rolOfGames[i].Team2 == team2) || (rolOfGames[i].Team1 == team2 && rolOfGames[i].Team2 == team1) {
			gameNotFound = false
			break
		}
	}
	if gameNotFound {
		var newGame = Game{Team1: team1, Team2: team2}
		rolOfGames = append(rolOfGames, newGame)
	}
}

func CreateGames(TournamentID uuid.UUID, CategoryID uuid.UUID) {
	//	var teams = []string{"Cosmos", "Red Bull", "Manchester United", "Liverpool", "Staples", "WestHill", "Giants", "Foo Fighters", "Iphoners"}
	var teams = []string{"Cosmos", "Red Bull", "Manchester United", "Liverpool", "Staples", "WestHill", "Giants"}

	var numOfTeams int = len(teams)

	//var k int =0

	for i := 0; i < numOfTeams; i++ {
		for x := 0; x < numOfTeams; x++ {
			if i != x {
				AddNewGame(i, x)
			}
		}
	}

	for i := 0; i < len(rolOfGames); i++ {
		fmt.Println("Game ", i+1, " --> Team 1: ", teams[rolOfGames[i].Team1], " vs Team 2:", teams[rolOfGames[i].Team2])
	}
}

func PostAssignGamesToTimeSlots(c *gin.Context) {

}
