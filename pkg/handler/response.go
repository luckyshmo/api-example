package handler

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type errorResponse struct {
	Message string `json:"message"`
}

type statusResponse struct {
	Status string `json:"status"`
}

func sendErrorResponse(c *gin.Context, statusCode int, message string) {
	logrus.Error(message)
	c.AbortWithStatusJSON(statusCode, errorResponse{message})
}

func sendStatusResponse(c *gin.Context, statusCode int, i interface{}) { //? Is using interface a good idea?
	logrus.Print(fmt.Sprintf("TEST RESP PRINT %s", i)) //TODO check lvls
	c.JSON(statusCode, i)
}
