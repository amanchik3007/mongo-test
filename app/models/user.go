package models

type User struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	UserID    int64  `json:"userId,omitempty" bson:"userId,omitempty"`
	FirstName string `json:"firstName" bson:"firstName,omitempty"`
	LastName  string `json:"lastName" bson:"lastName,omitempty"`
	Email     string `json:"email" bson:"email,omitempty"`
	Age       int    `json:"age" bson:"age,omitempty"`
	Weight    int64  `json:"weight" bson:"weight,omitempty"`
}
