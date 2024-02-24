package controllers

import (
	"chat_app/db"
	"chat_app/models"
	"chat_app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *services.UserService
}

func NewUserController(DatabaseClient *db.DbClient) *UserController {
	return &UserController{
		userService: services.NewUserService(DatabaseClient),
	}
}

func (userController *UserController) GetUser(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, "Could not parse the email param")
		return
	}
	user, err := userController.userService.FindUser(c, email)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	c.JSON(http.StatusAccepted, user)
}

func (userController *UserController) PostLogout(c *gin.Context) {
	c.HTML(http.StatusAccepted, "login-modal.html", nil)
}

func (userController *UserController) GetLogin(c *gin.Context) {
	c.HTML(http.StatusAccepted, "login-modal.html", nil)
}

func (userController *UserController) PostLogin(c *gin.Context) {
	var cred models.UserCredentials
	if err := c.ShouldBindJSON(&cred); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	c.Writer.WriteHeader(http.StatusAccepted)
}

func (userController *UserController) GetSignup(c *gin.Context) {
	c.HTML(http.StatusAccepted, "signup-modal.html", nil)
}

func (userController *UserController) PostSignup(c *gin.Context) {

	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	value, err := userController.userService.CreateUser(c, &newUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusAccepted, value)
}
