package service

import (
	"context"
	"ecommerce/config"
	"ecommerce/constants"
	"ecommerce/models"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Admin Login
func AdminLoginCheck(login *models.AdminData) (string, int) {

	var correctdata models.AdminData
	filter := bson.M{"email": login.Email}
	err := config.Admin_Collection.FindOne(context.Background(), filter).Decode(&correctdata)
	if err != nil {
		return "", 0
	}
	if correctdata.WrongInput == 4 {
		return "", 1
	}
	if correctdata.IP_Address == login.IP_Address {
		return "", 3
	}
	if correctdata.Password != login.Password {
		correctdata.WrongInput++
		update := bson.M{"$set": bson.M{"wronginput": correctdata.WrongInput}}
		config.Admin_Collection.UpdateOne(context.Background(), filter, update)
		return "", 2
	}

	if !ValidateOTP(login.TOTP, correctdata.SecretKey) {
		correctdata.WrongInput++
		update := bson.M{"$set": bson.M{"wronginput": correctdata.WrongInput}}
		config.Admin_Collection.UpdateOne(context.Background(), filter, update)
		return "", 4
	}

	token, err := CreateToken(correctdata.Email, correctdata.AdminID)
	if err != nil {
		return "", 5
	}

	log.Println(token)
	update := bson.M{"$set": bson.M{"wronginput": 0}}
	config.Admin_Collection.UpdateOne(context.Background(), filter, update)
	return token, 5

}

// TO get all Customer
func GetallCustomerdata() ([]models.Customer, string, error) {
	filter := bson.D{}
	cursor, err := config.Customer_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(context.Background())
	var Profiles []models.Customer
	for cursor.Next(context.Background()) {
		var profile models.Customer
		err := cursor.Decode(&profile)
		if err != nil {
			return nil, "Error in Decode", err
		}
		Profiles = append(Profiles, profile)
	}
	return Profiles, "Success", nil
}

// Get All Inventory
func Getinventorydata() []models.Inventory {
	filter := bson.D{}
	cursor, err := config.Inventory_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(context.Background())
	var Inventorydata []models.Inventory
	for cursor.Next(context.Background()) {
		var inventory models.Inventory
		err := cursor.Decode(&inventory)
		if err != nil {
			log.Println(err)
		}
		Inventorydata = append(Inventorydata, inventory)
	}
	return Inventorydata
}

// Get All Seller
func Getallsellerdata() []models.Seller {
	filter := bson.D{}
	cursor, err := config.Seller_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	defer cursor.Close(context.Background())
	var Seller []models.Seller
	for cursor.Next(context.Background()) {
		var seller models.Seller
		err := cursor.Decode(&seller)
		if err != nil {
			log.Println(err)
		}
		Seller = append(Seller, seller)
	}
	return Seller
}

// Create Seller
func CreateSeller(seller models.Seller) bool {
	if seller.Password != seller.ConfirmPassword {
		return false
	}
	filter := bson.M{"selleremail": seller.Seller_Email}
	cursor, err := config.Seller_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
		return false
	}
	defer cursor.Close(context.Background())

	if cursor.RemainingBatchLength() == 0 {
		seller.SellerId = GenerateUniqueCustomerID()
		seller.BlockedUser = false
		seller.WrongInput = 0
		seller.IsApproved = true
		seller.IsEmailVerified = true
		_, err := config.Seller_Collection.InsertOne(context.Background(), seller)
		if err != nil {
			log.Println(err)
			return false

		}
		go SendSellerInvitation(seller.Seller_Email, seller.Seller_Name, seller.Password, "https://anon.up.railway.app/seller/")
		return true
	}
	return false
}

