package router

import (
	"github.com/GraphZC/lmwn-jr-2024/internal/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	r.GET("/covid/summary", controller.GetCovidSummary)
}