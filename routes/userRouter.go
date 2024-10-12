package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/krnveersharma/golang-jwt-project/controllers"
	"github.com/krnveersharma/golang-jwt-project/middleware"
)

func userRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.POST("users", controllers.GetUsers())
	incomingRoutes.POST("users/:user_id", controllers.GetUser())
}
