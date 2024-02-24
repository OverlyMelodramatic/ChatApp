package db

import (
	envVars "chat_app/utils"
	"context"
	"fmt"
	"strconv"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/mongo/mongodriver"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DbClient struct {
	mongoClient *mongo.Client
	dbName      string
	authDbName  string
	maxAge      int
}

func CreateClient(ctx *context.Context) (*DbClient, error) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	mongoURI, err := envVars.GetEnvValue("URI_MONGODB")
	if err != nil {
		return nil, err
	}

	opts := options.Client().ApplyURI(mongoURI).SetServerAPIOptions(serverAPI)
	mongoClient, err := mongo.Connect(*ctx, opts)
	if err != nil {
		return nil, err
	}

	dbName, err := envVars.GetEnvValue("DB_NAME")
	if err != nil {
		return nil, err
	}
	authDbName, err := envVars.GetEnvValue("AUTH_DB_NAME")
	if err != nil {
		return nil, err
	}
	maxAgeString, err := envVars.GetEnvValue("MAX_DB_CONNEXION_AGE")
	if err != nil {
		return nil, err
	}
	maxAge, err := strconv.Atoi(maxAgeString)
	if err != nil {
		return nil, fmt.Errorf("could not convert the value found in 'MAX_DB_CONNEXION_AGE' to a string. Value found: %v", maxAgeString)
	}

	return &DbClient{
		dbName:      dbName,
		maxAge:      maxAge,
		authDbName:  authDbName,
		mongoClient: mongoClient,
	}, nil
}

func (client *DbClient) DisconnectClient(ctx *context.Context) error {
	return client.mongoClient.Disconnect(*ctx)
}

// TODO: Finish this so we can have an auth middleware
func (client *DbClient) AuthMiddleware() gin.HandlerFunc {
	c := client.GetCollection("sessions")
	store := mongodriver.NewStore(c, client.maxAge, true, []byte("secret"))
	return sessions.Sessions("mysession", store)
}

func (client *DbClient) GetCollection(collectionName string) *mongo.Collection {
	return client.mongoClient.Database(client.dbName).Collection(collectionName)
}
