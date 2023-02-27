package main

import (
	"os"

	"github.com/fmaulll/mandiy-go/controllers"
	"github.com/fmaulll/mandiy-go/initializers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.Migrate()
}

func main() {
	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"*"},
		AllowCredentials: true,
	}))

	router.POST("/signup", controllers.Signup)
	router.POST("/login", controllers.Login)
	router.GET("/customer", controllers.GetCustomer)
	router.POST("/customer", controllers.AddCustomer)
	router.POST("/transaction", controllers.AddTransaction)
	router.PATCH("/transaction", controllers.UsePoint)

	router.Run(":" + os.Getenv("PORT"))
}
