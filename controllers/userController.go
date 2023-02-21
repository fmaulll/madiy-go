package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/fmaulll/mandiy-go/initializers"
	"github.com/fmaulll/mandiy-go/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

func Signup(context *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read body"})

		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to has password "})

		return

	}
	user := models.User{Username: body.Username, Password: string(hash)}

	result := initializers.DB.Create(&user)

	if result.Error != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create user"})

		return

	}

	context.JSON(http.StatusCreated, gin.H{"message": "User " + body.Username + " created"})
}

func Login(context *gin.Context) {
	var body struct {
		Username string
		Password string
	}

	if context.Bind(&body) != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to read body"})

		return
	}
	var user models.User
	initializers.DB.First(&user, "username = ?", body.Username)

	if user.ID == 0 {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})

		return
	}

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid email or password"})

		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Failed to create token"})

		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":  "Login successfully",
		"token":    tokenString,
		"id":       user.ID,
		"username": body.Username,
	})
}
