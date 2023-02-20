package controllers

import (
	"net/http"

	"github.com/fmaulll/mandiy-go/initializers"
	"github.com/fmaulll/mandiy-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetCustomer(context *gin.Context) {
	phone := context.Query("phone")

	var customer models.Customer

	initializers.DB.Where("phone = ?", phone).Find(&customer)

	if customer.Id == uuid.MustParse("00000000-0000-0000-0000-000000000000") {
		context.JSON(http.StatusNotFound, gin.H{"message": "Customer not found in database!"})

		return
	}

	context.JSON(http.StatusOK, gin.H{"customer": customer})
}

func AddCustomer(context *gin.Context) {
	var body struct {
		Email     string `json:"email"`
		Phone     string `json:"phone"`
		FirstName string `json:"firstName"`
		LastName  string `json:"lastName"`
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read body"})

		return
	}

	uuid := uuid.New()

	customer := models.Customer{Id: uuid, Email: body.Email, Phone: body.Phone, FirstName: body.FirstName, LastName: body.LastName}

	if err := initializers.DB.Create(&customer).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create customer!"})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Success added user " + body.Phone})
}
