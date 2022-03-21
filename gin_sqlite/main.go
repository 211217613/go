package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	gin.SetMode(gin.DebugMode)
	// API v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("person", getPersons)
		// v1.GET("person/:id", getPersonById)
		// v1.POST("person", addPerson)
		// v1.PUT("person/:id", updatePerson)
		// v1.DELETE("person/:id", deletePerson)
		// v1.OPTIONS("person", options)
	}
	r.Run()
}

func getPersons(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "getPersons Called"})
}
