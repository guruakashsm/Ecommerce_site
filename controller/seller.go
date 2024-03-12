package controller

import (
	"ecommerce/constants"
	"ecommerce/models"
	"ecommerce/service"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Seller Login
func CheckSeller(c *gin.Context) {
	var check models.Login
	if err := c.BindJSON(&check); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	token, success, err := service.CheckSeller(check)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if success {
		c.JSON(http.StatusOK, gin.H{"token": token})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}

}

// Add Items to inventory
func Inventory(c *gin.Context) {
	var inventory models.Inventory
	if err := c.BindJSON(&inventory); err != nil {

		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(inventory)
	sellerid, err := service.ExtractCustomerID(inventory.SellerId, constants.SecretKey)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token : while Extracting"})
	}
	inventory.SellerId = sellerid
	success, err := service.Inventory(inventory)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if success {
		c.JSON(http.StatusOK, gin.H{"data": success})
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
	}

}

// Update Items in Inventory
func UpdateProductBySeller(c *gin.Context) {
	var update models.UpdateProduct
	if err := c.BindJSON(&update); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.UpdateProductBySeller(update)
	c.JSON(http.StatusOK, result)
}

// Delete Items in Inventory
func DeleteProductBySeller(c *gin.Context) {
	var delete models.DeleteBySeller
	if err := c.BindJSON(&delete); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.DeleteProductBySeller(delete)
	c.JSON(http.StatusOK, result)

}

//Display Customer Order
func Orders(c *gin.Context) {
	data := service.Orders()
	c.JSON(http.StatusOK, data)
}


//Delete Order
func DeleteOrder(c *gin.Context) {
	var delete models.DeleteOrder
	if err := c.BindJSON(&delete); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	service.DeleteOrder(delete)
	c.JSON(http.StatusOK, gin.H{"success": true})

}
