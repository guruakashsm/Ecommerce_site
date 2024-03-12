package service

import (
	"context"
	"ecommerce/config"
	"ecommerce/constants"
	"ecommerce/models"
	"log"
	"regexp"

	"go.mongodb.org/mongo-driver/bson"
)

func isValidNumber(s string) bool {
	numericRegex := regexp.MustCompile("^[0-9]+$")
	return numericRegex.MatchString(s)
}

func countdigits(number int) int {
	count := 0
	for number > 0 {
		count++
		number = number / 10
	}
	return int(count)
}

func Validatetoken(token string) bool {
	_, err := ExtractCustomerID(token, constants.SecretKey)
	return err == nil
}

func AddEvent(upload models.UploadCalender) error {
	_, err := config.Calender_Collection.InsertOne(context.Background(), upload)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

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
