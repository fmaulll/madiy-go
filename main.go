package main

import (
	"os"

	"github.com/fmaulll/mandiy-go/controllers"
	"github.com/fmaulll/mandiy-go/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDB()
	initializers.Migrate()
}

func main() {
	router := gin.Default()

	router.GET("/customer", controllers.GetCustomer)
	router.POST("/customer", controllers.AddCustomer)

	router.Run(":" + os.Getenv("PORT"))
}