// Update Any Data
func Update(update models.Update) bool {
	if update.Collection == "seller" {
		filter := bson.M{"selleremail": update.IdName}
		update1 := bson.M{"$set": bson.M{update.Field: update.New_Value}}
		options := options.Update()
		_, err := config.Seller_Collection.UpdateOne(context.TODO(), filter, update1, options)
		if err != nil {
			return false
		}
		go SendEditDataNotification(update.IdName, update.Field, update.New_Value)
		return true
	} else if update.Collection == "customer" {
		if update.Field == "phonenumber" || update.Field == "age" || update.Field == "pincode" {

			intValue, err := strconv.Atoi(update.New_Value)
			if err != nil {
				return false
			} else {
				update.New_Value = strconv.Itoa(intValue)
			}
			if !isValidNumber(update.New_Value) {
				return false
			}
			filter := bson.M{"email": update.IdName}
			update1 := bson.M{"$set": bson.M{update.Field: intValue}}
			options := options.Update()
			_, err1 := config.Customer_Collection.UpdateOne(context.TODO(), filter, update1, options)

			if err1 != nil {

				return false
			}

		}

		filter := bson.M{"email": update.IdName}
		update1 := bson.M{"$set": bson.M{update.Field: update.New_Value}}
		options := options.Update()
		_, err := config.Customer_Collection.UpdateOne(context.TODO(), filter, update1, options)
		fmt.Println("updated")
		if err != nil {
			return false
		}
		go SendEditDataNotification(update.IdName, update.Field, update.New_Value)
		return true

	} else if update.Collection == "inventory" {
		if update.Field == "price" {
			// Check if New_Value is a valid integer
			intValue, err := strconv.Atoi(update.New_Value)
			if err != nil {
				// Handle the error, e.g., return an error response or log it
				return false
			}

			// Check if the input value is a valid number (numeric characters only)
			if !isValidNumber(update.New_Value) {
				return false
			}

			filter := bson.M{"itemname": update.IdName}
			update1 := bson.M{"$set": bson.M{update.Field: intValue}}
			options := options.Update()
			_, err1 := config.Inventory_Collection.UpdateOne(context.TODO(), filter, update1, options)
			return err1 == nil

		}

		filter := bson.M{"itemname": update.IdName}
		update1 := bson.M{"$set": bson.M{update.Field: update.New_Value}}
		options := options.Update()
		_, err := config.Inventory_Collection.UpdateOne(context.TODO(), filter, update1, options)
		return err != nil
	}

	return false
}

// Delete Any data
func Delete(delete models.Delete) bool {
	if delete.Collection == "customer" {
		filter := bson.M{"email": delete.IdValue}
		_, err := config.Customer_Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	}
	if delete.Collection == "seller" {
		filter := bson.M{"selleremail": delete.IdValue}
		_, err := config.Seller_Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	}
	if delete.Collection == "inventory" {
		filter := bson.M{"itemname": delete.IdValue}
		_, err := config.Inventory_Collection.DeleteOne(context.Background(), filter)
		if err != nil {
			log.Println(err)
			return false
		}
		return true
	}
	return true
}

// Get Dataneed for Admin
func AdminNeededData() models.AdminPageData {
	var adminpagedata models.AdminPageData
	var sales models.Sales
	adminpagedata.ProductCount, _ = config.Inventory_Collection.CountDocuments(context.Background(), bson.D{})

	adminpagedata.UserCount, _ = config.Customer_Collection.CountDocuments(context.Background(), bson.D{})

	adminpagedata.SellerCount, _ = config.Seller_Collection.CountDocuments(context.Background(), bson.D{})

	config.Sales_Collection.FindOne(context.Background(), bson.M{}).Decode(&sales)

	adminpagedata.SalesCount = int64(sales.TotalNoOfSales)

	adminpagedata.TotalSalesAmount = int32(sales.TotalSalesAmount)

	return adminpagedata
}

func GetWorkerdata() []models.Workers {
	var workers []models.Workers

	filter := bson.M{}
	cursor, err := config.Worker_Collection.Find(context.Background(), filter)
	if err != nil {
		log.Println(err)
	}
	for cursor.Next(context.Background()) {
		var worker models.Workers
		err := cursor.Decode(&worker)
		if err != nil {
			log.Println(err)
		}
		workers = append(workers, worker)
	}
	return workers
}

// Create Worker
func CreateWorker(worker models.Workers) string {
	filter := bson.M{"email": worker.Email}
	result := config.Worker_Collection.FindOne(context.Background(), filter)
	if result.Err() == nil {
		return "User Already Exists"
	}
	if result.Err() != nil && result.Err() != mongo.ErrNoDocuments {
		return "Error in Query: " + result.Err().Error()
	}
	_, err := config.Worker_Collection.InsertOne(context.Background(), worker)
	if err != nil {
		return "Error in Creating: " + err.Error()
	}
	return "Created Successfully"
}

