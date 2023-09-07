package controllers

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func PostCreateTimeSlots(c *gin.Context) {

	var dailyFirstGameStart = time.Date(2023, 9, 6, 17, 0, 0, 0, time.Local)
	var dailyLastGameStart = time.Date(2023, 9, 6, 22, 0, 0, 0, time.Local)
	var gameDuration = 60 * time.Minute // in minutes
	var availableCourtCount = 3
	/*

		var roundRobinDays = 4

	*/
	var gamesPerDayPerCourt = int(dailyLastGameStart.Sub(dailyFirstGameStart) / gameDuration)
	var gamesPerDay = gamesPerDayPerCourt * availableCourtCount
	var roundrobinAvailableSlots = gamesPerDay * availableCourtCount

	fmt.Println("Games per court per day ", int(dailyLastGameStart.Sub(dailyFirstGameStart)/gameDuration))
	fmt.Println("Games per day", gamesPerDay)
	fmt.Println("Total round robin available slots", roundrobinAvailableSlots)

	/*

		for i := 0; i < 5; i++ {
			fmt.Println(startDate)
			startDate = startDate.Add(60 * time.Minute)
		}

			var body struct {
				TournamentID uuid.UUID
			}

			if c.Bind(&body) != nil {
				c.JSON(http.StatusBadRequest, gin.H{
					"error": "Fallo al leer body...",
				})
				return
			}

			var tournament models.Tournament

			result := initializers.DB.Where("ID = ?", body.TournamentID).First(&tournament)

			if result.Error != nil {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Torneo no encontrado...",
				})
				return

			}
			fmt.Println(tournament)
	*/
}
