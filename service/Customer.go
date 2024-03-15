package service

import (
	"context"
	"ecommerce/config"
	"ecommerce/constants"
	"ecommerce/models"
	"fmt"
	"log"
	"regexp"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// Create Customer
func CreateCustomer(profile models.Customer) int {
	name := profile.Name
	regexPattern := "^[a-zA-Z ]+$"
	regexpObject := regexp.MustCompile(regexPattern)
	match := regexpObject.FindString(name)

	if match == "" {
		return 4
	}

	phonecount := countdigits(profile.Phone_No)

	if phonecount != 10 {
		return 5
	}

	if profile.Password != profile.ConfirmPassword {
		return 3
	}

	filter := bson.M{"email": profile.Email}
	existingCustomer := config.Customer_Collection.FindOne(context.Background(), filter)

	// Check if the document already exists
	if existingCustomer.Err() == nil {
		var existingProfile models.Customer // Replace with the actual type of your document
		if err := existingCustomer.Decode(&existingProfile); err != nil {
			log.Println(err)
			return 0
		}

		// Check if the email is verified
		if existingProfile.IsEmailVerified {
			return 2
		} else {
			profile.CustomerId = GenerateUniqueCustomerID()
			profile.IsEmailVerified = false
			profile.WrongInput = 0
			profile.VerificationString = GenerateOTP(6)
			_, updateErr := config.Customer_Collection.ReplaceOne(context.Background(), filter, profile)
			if updateErr != nil {
				return 0
			}
			go SendEmailforCustomerVerification(profile.Email, profile.VerificationString, profile.Name)

			return 1
		}
	} else if existingCustomer.Err() == mongo.ErrNoDocuments {

		profile.CustomerId = GenerateUniqueCustomerID()
		profile.IsEmailVerified = false
		profile.BlockedUser = false
		profile.WrongInput = 0
		profile.VerificationString = GenerateOTP(6)

		inserted, insertErr := config.Customer_Collection.InsertOne(context.Background(), profile)
		if insertErr != nil {
			log.Println(insertErr)
			return 0
		}
		go SendEmailforCustomerVerification(profile.Email, profile.VerificationString, profile.Name)
		fmt.Println("Inserted", inserted.InsertedID)
		return 1
	} else {
		log.Println(existingCustomer.Err())
		return 0
	}

}

// Email Verification for Customer
func EmailVerification(data models.VerifyEmail) (string, error) {
	filter := bson.M{"email": data.Email}
	var customer models.Customer
	err := config.Customer_Collection.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		return "", err
	}
	if customer.VerificationString == data.VerificationString {
		update := bson.M{"$set": bson.M{"isemailverified": true}}
		filter := bson.M{"email": data.Email}

		_, err := config.Customer_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "", err
		}
		go SendThankYouEmail(customer.Email, customer.Name)
		return "Signup Successful", nil
	} else {
		return "Wrong OTP", nil
	}

}

// Login Customer
func Login(details models.Login) (string, int, error) {
	var customer models.Customer

	filter := bson.M{"email": details.Email}
	err := config.Customer_Collection.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		return "User not found", 0, err
	}
	if customer.WrongInput == 10 {
		return "Too many no of try", 0, nil
	}
	if !customer.IsEmailVerified {
		return "Please verify your email", 0, nil
	}
	if customer.BlockedUser {
		return "Your ID has been Blocked", 0, nil
	}
	if customer.Password != details.Password {
		customer.WrongInput++
		update := bson.M{"$set": bson.M{"wronginput": customer.WrongInput}}
		config.Customer_Collection.UpdateOne(context.Background(), filter, update)
		return "Wrong Password", 0, nil
	}

	token, err := CreateToken(customer.Email, customer.CustomerId)
	if err != nil {
		return "Internal server error", 0, nil

	}

	return token, 1, nil
}

