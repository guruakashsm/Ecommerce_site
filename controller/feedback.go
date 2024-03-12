package controller

import (
	"ecommerce/models"
	"ecommerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Feedback(c *gin.Context) {
	var feedback models.Feedback

	if err := c.BindJSON(&feedback); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Feedback(feedback)
	c.JSON(http.StatusOK, result)
}

func SellerFeedback(c *gin.Context) {
	result := service.SellerFeedback()
	c.JSON(http.StatusOK, result)
}

func CustomerFeedback(c *gin.Context) {
	result := service.CustomerFeedback()
	c.JSON(http.StatusOK, result)
}
