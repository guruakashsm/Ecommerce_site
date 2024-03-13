package service

import (
	"context"
	"ecommerce/config"
	"ecommerce/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
)


// Instert FeedBack
func Feedback(feedback models.Feedback) int {
	insertedid, err := config.Feedback_Collection.InsertOne(context.Background(), feedback)
	if err != nil {
		log.Println(err)
		return 3
	}
	fmt.Println(insertedid.InsertedID)
	return 1
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

//Delete FeedBack
func Deletefeedback(delete models.Feedback) int32 {
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

//Get all Feedbacks
func GetFeedBacks() []models.Feedback {
	filter := bson.M{}
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