// Add To Cart
func Addtocart(addtocart models.Addtocart) (string, error) {
	id, err := ExtractCustomerID(addtocart.Token, constants.SecretKey)
	if err != nil {
		log.Println(err)
		return "Token Expired (or) Invalid", nil
	}
	filter := bson.D{
		{Key: "$and", Value: []interface{}{
			bson.D{{Key: "customerid", Value: id}},
			bson.D{{Key: "productname", Value: addtocart.ProductName}},
		}},
	}

	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return "Internal Server Error", err
	}

	var inventoryData models.Inventory
	inventoryfilter := bson.M{"itemname": addtocart.ProductName}
	err = config.Inventory_Collection.FindOne(context.Background(), inventoryfilter).Decode(&inventoryData)
	if err != nil {
		log.Println(err)
		return "Error in Fetching", err
	}
	defer cursor.Close(context.Background())

	if cursor.RemainingBatchLength() == 0 {
		// Item not found, so insert a new item with quantity 1
		cart := models.Addcart{
			CustomerId:   id,
			ProductName:  addtocart.ProductName,
			Price:        inventoryData.Price,
			Quantity:     1,
			Image:        inventoryData.Image,
			SellerID:     inventoryData.SellerId,
			SellerName:   inventoryData.SellerName,
			ItemCategory: inventoryData.ItemCategory,
			TotalPrice:   inventoryData.Price,
		}
		_, err := config.Cart_Collection.InsertOne(context.Background(), cart)
		if err != nil {
			log.Println(err)
			return "Error in Inserting", err
		}

	} else {
		// Item already exists, update its quantity
		var cart models.Addcart
		for cursor.Next(context.Background()) {
			err = cursor.Decode(&cart)

			if err != nil {

				return "Error in Decode", err
			}
		}
		// Item already exists, update its quantity
		cart.Quantity++
		cart.TotalPrice = cart.TotalPrice + inventoryData.Price
		// Use the UpdateOne method to increment the quantity

		update := bson.M{"$set": bson.M{"quantity": cart.Quantity, "totalprice": cart.TotalPrice}}
		_, err = config.Cart_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			return "Error In Updating", err
		}

	}

	inventoryData.Stock_Available--
	update := bson.M{"$set": bson.M{"sellerquantity": inventoryData.Stock_Available}}
	_, err = config.Inventory_Collection.UpdateOne(context.Background(), inventoryfilter, update)
	if err != nil {
		return "Error In Updating", err
	}

	return "Item Added to Cart", nil
}

// Search and get Products
func Search(productName string) []models.Inventory {

	filter := bson.M{"itemcategory": productName}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(context.Background())
	var Inventory []models.Inventory
	for cursor.Next(context.Background()) {
		var inventory models.Inventory
		err := cursor.Decode(&inventory)
		if inventory.Stock_Available <= 0 {
			continue
		}
		if err != nil {
			log.Println(err)
		}

		Inventory = append(Inventory, inventory)
	}

	return Inventory
}

