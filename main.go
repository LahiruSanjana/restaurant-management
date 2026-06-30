package main

import (
	"os"
	"golang-restaurant-management/database"
	"golang-restaurant-management/routes"
	"golang-restaurant-management/middleware"
	"github.com/gin-gonic/gin"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main(){
	port := os.Getenv("PORT")

	if port == ""{
		port = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.UserRoutes(router)
	router.Use(middleware.Authentication())

	routes.FoodRoutes(router)
	routes.OrderRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrederItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}