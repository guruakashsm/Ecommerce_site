package router

import (
	"ecommerce/controller"

	"github.com/gin-gonic/gin"
)

// Router creates and configures the Gin router.
func Router() *gin.Engine {
	router := gin.Default()
	router.Use(controller.CorsMiddleware())
	//Serve static files for specific routes
	router.Static("/index", "./frontend/index")
	router.Static("/home", "./frontend/home")
	router.Static("/signup", "./frontend/signup")
	router.Static("/signin", "./frontend/signin")
	//router.Static("/additems", "./frontend/inventory")
	router.Static("/inventory","./frontend/inven")
	router.Static("/items","./frontend/items")
	router.Static("/admin", "./frontend/admin")
	router.Static("/cart","./frontend/cart")
	//router.Static("/edit", "./frontend/edit")
	//router.Static("/delete", "./frontend/delete")
	router.Static("/seller", "./frontend/seller")
	router.Static("/feedback", "./frontend/feedback")
	//router.Static("/feedbacks","./frontend/feedbackReview")
	router.Static("/order", "./frontend/Payment")
    router.Static("/ordereditems","./frontend/order")

	// Define your routes
	router.GET("/getallcustomerdata", controller.Getalldata)
	router.GET("/getallinventorydata", controller.Getinventorydata)
	router.POST("/getallsellerdata", controller.Getallsellerdata)
	router.POST("/createseller", controller.CreateSeller)
	router.POST("/getdata", controller.GetData)
	router.POST("/create", controller.CreateProfile)
	router.POST("/deletedata", controller.Delete)
	router.POST("/addtocart", controller.Addtocart)
	router.POST("/verifyemail", controller.VerifyEmail)
	router.POST("/login", controller.Login)
	router.POST("/products", controller.Products)
	router.POST("/updatecart", controller.UpdateCart)
	router.POST("/inventory", controller.Inventory)
	router.POST("/search", controller.Search)
	router.POST("/update", controller.Update)
	router.GET("/inventorydata", controller.Getallinventorydata)
	router.POST("/sellercheck", controller.CheckSeller)
	router.POST("/deleteproduct", controller.DeleteProduct)
	router.POST("/deleteproductbyseller", controller.DeleteProductBySeller)
	router.POST("/updateproductbyseller", controller.UpdateProductBySeller)
	router.POST("/sitefeedback", controller.Feedback)
	router.POST("/getfeedback", controller.GetFeedback)
	router.GET("/sellerfeedback", controller.SellerFeedback)
	router.GET("/customerfeedback", controller.CustomerFeedback)
	router.POST("/deletefeedback", controller.Deletefeedback)
	router.POST("/getuser", controller.GetUser)
	router.POST("/totalamount", controller.TotalAmount)
	router.GET("/orders",controller.Orders)
	router.DELETE("/deleteorder",controller.DeleteOrder)
	router.POST("/customerorders",controller.CustomerOrder)
	router.POST("/validatetoken", controller.ValidateToken)
	router.POST("/adminlogin", controller.AdminLogin)
	router.POST("/adminpage", controller.GetAllDetailsForAdmin)
	router.POST("/getworkers", controller.GetWorkers)
	router.POST("/createworker",controller.CreateWorker)
	router.POST("/createadmin",controller.CreateAdmin)
	router.POST("/addevent",controller.AddEvent)
	router.POST("/getevent",controller.GetEvent)





	


	return router
}