// Search and Find Single Data
func FetchInventoryData(search models.Search) (*models.Inventory, error) {
	var data models.Inventory
	filter := bson.M{"itemname": search.ProductName}
	err := config.Inventory_Collection.FindOne(context.Background(), filter).Decode(&data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &data, nil

}

// Get All items in Cart
func GetAllItemsinCart(token models.Token) ([]models.Addcart, string, error) {
	id, err := ExtractCustomerID(token.Token, constants.SecretKey)
	if err != nil {
		log.Println("Login Expired")
		return nil, "Login Expired", err
	}
	filter := bson.M{"customerid": id}
	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {
		return nil, "No result Found", err
	}
	defer cursor.Close(context.Background())
	var Cart []models.Addcart
	for cursor.Next(context.Background()) {
		var cart models.Addcart
		err := cursor.Decode(&cart)
		if err != nil {
			log.Println(err)
			return nil, "Error in Decode", err
		}
		Cart = append(Cart, cart)
	}
	return Cart, "Success", nil
}

// Update Cart
func UpdateCart(cart models.Addcart) (bool, string) {
	Id, err := ExtractCustomerID(cart.CustomerId, constants.SecretKey)
	if err != nil {
		log.Println(err)
		return false, "Login Expired"
	}
	cart.CustomerId = Id
	filter := bson.D{
		{Key: "$and", Value: []interface{}{
			bson.D{{Key: "customerid", Value: cart.CustomerId}},
			bson.D{{Key: "productname", Value: cart.ProductName}},
		}},
	}
	if cart.Quantity == 0 {
		_, err = config.Cart_Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Println(err)
			return false, "Can not find the Product"
		}
		return true, "Updated Successfully"
	}
	var data models.Addcart
	err = config.Cart_Collection.FindOne(context.Background(), filter).Decode(&data)
	if err != nil {
		log.Println(err)
		return false, "No items in your Cart"
	}

	if data.Quantity > cart.Quantity {
		var inventory models.Inventory
		filter := bson.M{"itemname": cart.ProductName}
		err := config.Inventory_Collection.FindOne(context.Background(), filter).Decode(&inventory)
		if err != nil {
			log.Println(err)
			return false, "Can not find the Product"
		}

		quantity := data.Quantity - cart.Quantity
		inventory.Stock_Available = inventory.Stock_Available + quantity
		update := bson.M{"$set": bson.M{"sellerquantity": inventory.Stock_Available}}
		_, err = config.Inventory_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Println(err)
			return false, "Problem in Updating"
		}
	}
	if data.Quantity < cart.Quantity {
		var inventory models.Inventory
		filter := bson.M{"itemname": cart.ProductName}
		err := config.Inventory_Collection.FindOne(context.Background(), filter).Decode(&inventory)
		if err != nil {
			log.Println(err)
			return false, "Can not find the Product"
		}
		if inventory.Stock_Available == 0 {
			return false, "No more Stock Available"
		}
		quantity := cart.Quantity - data.Quantity
		inventory.Stock_Available = inventory.Stock_Available - quantity
		update := bson.M{"$set": bson.M{"sellerquantity": inventory.Stock_Available}}
		_, err = config.Inventory_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Println(err)
			return false, "Problem in Updating"
		}
	}

	update := bson.M{"$set": bson.M{"name": cart.ProductName, "quantity": cart.Quantity, "totalprice": data.Price * float64(cart.Quantity)}}
	_, err = config.Cart_Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
		return false, "Problem in Updating"
	}
	return true, "Updated Successfully"

}

// Delete Products in Cart
func DeleteProduct(delete models.DeleteProduct) bool {
	customerid, err := ExtractCustomerID(delete.Token, constants.SecretKey)
	if err != nil {
		log.Println(err)
		return false
	}
	filter1 := bson.M{"customerid": customerid}
	filter2 := bson.M{"productname": delete.Name}
	combinedFilter := bson.M{
		"$and": []bson.M{filter1, filter2},
	}
	filter3 := bson.M{"itemname": delete.Name}
	var data models.Inventory
	err = config.Inventory_Collection.FindOne(context.Background(), filter3).Decode(&data)
	if err != nil {
		log.Println(err)
		return false
	}
	delete.Quantity = delete.Quantity + int(data.Stock_Available)
	update1 := bson.M{"$set": bson.M{"sellerquantity": delete.Quantity}}
	_, err = config.Inventory_Collection.UpdateOne(context.Background(), filter3, update1)
	if err != nil {
		log.Println(err)
		return false
	}
	log.Println("Updated In Seller")
	_, err = config.Cart_Collection.DeleteOne(context.Background(), combinedFilter)
	if err != nil {
		log.Println(err)
		return false
	}

	return true
}

// Get Customer Address
func GetUserAddress(token models.Token) (models.Address, string, error) {
	var address models.Address
	id, err := ExtractCustomerID(token.Token, constants.SecretKey)
	if err != nil {
		log.Println(err)
		return address, "Login Expired", err
	}
	filter1 := bson.M{"customerid": id}
	err = config.Customer_Collection.FindOne(context.Background(), filter1).Decode(&address)
	if err != nil {
		log.Println(err)
		return address, "Unable to find Address", err
	}
	return address, "Success", nil
}

