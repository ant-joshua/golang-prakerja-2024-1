package models

import "github.com/gin-gonic/gin"

type BaseResponse struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
	Data       any    `json:"data"`
	Errors     any    `json:"errors"`
}

func SuccessResponse(c *gin.Context, data any) {
	c.JSON(200, BaseResponse{
		StatusCode: 200,
		Message:    "Success",
		Data:       data,
	})
}

func BadRequestResponse(c *gin.Context, errors any) {
	c.JSON(400, BaseResponse{
		StatusCode: 400,
		Message:    "Bad Request",
		Errors:     errors,
	})
}

func NotFoundResponse(c *gin.Context) {
	c.JSON(404, BaseResponse{
		StatusCode: 404,
		Message:    "Not Found",
		Errors:     nil,
	})
}

func UnauthenticatedResponse(c *gin.Context) {
	c.JSON(401, BaseResponse{
		StatusCode: 401,
		Message:    "Unauthenticated",
		Errors:     nil,
	})
}

func InternalServerErrorResponse(c *gin.Context) {
	c.JSON(500, BaseResponse{
		StatusCode: 500,
		Message:    "Internal Server Error",
		Errors:     nil,
	})
}
