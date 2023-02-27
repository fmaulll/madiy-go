package controllers

import (
	"net/http"

	"github.com/fmaulll/mandiy-go/initializers"
	"github.com/fmaulll/mandiy-go/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AddTransaction(context *gin.Context) {
	var body struct {
		CustomerId string `json:"customerId"`
		UserId     string `json:"userId"`
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read body"})

		return
	}

	uid := uuid.New()
	customerId := uuid.MustParse(body.CustomerId)
	userId := uuid.MustParse(body.UserId)

	transaction := models.Transaction{Id: uid, CustomerId: customerId, UserId: userId, Price: 10000}

	if err := initializers.DB.Create(&transaction).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create transaction!"})

		return
	}

	var totalTransaction []models.Transaction

	if err := initializers.DB.Where("customer_id = ?", customerId).Find(&totalTransaction).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to count transaction!"})

		return
	}

	if len(totalTransaction)%5 == 0 {
		var customer models.Customer

		if err := initializers.DB.Where("id = ?", customerId).Find(&customer).Update("point", customer.Point+1).Error; err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to add point!"})

			return
		}

		context.JSON(http.StatusCreated, gin.H{"message": "Succes added transaction, you gained 1 point!", "length": len(totalTransaction), "transactions": totalTransaction})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Succes added transaction"})
}

func UsePoint(context *gin.Context) {
	var body struct {
		CustomerId string `json:"customerId"`
		UserId     string `json:"userId"`
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read body"})

		return
	}

	var customer models.Customer

	uid := uuid.New()
	customerId := uuid.MustParse(body.CustomerId)
	userId := uuid.MustParse(body.UserId)

	transaction := models.Transaction{Id: uid, CustomerId: customerId, UserId: userId, Price: 0}

	if err := initializers.DB.Create(&transaction).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create transaction!"})

		return
	}

	if err := initializers.DB.Where("id = ?", customerId).Find(&customer).Update("point", customer.Point-1).Error; err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to use point!"})

		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Succes using point for transaction!"})
}
