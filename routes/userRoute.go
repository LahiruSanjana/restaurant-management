package routes

import (
	"golang-restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine){
	incomingRoutes.GET("/users",controllers.GetAllUsers())
	incomingRoutes.GET("/users/:user_id",controllers.GetUserById())
	incomingRoutes.POST("/users/signup",controllers.SignUp())
	incomingRoutes.POST("/users/login",controllers.Login())
}