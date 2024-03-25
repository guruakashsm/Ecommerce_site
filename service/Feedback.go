package service

import (
	"context"
	"ecommerce/config"
	"ecommerce/constants"
	"ecommerce/models"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)

// Instert FeedBack
func InstertSellerFeedback(feedback models.Feedback) (string, error) {
	id, err := ExtractCustomerID(feedback.Token, constants.SecretKey)
	if err != nil {
		return "Login Expired", err
	}
	var seller models.Seller
	filter := bson.M{"sellerid": id}
	err = config.Seller_Collection.FindOne(context.Background(), filter).Decode(&seller)
	if err != nil {
		return "Details not Found", nil
	}

	var FeedbackDB models.FeedbackDB
	FeedbackDB.Email = seller.Seller_Email
	FeedbackDB.Feedback = feedback.Feedback
	FeedbackDB.Role = "seller"
	_, err = config.Feedback_Collection.InsertOne(context.Background(), FeedbackDB)
	if err != nil {
		log.Println(err)
		return "Error in Inserting", err
	}

	return "FeedBack Submited Successfully", nil
}

// Instert FeedBack
func InsertCustomerFeedback(feedback models.Feedback) (string, error) {
	id, err := ExtractCustomerID(feedback.Token, constants.SecretKey)
	if err != nil {
		return "Login Expired", err
	}
	var customer models.Customer
	filter := bson.M{"customer": id}
	err = config.Customer_Collection.FindOne(context.Background(), filter).Decode(&customer)
	if err != nil {
		return "Details not Found", nil
	}

	var FeedbackDB models.FeedbackDB
	FeedbackDB.Email = customer.Email
	FeedbackDB.Feedback = feedback.Feedback
	FeedbackDB.Role = "customer"
	_, err = config.Feedback_Collection.InsertOne(context.Background(), FeedbackDB)
	if err != nil {
		log.Println(err)
		return "Error in Inserting", err
	}

	return "FeedBack Submited Successfully", nil
}

// Get Customer FeedBack
func CustomerFeedback() []models.Feedback {
	filter := bson.M{"role": "customer"}
	cursor, err := config.Feedback_Collection.Find(context.Background(), filter)
	var Feedback []models.Feedback
	if err != nil {
		log.Println(err)
	}
	for cursor.Next(context.Background()) {
		var feedback models.Feedback
		err := cursor.Decode(&feedback)
		if err != nil {
			log.Println(err)
		}
		Feedback = append(Feedback, feedback)
	}
	return Feedback
}

// Get Seller FeedBack
func SellerFeedback() []models.Feedback {
	filter := bson.M{"role": "seller"}
	cursor, err := config.Feedback_Collection.Find(context.Background(), filter)
	var Feedback []models.Feedback
	if err != nil {
		log.Println(err)
	}
	for cursor.Next(context.Background()) {
		var feedback models.Feedback
		err := cursor.Decode(&feedback)
		if err != nil {
			log.Println(err)
		}
		Feedback = append(Feedback, feedback)
	}
	return Feedback
}

// Delete FeedBack
func Deletefeedback(delete models.FeedbackDB) int32 {
	filter1 := bson.M{"email": delete.Email}
	filter2 := bson.M{"feedback": delete.Feedback}
	combinedFilter := bson.M{
		"$and": []bson.M{filter1, filter2},
	}
	_, err := config.Feedback_Collection.DeleteMany(context.Background(), combinedFilter)
	if err != nil {
		return 0
	}
	return 1

}

// Get all Feedbacks
func GetFeedBacks() []models.FeedbackDB {
	filter := bson.M{}
	cursor, err := config.Feedback_Collection.Find(context.Background(), filter)
	var Feedback []models.FeedbackDB
	if err != nil {
		log.Println(err)
	}

	for cursor.Next(context.Background()) {
		var feedback models.FeedbackDB
		err := cursor.Decode(&feedback)
		if err != nil {
			log.Println(err)
		}
		Feedback = append(Feedback, feedback)
	}
	return Feedback
}
