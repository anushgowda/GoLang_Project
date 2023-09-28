package controllers

import (
	"fmt"
	"net/http"

	"github.com/anushgowda/GoLang_Project/entities"
	"github.com/anushgowda/GoLang_Project/interfaces"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService interfaces.IUser
}

func InitAuthController(authService interfaces.IUser) *AuthController {
	return &AuthController{AuthService: authService}
}



func (a *AuthController) Register(c *gin.Context) {
	fmt.Println("Invoked controller")
	var register entities.Register
	err := c.BindJSON(&register)
	if err != nil {
		fmt.Println("controller not invoked")
		return
	}
	result, err := a.AuthService.Register(&register)
	fmt.Println(result)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}

func (a *AuthController) Login(c *gin.Context) {
	fmt.Println("Invoked controller")
	var user entities.Login
	err := c.BindJSON(&user)
	if err != nil {
		return
	}
	result, err := a.AuthService.Login(&user)
	fmt.Println(result)
	if err != nil {
		return
	} else {
		c.IndentedJSON(http.StatusCreated, result)
	}
}

func (a *AuthController) Logout(c *gin.Context) {
	result := a.AuthService.Logout()
	c.IndentedJSON(http.StatusOK, result)
}