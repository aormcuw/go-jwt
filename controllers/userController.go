package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/aormcuw/go-jwt/initializer"
	"github.com/aormcuw/go-jwt/models"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

func Signup(ctx *gin.Context) {
	// Get the email/password off req body
	var body struct {
		Email    string
		Password string
	}

	if err := ctx.Bind(&body); err != nil {

	}

	// Hash the password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), 10)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to hash password",
		})

		return
	}

	// Create user
	user := models.User{
		Email:    body.Email,
		Password: string(hash),
	}
	result := initializer.DB.Create(&user)
	if result.Error != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create user",
		})

		return
	}

	// Respond

	ctx.JSON(200, gin.H{})

}

func Login(ctx *gin.Context) {
	// Get the email/password off req body
	var body struct {
		Email    string
		Password string
	}

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "email or password incorrect",
		})

		return
	}

	// Check is the email or correct
	var user models.User
	initializer.DB.First(&user, "email = ?", body.Email)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "email or password incorrect",
		})

		return
	}

	// Compare the password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "email or password incorrect",
		})

		return
	}

	// Generate JWT token
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "failed to create token",
		})

		return
	}

	// send it back
	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)
	ctx.JSON(200, gin.H{})
}

func Validate(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"text": "i'm logged in",
	})
}
