package main

import (
	"fmt"
	"net/http"

	"github.com/Kratos-28/db"
	"github.com/Kratos-28/models"
	"github.com/gin-gonic/gin"
)

func main() {
	err := db.InitDB()
	if err != nil {
		fmt.Errorf("Could not establist DB connecttion %v", err)
	}
	server := gin.Default()          //Give Engine
	server.GET("/events", getEvents) //GET,POST,PUT,PATCH,DELETE
	server.POST("/events", createEvent)
	server.Run(":8080") //localhost:8080

}

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not fetch the events.Try Again Later",
		})

	}
	context.JSON(http.StatusOK, events)

}

func createEvent(context *gin.Context) {
	var events models.Event
	err := context.ShouldBindJSON(&events)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parsed",
		})
	}

	events.ID = 1
	events.UserID = 1

	err = events.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not create the event.Try Again Later",
		})

	}

	context.JSON(http.StatusCreated, gin.H{
		"message": "Event Created",
		"event":   events,
	})

}