// Create Admin
func CreateAdmin(admin models.AdminSignup) (string, string) {
	filter := bson.M{"email": admin.Email}

	result := config.Admin_Collection.FindOne(context.Background(), filter)
	if result.Err() == nil {
		return "User Already Exists", ""
	}
	if result.Err() != nil && result.Err() != mongo.ErrNoDocuments {
		return "Error in Query: " + result.Err().Error(), ""
	}
	key, err := GenerateSecret()
	if err != nil {
		return "Error In Generating TOTP", ""
	}
	var AdminData models.AdminData
	AdminData.Email = admin.Email
	AdminData.Password = admin.Password
	AdminData.IP_Address = admin.IP
	AdminData.SecretKey = key
	AdminData.Token = ""
	AdminData.WrongInput = 0
	AdminData.AdminID = GenerateUniqueAdminID()
	_, err = config.Admin_Collection.InsertOne(context.Background(), AdminData)
	if err != nil {
		return "Error in Creating: " + err.Error(), ""
	}
	go SendAdminInvitation(admin.Email, admin.AdminName, admin.Password, "https://anon.up.railway.app/admin/", admin.IP, key)
	return "Created Successfully", key
}

// Get Single Data
func GetData(data models.Getdata) (*models.ReturnData, error) {
	var returndata models.ReturnData

	if data.Collection == "customer" {
		log.Println("In customer")
		var profile models.Customer
		filter := bson.M{"email": data.Id}
		err := config.Customer_Collection.FindOne(context.Background(), filter).Decode(&profile)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		returndata.Name = profile.Name
		returndata.CustomerId = profile.CustomerId
		returndata.Address = profile.Address
		returndata.Email = profile.Email
		returndata.Phone_No = profile.Phone_No
		returndata.Password = profile.Password
		returndata.IsEmailVerified = profile.IsEmailVerified
		returndata.BlockedUser = profile.BlockedUser
		returndata.WrongInput = profile.WrongInput
		return &returndata, nil

	} else if data.Collection == "seller" {
		log.Println("In seller")
		var profile models.Seller
		filter := bson.M{"selleremail": data.Id}
		log.Println()
		err := config.Seller_Collection.FindOne(context.Background(), filter).Decode(&profile)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		returndata.Seller_Name = profile.Seller_Name
		returndata.Phone_No = profile.Phone_No
		returndata.Address = profile.Address
		returndata.Password = profile.Password
		returndata.SellerId = profile.SellerId
		returndata.Seller_Email = profile.Seller_Email
		returndata.Seller_Name = profile.Seller_Name
		returndata.Image = profile.Image
		returndata.BlockedUser = profile.BlockedUser
		returndata.WrongInput = profile.WrongInput
		return &returndata, nil
	} else if data.Collection == "inventory" {
		log.Println("In inventory")
		var profile models.Inventory
		filter := bson.M{"itemname": data.Id}
		err := config.Inventory_Collection.FindOne(context.Background(), filter).Decode(&profile)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		returndata.ItemCategory = profile.ItemCategory
		returndata.ItemName = profile.ItemName
		returndata.Quantity = profile.Quantity
		returndata.Seller_Name = profile.SellerName
		returndata.Price = profile.Price
		returndata.Stock_Available = profile.Stock_Available
		returndata.Image = profile.Image
		return &returndata, nil
	} else if data.Collection == "worker" {
		log.Println("In worker")
		var profile models.Workers
		filter := bson.M{"email": data.Id}
		err := config.Worker_Collection.FindOne(context.Background(), filter).Decode(&profile)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		returndata.Email = profile.Email
		returndata.No = profile.No
		returndata.Role = profile.Role
		returndata.Status = profile.Status
		returndata.UserName = profile.UserName
		returndata.Salary = profile.Salary
		returndata.Image = profile.Image
		return &returndata, nil
	}
	return nil, nil

}

