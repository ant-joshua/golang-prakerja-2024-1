package controllers

import (
	"log"
	"todo-gin/pkg/domains/models"
	"todo-gin/pkg/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type AuthController struct {
	db *gorm.DB
}

func NewAuthController(db *gorm.DB) *AuthController {
	return &AuthController{
		db: db,
	}
}

func (a *AuthController) Routes(g *gin.RouterGroup) {
	g.POST("/login", a.Login)
	g.POST("/register", a.Register)
}

func (a *AuthController) Login(c *gin.Context) {
	// Do something here

	var loginRequest models.LoginRequest

	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		models.BadRequestResponse(c, nil)
		return
	}

	var user models.User

	err := a.db.Where("username = ?", loginRequest.Username).
		Or("email = ?", loginRequest.Username).
		First(&user).Error

	if err != nil {
		log.Printf("Error when get user: %v", err)

		models.NotFoundResponse(c)
		return
	}

	if ok, _ := helpers.ComparePassword(user.Password, loginRequest.Password); !ok {
		log.Printf("Password not match")
		models.UnauthenticatedResponse(c)
		return
	}

	token, err := helpers.GenerateJwtToken(user.Username, user.Email, user.ID)

	if err != nil {
		models.InternalServerErrorResponse(c)
		return
	}

	resp := models.LoginResponse{
		AccessToken: token,
		Username:    user.Username,
	}

	models.SuccessResponse(c, resp)

}

func (a *AuthController) Register(c *gin.Context) {
	// Do something here
	var createUserRequest models.RegisterRequest

	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		models.BadRequestResponse(c, nil)
		return
	}

	newUser := models.User{
		Name:     createUserRequest.Name,
		Username: createUserRequest.Username,
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	err := a.db.Create(&newUser).Error

	if err != nil {
		log.Printf("Error when create user: %v", err)

		models.InternalServerErrorResponse(c)
		return
	}

	resp := models.RegisterResponse{
		ID:       newUser.ID,
		Username: newUser.Username,
		Name:     newUser.Name,
		Email:    newUser.Email,
	}

	models.SuccessResponse(c, resp)

}
