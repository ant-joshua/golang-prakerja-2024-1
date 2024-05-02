package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"todo-gin/pkg/domains/models"
	"todo-gin/pkg/helpers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type UserController struct {
	db *gorm.DB
}

func NewUserController(db *gorm.DB) *UserController {
	return &UserController{
		db: db,
	}
}

func (u *UserController) Routes(g *gin.RouterGroup, customMiddleware ...gin.HandlerFunc) {

	userGroup := g.Group("/users")
	userGroup.Use(customMiddleware...)

	userGroup.GET("", u.GetAllUser)
	userGroup.GET("/:id", u.GetUser)
	userGroup.POST("", u.CreateUser)
	userGroup.PUT("/:id", u.UpdateUser)
	userGroup.DELETE("/:id", u.DeleteUser)
}

func (u *UserController) GetAllUser(c *gin.Context) {

	auth := c.MustGet("user").(*helpers.JwtCustomClaims)

	fmt.Printf("authUser: %+v\n", auth)

	var users []models.User

	err := u.db.Find(&users).Error

	if err != nil {
		log.Printf("Error when get all user: %v", err)

		c.JSON(500, gin.H{
			"statusCode": http.StatusInternalServerError,
			"message":    "Internal Server Error",
			"data":       nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"statusCode": http.StatusOK,
		"message":    "Success",
		"data":       users,
	})
}

func (u *UserController) GetUser(c *gin.Context) {
	// get param id
	id := c.Param("id")
	parseId, err := strconv.Atoi(id)

	if err != nil {
		BadRequestResponse(c, nil)
		return
	}
	var user *models.User

	// err = u.db.Where("id = ?", parseId).First(&user).Error

	err = u.db.First(&user, parseId).Error

	if err != nil {
		InternalServerErrorResponse(c)
		return
	}

	if user == nil {
		NotFoundResponse(c)
		return
	}

	SuccessResponse(c, user)
}

func (u *UserController) CreateUser(c *gin.Context) {
	var createUserRequest models.CreateUserRequest

	if err := c.ShouldBindJSON(&createUserRequest); err != nil {
		BadRequestResponse(c, nil)
		return
	}

	newUser := models.User{
		Name:     createUserRequest.Name,
		Username: createUserRequest.Username,
		Email:    createUserRequest.Email,
		Password: createUserRequest.Password,
	}

	err := u.db.Create(&newUser).Error

	if err != nil {
		log.Printf("Error when create user: %v", err)

		InternalServerErrorResponse(c)
		return
	}

	SuccessResponse(c, newUser)
}

func (u *UserController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	parseId, err := strconv.Atoi(id)

	if err != nil {
		BadRequestResponse(c, nil)
		return
	}

	var user *models.User

	err = u.db.First(&user, parseId).Error

	if err != nil {
		InternalServerErrorResponse(c)
		return
	}

	if user == nil {
		NotFoundResponse(c)
		return
	}

	err = u.db.Delete(&user).Error

	if err != nil {
		InternalServerErrorResponse(c)
		return
	}

	SuccessResponse(c, nil)
}

func (u *UserController) UpdateUser(c *gin.Context) {
	var updateUserRequest models.UpdateUserRequest

	id := c.Param("id")
	parseId, err := strconv.Atoi(id)

	if err != nil {
		BadRequestResponse(c, nil)
		return
	}

	if err := c.ShouldBindJSON(&updateUserRequest); err != nil {
		BadRequestResponse(c, nil)
		return
	}

	var user *models.User

	err = u.db.First(&user, parseId).Error

	if err != nil {
		InternalServerErrorResponse(c)
		return
	}

	if user == nil {
		NotFoundResponse(c)
		return
	}

	if updateUserRequest.Username != nil {
		user.Username = *updateUserRequest.Username
	}

	if updateUserRequest.Name != nil {
		user.Name = *updateUserRequest.Name
	}

	if updateUserRequest.Email != nil {
		user.Email = *updateUserRequest.Email
	}

	err = u.db.Save(&user).Error

	if err != nil {
		log.Printf("Error when create user: %v", err)

		InternalServerErrorResponse(c)
		return
	}

	SuccessResponse(c, user)
}