// Block User & Admin
func Block(data models.Block) (string, error) {
	if data.Collection == "customer" {
		var customer models.Customer
		filter := bson.M{"email": data.Email}
		err := config.Customer_Collection.FindOne(context.Background(), filter).Decode(&customer)
		if err != nil {
			log.Println(err)
			return "No result Found", err
		}
		message := ""
		if customer.BlockedUser {
			customer.BlockedUser = false
			message = "Customer has been Unblocked"
		} else {
			customer.BlockedUser = true
			message = "Customer has been Blocked"
		}
		update := bson.M{"$set": bson.M{"blockeduser": customer.BlockedUser}}
		_, err = config.Customer_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Println(err)
			return "Can't Update Data", err
		}
		go SendBlockingNotification(customer.Email, customer.Name, "Due to improper behaviour")
		return message, nil
	} else if data.Collection == "seller" {
		var Seller models.Seller
		filter := bson.M{"selleremail": data.Email}
		err := config.Seller_Collection.FindOne(context.Background(), filter).Decode(&Seller)
		if err != nil {
			log.Println(err)
			return "No result Found", err
		}
		message := ""
		if Seller.BlockedUser {
			Seller.BlockedUser = false
			message = "Seller has been Unblocked"
		} else {
			Seller.BlockedUser = true
			message = "Seller has been Blocked"
		}
		update := bson.M{"$set": bson.M{"blockeduser": Seller.BlockedUser}}
		_, err = config.Seller_Collection.UpdateOne(context.Background(), filter, update)
		if err != nil {
			log.Println(err)
			return "Can't Update Data", err
		}
		go SendBlockingNotification(Seller.Seller_Email, Seller.Seller_Name, "Due to improper behaviour")
		return message, nil

	}
	return "Invalid Collection", nil
}

// Add Event To Calender
func AddEvent(upload models.UploadCalender) error {
	_, err := config.Calender_Collection.InsertOne(context.Background(), upload)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// Get Event from Calender
func GetEvent(GetData models.GetCalender) ([]models.UploadCalender, error) {
	filter := bson.M{"email": GetData.AdminEmail}
	cursor, err := config.Calender_Collection.Find(context.Background(), filter)
	var Data []models.UploadCalender
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for cursor.Next(context.Background()) {
		var data models.UploadCalender
		err := cursor.Decode(&data)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		Data = append(Data, data)
	}
	return Data, nil
}

// ShutDown
func ShutDown(token models.ShutDown) (string, error) {
	if token.Password != constants.ShutDownKey {
		return "Key Mismatch", nil
	}
	id, err := ExtractCustomerID(token.Token, constants.SecretKey)
	if err != nil {
		log.Println("Login Exp")
		return "Login Expired", err
	}
	var admin models.AdminData
	filter := bson.M{"adminid": id}
	err = config.Admin_Collection.FindOne(context.Background(), filter).Decode(&admin)
	if err != nil {
		return "Login as Admin", err
	}

	if id != admin.AdminID {
		return "Login as Admin", err
	}

	shutdownComplete := make(chan bool)

	go func() {
		ShutDownExe()
		shutdownComplete <- true
	}()

	return "Shutdown initiated", nil
}

func ShutDownExe() {
	time.Sleep(3 * time.Second)
	os.Exit(0)
}

// Clear DataBase
func ClearDB(details models.Getdata) (string, error) {
	id, err := ExtractCustomerID(details.Id, constants.SecretKey)
	if err != nil {
		return "Login Expired", err
	}
	var admin models.AdminData
	filter := bson.M{"adminid": id}
	err = config.Admin_Collection.FindOne(context.Background(), filter).Decode(&admin)
	if err != nil {
		return "Data not Found", err
	}
	if admin.Email == "" {
		return "Data not Found", nil
	}
	result, err := DeleteDBCollection(details.Collection)
	if err != nil {
		return result, err
	}
	return result, nil
}

// Delete colletion
func DeleteDBCollection(collection string) (string, error) {
	if collection == "all" {
		err := config.Admin_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Admin Collection", err
		}
		err = config.Buynow_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Orders Collection", err
		}
		err = config.Calender_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Calender Collection", err
		}
		err = config.Cart_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Cart Collection", err
		}
		err = config.Customer_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Customers Collection", err
		}
		err = config.Feedback_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting FeedBack Collection", err
		}
		err = config.Inventory_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Inventory Collection", err
		}
		err = config.Seller_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Sellers Collection", err
		}
		err = config.Worker_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Workers Collection", err
		}
		return "All Database Deleted Successfully", nil
	} else if collection == "sellerall" {
		err := config.Buynow_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Orders Collection", err
		}
		err = config.Feedback_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting FeedBack Collection", err
		}
		err = config.Inventory_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Inventory Collection", err
		}
		err = config.Seller_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Sellers Collection", err
		}
		return "Seller Related Database Deleted Successfully", nil
	} else if collection == "customerall" {
		err := config.Buynow_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Orders Collection", err
		}
		err = config.Cart_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Cart Collection", err
		}
		err = config.Customer_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Customers Collection", err
		}
		err = config.Feedback_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting FeedBack Collection", err
		}
		return "Customer Related Database Deleted Successfully", nil
	} else if collection == "adminall" {
		err := config.Admin_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Admin Collection", err
		}
		err = config.Calender_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Calender Collection", err
		}
		return "Adim related Database Deleted Successfully", nil
	} else if collection == "seller" {
		err := config.Seller_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Sellers Collection", err
		}
		return "Seller Database Deleted Successfully", nil
	} else if collection == "inventory" {
		err := config.Inventory_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Inventory Collection", err
		}
		return "Inventory Database Deleted Successfully", nil
	} else if collection == "orders" {
		err := config.Buynow_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Orders Collection", err
		}
		return "Order Database Deleted Successfully", nil
	} else if collection == "feedback" {
		err := config.Feedback_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting FeedBack Collection", err
		}
		return "Feedback Database Deleted Successfully", nil
	} else if collection == "worker" {
		err := config.Worker_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Workers Collection", err
		}
		return "Worker Database Deleted Successfully", nil
	} else if collection == "cart" {
		err := config.Cart_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Cart Collection", err
		}
		return "Cart Database Deleted Successfully", nil
	} else if collection == "calender" {
		err := config.Calender_Collection.Drop(context.Background())
		if err != nil {
			return "Error in Delting Calender Collection", err
		}
		return "Calender Database Deleted Successfully", nil
	}
	return "Collection Not Found", nil

}

