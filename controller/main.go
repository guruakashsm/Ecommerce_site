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

func Getinventorydata(c *gin.Context) {
	Inventorydata := service.Getinventorydata()
	c.JSON(http.StatusOK, Inventorydata)
}
func Getalldata(c *gin.Context) {
	alltransaction := service.Getalldata()
	c.JSON(http.StatusOK, alltransaction)
}
func CreateSeller(c *gin.Context) {
	var seller models.Seller
	if err := c.BindJSON(&seller); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	createseller := service.CreateSeller(seller)
	c.JSON(http.StatusOK, createseller)

}
func Getallsellerdata(c *gin.Context) {
	Getallsellerdata := service.Getallsellerdata()
	c.JSON(http.StatusOK, Getallsellerdata)

}
func UpdateCart(c *gin.Context) {
	var cart models.Cart
	if err := c.BindJSON(&cart); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	t1, err := service.ExtractCustomerID(cart.CustomerId, constants.SecretKey)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	cart.CustomerId = t1

	result := service.UpdateCart(cart)
	c.JSON(http.StatusOK, result)
}
func Login(c *gin.Context) {
	var request models.Login
	if err := c.BindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token, success, err := service.Login(request)
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

func Products(c *gin.Context) {
	var cartproducts *models.Addtocart
	if err := c.BindJSON(&cartproducts); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token, err := service.ExtractCustomerID(cartproducts.Token, constants.SecretKey)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	cart := service.Cart(token)
	c.JSON(http.StatusOK, cart)
}

func Addtocart(c *gin.Context) {
	var addtocart models.Addtocart
	if err := c.BindJSON(&addtocart); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	token, err := service.ExtractCustomerID(addtocart.Token, constants.SecretKey)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Token"})
		return
	}
	var addtocart1 models.Addtocart1
	addtocart1.CustomerId = token
	addtocart1.Name = addtocart.Name
	addtocart1.Price = addtocart.Price
	addtocart1.SellerQuantity = addtocart.Sellerquantity
	result := service.Addtocart(addtocart1)
	c.JSON(http.StatusOK, result)

}

func CreateProfile(c *gin.Context) {
	var profile models.Customer
	if err := c.BindJSON(&profile); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Insert(profile)
	c.JSON(http.StatusOK, result)
}
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
func Getallinventorydata(c *gin.Context) {
	result := service.Search(SearchName)
	c.JSON(http.StatusOK, result)
}

var SearchName string

func Search(c *gin.Context) {
	type Serarch struct {
		ProductName string `json:"productName" bson:"productName"`
	}
	var search Serarch
	if err := c.BindJSON(&search); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	SearchName = search.ProductName
}

func Update(c *gin.Context) {
	var update models.Update
	if err := c.BindJSON(&update); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Update(update)
	c.JSON(http.StatusOK, result)
}

func Delete(c *gin.Context) {
	var delete models.Delete
	if err := c.BindJSON(&delete); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Delete(delete)
	c.JSON(http.StatusOK, result)

}

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

func DeleteProduct(c *gin.Context) {
	var delete models.DeleteProduct
	if err := c.BindJSON(&delete); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	fmt.Println(delete)
	result := service.DeleteProduct(delete)
	c.JSON(http.StatusOK, result)

}

func DeleteProductBySeller(c *gin.Context) {
	var delete models.DeleteBySeller
	if err := c.BindJSON(&delete); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.DeleteProductBySeller(delete)
	c.JSON(http.StatusOK, result)

}

func UpdateProductBySeller(c *gin.Context) {
	var update models.UpdateProduct
	if err := c.BindJSON(&update); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.UpdateProductBySeller(update)
	c.JSON(http.StatusOK, result)
}

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

func Deletefeedback(c *gin.Context) {
	var feedback models.FeedbacktoAdmin

	if err := c.BindJSON(&feedback); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	result := service.Deletefeedback(feedback)
	c.JSON(http.StatusOK, result)

}





func GetUser(c *gin.Context) {
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
	
	service.Buynow(id)
	c.JSON(http.StatusOK, data)
}

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

func Orders(c *gin.Context) {
	data := service.Orders()
	c.JSON(http.StatusOK, data)
}

func DeleteOrder(c *gin.Context) {
	var delete models.DeleteOrder
	if err := c.BindJSON(&delete); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}
	service.DeleteOrder(delete)
	c.JSON(http.StatusOK, gin.H{"success": true})

}
func CustomerOrder(c *gin.Context) {
	var token models.Token
	if err := c.BindJSON(&token); err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON data"})
		return
	}

	data := service.CustomerOrder(token.Token)
	c.JSON(http.StatusOK, data)

}


