package models

//FeedBack Given
type Feedback struct {
	Role     string `json:"role" bson:"role"`
	Email    string `json:"email" bson:"email"`
	Feedback string `json:"feedback" bson:"feedback"`
}

//