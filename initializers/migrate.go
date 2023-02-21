package initializers

import "github.com/fmaulll/mandiy-go/models"

func Migrate() {
	DB.AutoMigrate(&models.Customer{})
	DB.AutoMigrate(&models.Transaction{})
}
