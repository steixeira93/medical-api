package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateAppointment(c *gin.Context) {
	request := CreateAppointmentRequest{}

	c.BindJSON(&request)

	// validate request
	if err := request.validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
}