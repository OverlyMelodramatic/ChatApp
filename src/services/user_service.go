package services

import (
	"chat_app/db"
	"chat_app/models"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserService struct {
	databaseClient *db.DbClient
	validator      *validator.Validate
}

func NewUserService(databaseClient *db.DbClient) *UserService {
	return &UserService{
		databaseClient: databaseClient,
		validator:      validator.New(),
	}
}

func (userService *UserService) CreateUser(c *gin.Context, newUser *models.User) (*mongo.InsertOneResult, error) {
	userCollection := userService.databaseClient.GetCollection("users")
	if validationErr := userService.validator.Struct(newUser); validationErr != nil {
		return nil, validationErr
	}

	newUser.ID = primitive.NewObjectID()
	ctx, cancel := context.WithCancel(c)
	defer cancel()
	value, err := userCollection.InsertOne(ctx, newUser)
	return value, err
}

func (userService *UserService) FindUser(c *gin.Context, email string) (*models.User, error) {
	var foundUser models.User
	filter := bson.D{
		{Key: "credentials.email", Value: email},
	}
	ctx, cancel := context.WithCancel(c)
	defer cancel()

	err := userService.databaseClient.GetCollection("users").FindOne(ctx, filter).Decode(&foundUser)
	if err != nil {
		return nil, err
	}

	return &foundUser, nil
}