func GetAllNotApprovedSeller(token models.Token) ([]models.Seller, string, error) {
	id, err := ExtractCustomerID(token.Token, constants.SecretKey)
	if err != nil {
		return nil, "Login Expired", err
	}
	var admin models.AdminData
	filter := bson.M{"adminid": id}
	err = config.Admin_Collection.FindOne(context.Background(), filter).Decode(&admin)
	if err != nil {
		return nil,"Data not Found", err
	}
	if admin.Email == "" {
		return nil,"Data not Found", nil
	}
    var Seller []models.Seller
	filter = bson.M{"isapproved":false}
	filter2 := bson.M{"isemailverified":true}
	filter3 := bson.M{"blockeduser":false}
	combinedFilter := bson.M{
		"$and": []bson.M{filter, filter2,filter3},
	}
    cursor,err := config.Seller_Collection.Find(context.Background(),combinedFilter)
	if err != nil{
		return nil,"Error in Finding",err
	}
	for cursor.Next(context.Background()) {
		var seller models.Seller
		err := cursor.Decode(&seller)
		if err != nil {
			log.Println(err)
			return nil,"Internal Server Error",err
		}
		Seller = append(Seller, seller)
	}
   defer cursor.Close(context.Background())
   return Seller,"Success",nil
}


func ApproveSeller(details models.ApproveSeller)(string,error){
	id, err := ExtractCustomerID(details.Token, constants.SecretKey)
	if err != nil {
		return "Login Expired", err
	}
	var admin models.AdminData
	filter := bson.M{"adminid": id}
	err = config.Admin_Collection.FindOne(context.Background(), filter).Decode(&admin)
	if err != nil {
		return "Data not Found", err
	}
	if admin.Email == "" {
		return "Data not Found", nil
	}
	filter =  bson.M{"sellerid": id}
	update := bson.M{"$set": bson.M{"isapproved": true}}
	_,err = config.Seller_Collection.UpdateOne(context.Background(),filter,update)
    if err != nil{
		return "Data Not Found",err
	}
	return "Success",nil
}
