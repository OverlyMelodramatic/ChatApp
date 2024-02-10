package main

import (
	"context"
	"fmt"
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const (
	defaultAddr = "127.0.0.1:80"
)

func main() {
	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"formatAsTime": fomartAsTime,
	})
	router.SetTrustedProxies(nil)

	router.GET("/", getIndex)
	router.POST("/login", postLogin)
	router.GET("/login", getLogin)
	router.POST("/signup", postSignup)
	router.GET("/signup", getSignup)
	router.POST("/logout", postLogout)

	router.POST("/message", postMessage)
	router.GET("/messages", getMessages)

	router.Static("/css", "../templates/css")
	router.Static("/js", "../js")
	router.Static("/assets", "../assets")
	router.LoadHTMLGlob("../templates/html/*")

	envAddr, addrDefined := os.LookupEnv("PORT")
	if addrDefined {
		router.Run(envAddr)
	} else {
		router.Run(defaultAddr)
	}

	// Use the SetServerAPIOptions() method to set the version of the Stable API on the client
  serverAPI := options.ServerAPI(options.ServerAPIVersion1)
  opts := options.Client().ApplyURI("mongodb+srv://mongodbcorporate845:tuE?&^2J[C^fpka$@chatapp.m6wv2cf.mongodb.net/?retryWrites=true&w=majority").SetServerAPIOptions(serverAPI)
  // Create a new client and connect to the server
  client, err := mongo.Connect(context.TODO(), opts)
  if err != nil {
    panic(err)
  }
  defer func() {
    if err = client.Disconnect(context.TODO()); err != nil {
      panic(err)
    }
  }()
  // Send a ping to confirm a successful connection
  if err := client.Database("admin").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
    panic(err)
  }
  fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")
}
