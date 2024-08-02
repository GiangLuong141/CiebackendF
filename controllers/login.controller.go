package controllers

import (
	"context"
	"example/configs"
	"example/models"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var adminCollection *mongo.Collection = configs.GetCollection(configs.DB, "admins")

func LoginAdmin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var loginRequest LoginRequest
		if err := c.ShouldBindJSON(&loginRequest); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		var admin models.Admin
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()
		err := adminCollection.FindOne(ctx, bson.M{"email": loginRequest.Email}).Decode(&admin)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid email or password",
			})
			return
		}

		err = bcrypt.CompareHashAndPassword([]byte(admin.Password), []byte(loginRequest.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"error": "Invalid email or password",
			})
			return
		}

		// Generate token here
		token, err := GenerateToken(admin.Email)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Failed to generate token",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"token":   token,
			"message": "Logged in successfully",
		})
	}
}

func GenerateToken(userEmail string) (string, error) {
	claims := jwt.MapClaims{
		"user_email": userEmail,
		"exp":        time.Now().Add(time.Hour * 24).Unix(), // Token expires in 24 hours
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("access_token_key"))
}
