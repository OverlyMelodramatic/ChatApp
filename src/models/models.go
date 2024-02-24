package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Message struct {
	ID      primitive.ObjectID `bson:"_id" validate:"omitempty"`
	Content string             `json:"content" validate:"required"`
	SentAt  string             `json:"sent_at" validate:"required"`
	Sender  string             `json:"sender" validate:"required"`
}

type User struct {
	ID          primitive.ObjectID `bson:"_id" validate:"omitempty"`
	FirstName   string             `json:"first_name" validate:"required,alpha"`
	LastName    string             `json:"last_name" validate:"required,alpha"`
	Credentials UserCredentials    `json:"credentials,omitempty"`
}

type UserCredentials struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type Conversation struct {
	ID           primitive.ObjectID   `bson:"_id" validate:"omitempty"`
	Participants []primitive.ObjectID `json:"participants_id" validate:"required"`
	LastMessages []Message            `json:"messages" validate:"omitempty"`
}