// Add Ordered Items to Db
func CustomerOrders(Buynow models.BuyNow) (string, error) {
	for _,value := range Buynow.ItemsToBuy{
		var order models.AddOrder
		order.Address = Buynow.Address
		order.ItemsToBuy = value
		order.TotalAmount = float64(value.TotalPrice)
		order.CustomerId = Buynow.CustomerId
		order.NoofItems = value.Quantity
		order.EstimatedDeliverydate = Buynow.EstimatedDeliverydate 
		order.SellerId = value.SellerId
		order.OrderedDate = Buynow.OrderedDate
		order.OrderID = GenerateUniqueOrderID()
		order.Status.Confirmed_Order = `completed`
		order.Status.Processing_Order = "pending"
		order.Status.Product_Delivered = "pending"
		order.Status.Product_Dispatched = "pending"
		order.Status.Quality_Check = "pending"
		_, err := config.Buynow_Collection.InsertOne(context.Background(), order)
		if err != nil {
			log.Println(err)
			return "Error in inserting", err
		}
	}
	return "Success", nil
}

// Convert Id To string
func convertIDToString(insertResult *mongo.InsertOneResult) string {
	insertedID := insertResult.InsertedID.(primitive.ObjectID)
	return insertedID.Hex()
}

// Delete Orders
func DeleteOrder(delete models.DeleteOrder)(string,error) {
	filter := bson.M{"orderid": delete.OrderID}

	_, err := config.Buynow_Collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return "No result Found",err
	}

	return "Cancelled Successfully",nil
}

// Display Customer Order
func CustomerOrder(token models.Token)([]models.AddOrder,string,error) {
	id, err := ExtractCustomerID(token.Token, constants.SecretKey)
	if err != nil {
		log.Println(err)
		return nil,"Login Expired",err
	}
	filter := bson.M{"customerid": id}
	var Order []models.AddOrder
	cursor, err := config.Buynow_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return nil,"Error in Fetching Data",err
	}
	for cursor.Next(context.Background()) {
		var order models.AddOrder
		err := cursor.Decode(&order)
		if err != nil {
			log.Println(err)
			return nil,"Error in Decoding Data",err
		}
		Order = append(Order, order)
	}
	return Order,"Success",nil
}

// Delete Items In Cart
func DeleteItemsInCart(id string) (string, error) {
	filter := bson.M{"customerid": id}
	_, err := config.Cart_Collection.DeleteMany(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return "Error in Deleting", err
	}
	return "Success", nil
}

// Fetch itmes from Cart
func Itemstobuy(id string) ([]models.Item, int, int, error) {
	filter := bson.M{"customerid": id}
	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	var Item []models.Item
	var price int
	var count int
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var item models.Item
		err := cursor.Decode(&item)
		if err != nil {
			log.Println(err)
			return Item, 0, count, err
		}
		count++
		price += item.TotalPrice
		Item = append(Item, item)
	}
	log.Println(Item)
	return Item, price, count, nil
}

// Display Total Amount in Cart
func TotalAmount(id string) float64 {
	filter := bson.M{"customerid": id}
	cursor, err := config.Cart_Collection.Find(context.Background(), filter)
	if err != nil {

		log.Println(err)
	}
	var Cart float64
	defer cursor.Close(context.Background())
	for cursor.Next(context.Background()) {
		var cart models.Addcart
		err := cursor.Decode(&cart)
		if err != nil {

			log.Println(err)
		}
		// if cart.TotalPrice == 0 {
		// 	Cart = Cart + cart.Price
		// } else {
		// 	Cart = Cart + cart.TotalPrice
		// }
	}
	var currentValues models.Sales
	err = config.Sales_Collection.FindOne(context.Background(), bson.M{}).Decode(&currentValues)
	if err != nil {
		log.Println(err)
		return 0
	}
	// Update the values
	updatedTotalSalesAmount := currentValues.TotalSalesAmount + int(Cart)
	updatedTotalNoOfSales := currentValues.TotalNoOfSales + 1

	// Update the document in the collection
	_, err = config.Sales_Collection.UpdateOne(
		context.Background(),
		bson.M{},
		bson.M{
			"$set": bson.M{
				"totalsalesamount": updatedTotalSalesAmount,
				"totalnoofsales":   updatedTotalNoOfSales,
			},
		},
	)
	if err != nil {
		return 0
	}
	return Cart
}

