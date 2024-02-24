package services

import (
	"chat_app/db"
	"chat_app/models"
	"context"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
)

type MessageService struct {
	databaseClient *db.DbClient
	validator      *validator.Validate
}

func NewMessageService(DatabaseClient *db.DbClient) *MessageService {
	return &MessageService{
		databaseClient: DatabaseClient,
		validator:      validator.New(),
	}
}

func (mc *MessageService) CreateMessage(c *gin.Context, newMessage *models.Message) (*models.Message, error) {
	userCollection := mc.databaseClient.GetCollection("messges")
	//use the validator library to validate required fields
	if validationErr := mc.validator.Struct(&newMessage); validationErr != nil {
		return nil, validationErr
	}

	ctx, cancel := context.WithCancel(c)
	defer cancel()
	_, err := userCollection.InsertOne(ctx, newMessage)
	if err != nil {
		return nil, err
	}

	return newMessage, nil
}

func (mc *MessageService) GetAllMessages(c *gin.Context, correspondantsId *[]string) (*[]models.Message, error) {
	var foundMessages []models.Message
	var filter bson.D
	// filter := bson.D{
	// 	{Key: "user1", Value: user1.ID},
	// 	{Key: "user2", Value: user2.ID},
	// }
	ctx, cancel := context.WithCancel(c)
	defer cancel()
	curr, err := mc.databaseClient.GetCollection("message").Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err := curr.Decode(&foundMessages); err != nil {
		return nil, err
	}

	return &foundMessages, nil
}
