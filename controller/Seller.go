package controller

import (
	"ecommerce/models"
	"ecommerce/service"
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
		return
	}
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": token})
		return
	}
	if !success {
		c.JSON(http.StatusOK, gin.H{"message": token})
	}

}

// Add Items to inventory
func Inventory(c *gin.Context) {
	var inventory models.Inventory
	if err := c.BindJSON(&inventory); err != nil {

		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	message, err := service.Inventory(inventory)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})

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

// Display Customer All Orders
func Orders(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	data, message, err := service.Orders(token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})

}

// Display Customer Completed Orders
func CompletedOrders(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	data, message, err := service.CompletedOrders(token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})

}

// Display Customer Pending Orders
func PendingOrders(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	data, message, err := service.PendingOrders(token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})
}

// Yet to Deliver Orders
func YettoDeliverOrders(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	data, message, err := service.YettoDeliverOrders(token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	log.Println(data)
	c.JSON(http.StatusOK, gin.H{"message": data})
}

// Delete Order
func DeleteOrder(c *gin.Context) {
	var delete models.DeleteOrder
	if err := c.BindJSON(&delete); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	message, err := service.DeleteOrder(delete)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})

}

// Signup Seller
func RegisterSeller(c *gin.Context) {
	var register models.Seller
	if err := c.BindJSON(&register); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	message, err := service.RegisterSeller(register)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// Verify Seller Email
func VerifySellerEmail(c *gin.Context) {
	var register models.VerifyEmail
	if err := c.BindJSON(&register); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	message, err := service.EmailVerificationforSeller(register)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}

// Data needed for Seller DrashBoard
func SellerDrashbordDetails(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	data, message, err := service.SellerDrashbordDetails(token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})
}

// Get All Products of seller
func GetAllProducts(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	data, message, err := service.GetAllProducts(token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})
}

// List BuyedCustomer
func BuyedCustomer(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	data, message, err := service.GetAllBuyedCustomer(token)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})
}

// Get Customer Order
func GetOrderbySeller(c *gin.Context) {
	var details models.GetOrder
	if err := c.BindJSON(&details); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	data, message, err := service.GetCustromerOrder(details)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})
}


// Update OrderTracking Details
func UpdateOrderTracking(c *gin.Context) {
	var details models.OrderTracking
	if err := c.BindJSON(&details); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	message, err := service.UpdateOrderTracking(details)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusOK, gin.H{"error": message})
	}
	c.JSON(http.StatusOK, gin.H{"message": message})
}