// Add User Address
func AddUserAddress(address models.AddAddress) (string, error) {
	id, err := ExtractCustomerID(address.Token, constants.SecretKey)
	if err != nil {
		return "Login Expired", err
	}
	filter := bson.M{"customerid": id}
	update := bson.M{"$set": bson.M{"deliveryphoneno": address.DeliveryPhoneno, "deliveryemail": address.DeliveryEmail, "firstname": address.FirstName, "lastname": address.LastName, "streetname": address.Street_Name, "city": address.City, "pincode": address.Pincode}}
	_, err = config.Customer_Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Println(err)
		return "In Updating", err
	}
	return "Update SuccessFully", nil
}

// Add Customer Orders to DB
func AddCustomerOrders(Token models.Token) (string, error) {
	var BuyNow models.BuyNow
	id, err := ExtractCustomerID(Token.Token, constants.SecretKey)
	if err != nil {
		log.Println(err)
		return "Login Expired", err
	}
	BuyNow.CustomerId = id

	data, message, err := GetUserAddress(Token)
	if err != nil {
		log.Println(err)
		return message, err
	}
	BuyNow.Address = data

	ItemsToBuy, price, count, err := Itemstobuy(id)
	if err != nil {
		log.Println(err)
		return intToString(price), err
	}
	BuyNow.ItemsToBuy = ItemsToBuy
	if price <= 500 {
		BuyNow.TotalAmount = float64(price) + 50
	} else {
		BuyNow.TotalAmount = float64(price)
	}
	BuyNow.EstimatedDeliverydate = DeliveryDate()
	BuyNow.NoofItems = count
	BuyNow.OrderedDate = CurrentDate()
	orderid, err := CustomerOrders(BuyNow)
	if err != nil {
		log.Println(err)
		return message, err
	}


	message, err = DeleteItemsInCart(id)
	if err != nil {
		log.Println(err)
		return message, err
	}

	go SendOrderConformation(BuyNow.Address.DeliveryEmail, floatToString(BuyNow.TotalAmount), floatToString(float64(price)), BuyNow.EstimatedDeliverydate, orderid, intToString(BuyNow.NoofItems), BuyNow.Address)

	return "Success", nil

}

// Get DeliveryDate()
func DeliveryDate() string {
	currentTime := time.Now()
	estimatedDeliveryDate := currentTime.Add(15 * 24 * time.Hour)
	formattedDeliveryDate := estimatedDeliveryDate.Format("January 2, 2006")
	return formattedDeliveryDate
}


// To Generate Order ID
func GenerateUniqueOrderID() string {
	return fmt.Sprintf("%d%s", time.Now().UnixNano(), GetRandomString(7))
}

//Give Current Date in Format
func CurrentDate() string {
	currentTime := time.Now()
	formattedCurrentDate := currentTime.Format("January 2, 2006")
	return formattedCurrentDate
}

//Get Order Details
func GetCustromerOrder(details models.GetOrder)(models.AddOrder,string,error){
	var orderDetails models.AddOrder
	filter := bson.M{"orderid":details.OrderID}
	err:= config.Buynow_Collection.FindOne(context.Background(),filter).Decode(&orderDetails)
	if err != nil{
		return orderDetails,"No Result found",err
	}
	return orderDetails,"Success",nil
}



