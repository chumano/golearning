package api

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Username string `json:"username" binding:"required,min=5,max=30"`
	Password string `json:"password" validate:"required,min=5,max=20"`
}

var Users []User = []User{
	User{
		Username: "loc",
		Password: "chumano",
	},
}

func SignIn(ctx *gin.Context) {
	user := new(User)

	if err := ctx.Bind(user); err != nil {

		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	//https://dev.to/kittipat1413/a-guide-to-input-validation-in-go-with-validator-v10-56bp
	// Create a new validator instance
	validate := validator.New()

	// Validate the User struct
	err := validate.Struct(user)
	validationErrors := err.(validator.ValidationErrors)
	log.Printf("validationErrors: %v", validationErrors)

	log.Printf("user: %v", user)

	for _, u := range Users {
		if u.Username == user.Username && u.Password == user.Password {
			ctx.JSON(http.StatusOK, gin.H{
				"msg": "Signed in successfully.",
				"jwt": "123456789",
			})
			return
		}
	}
	ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"err": "Sign in failed.", "data": user})
}
