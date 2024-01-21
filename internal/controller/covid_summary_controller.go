package controller

import (
	"net/http"

	"github.com/GraphZC/lmwn-jr-2024/internal/service"
	"github.com/gin-gonic/gin"
)

func GetCovidSummary(c *gin.Context) {
	// Get covid cases 
	covidCases, err := service.GetCovidCases()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	// Get covid summary
	covidSummary := service.GetCovidSummary(covidCases)

	c.JSON(http.StatusOK,
		covidSummary,
	)
}