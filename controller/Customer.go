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

// Signup Function
func CreateProfile(c *gin.Context) {
	var profile models.Customer
	if err := c.BindJSON(&profile); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.CreateCustomer(profile)
	c.JSON(http.StatusOK, result)
}

// Customer Email Verification
func VerifyEmail(c *gin.Context) {
	var Data models.VerifyEmail
	if err := c.BindJSON(&Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	data, err := service.EmailVerification(Data)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Internal Server Error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})
}


// Signin Function
func Login(c *gin.Context) {
	var request models.Login
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token, no,err := service.Login(request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal Server Error"})
		return
	}
	if no == 1 {
		c.JSON(http.StatusOK, gin.H{"token": token})
		return
	} else if no == 0 {
		c.JSON(http.StatusOK, gin.H{"message": token})
		return
	}
}

// Validate Customer Token
func ValidateToken(c *gin.Context) {
	var userdata models.Token
	if err := c.BindJSON(&userdata); err != nil {

		log.Fatal(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Validatetoken(userdata.Token,constants.SecretKey)
	c.JSON(http.StatusOK, gin.H{"result": result})
}

// Search Products
func Search(c *gin.Context) {
	var search models.Search
	if err := c.BindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Search(search.ProductName)
	c.JSON(http.StatusOK, gin.H{"data": result})

}

// Add In Customer Cart
func Addtocart(c *gin.Context) {
	var addtocart models.Addtocart
	if err := c.BindJSON(&addtocart); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result, err := service.Addtocart(addtocart)
	if err != nil {
		log.Println(result, err)
		c.JSON(http.StatusOK, gin.H{"error": result})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": result})

}

// List Product In Customer Cart
func Products(c *gin.Context) {
	var cartproducts models.Token
	if err := c.BindJSON(&cartproducts); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	cart, message, err := service.GetAllItemsinCart(cartproducts)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": message})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": cart})
}

// Updating Customer Cart
func UpdateCart(c *gin.Context) {
	var cart models.Addcart
	if err := c.BindJSON(&cart); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	log.Println(cart)
	result,message := service.UpdateCart(cart)
	if !result{
		c.JSON(http.StatusOK, gin.H{"error":message})
		return
	}
	c.JSON(http.StatusOK,gin.H{"message":message})
}

// Delete Items In Cart
func DeleteProduct(c *gin.Context) {
	var delete models.DeleteProduct
	if err := c.BindJSON(&delete); err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.DeleteProduct(delete)
	c.JSON(http.StatusOK, result)

}

// Cart To Buy Items
func BuyNow(c *gin.Context) {
	var data models.Address
	var token models.Token
	var ItemsToBuy []models.Item
	if err := c.BindJSON(&token); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(token.Token)
	id, _ := service.ExtractCustomerID(token.Token, constants.SecretKey)

	data = service.GetUser(id)
	ItemsToBuy = service.Itemstobuy(id)
	fmt.Println(ItemsToBuy)
	service.CustomerOrders(ItemsToBuy, data)

	service.DeleteItemsInCart(id)
	c.JSON(http.StatusOK, data)
}

//Display Total Amount
func TotalAmount(c *gin.Context) {
	var data models.TotalAmount
	var token models.Token
	if err := c.BindJSON(&token); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(token.Token)
	id, _ := service.ExtractCustomerID(token.Token, constants.SecretKey)
	data.TotalAmount = service.TotalAmount(id)
	c.JSON(http.StatusOK, data)
}

//Display Customer Orders
func CustomerOrder(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	data := service.CustomerOrder(token.Token)
	c.JSON(http.StatusOK, data)

}


//Get Items Based On Search
func GetInventoryData(c *gin.Context) {
	var search models.Search
	if err := c.BindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	log.Println(search)
	data, err := service.FetchInventoryData(search)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"error": "No result Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": data})

}