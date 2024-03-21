package controller

import (
	"ecommerce/models"
	"ecommerce/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func InsertSellerFeedback(c *gin.Context) {
	var feedback models.Feedback

	if err := c.BindJSON(&feedback); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result,err := service.InstertSellerFeedback(feedback)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error":result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":result})
}

func InsertCustomerFeedback(c *gin.Context) {
	var feedback models.Feedback

	if err := c.BindJSON(&feedback); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result,err := service.InsertCustomerFeedback(feedback)
	if err != nil{
		c.JSON(http.StatusOK, gin.H{"error":result})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message":result})
}

func SellerFeedback(c *gin.Context) {
	result := service.SellerFeedback()
	c.JSON(http.StatusOK, result)
}

func CustomerFeedback(c *gin.Context) {
	result := service.CustomerFeedback()
	c.JSON(http.StatusOK, result)
}
