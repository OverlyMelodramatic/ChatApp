package main

import (
	"chat_app/controllers"
	"chat_app/db"
	envVars "chat_app/utils"
	"context"
	"log"
	"text/template"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	router := gin.Default()
	router.SetFuncMap(template.FuncMap{
		"formatAsTime": fomartAsTime,
	})
	router.SetTrustedProxies(nil)
	router.GET("/", controllers.GetIndex)
	mainContext := context.TODO()
	DatabaseClient, err := db.CreateClient(&mainContext)
	defer func() {
		err := DatabaseClient.DisconnectClient(&mainContext)
		if err != nil {
			log.Fatalf("Error while disconnecting the Db Client: %v", err.Error())
		}
	}()
	if err != nil {
		log.Fatalf("Error while creating the Db Client: %v", err.Error())
		return
	}

	uc := controllers.NewUserController(DatabaseClient)
	router.GET("/login", uc.GetLogin)
	router.GET("/signup", uc.GetSignup)
	router.POST("/logout", uc.PostLogout)
	router.POST("/login", uc.GetLogin)
	router.GET("/user", uc.GetUser)
	router.POST("/signup", uc.PostSignup)

	mc := controllers.NewMessageController(DatabaseClient)
	router.POST("/message", mc.PostMessage)
	router.GET("/messages", mc.GetMessages)

	router.Static("/css", "../templates/css")
	router.Static("/js", "../js")
	router.Static("/assets", "../assets")
	router.LoadHTMLGlob("../templates/html/*")

	envAddr, err := envVars.GetEnvValue("SERVER_ADDRESS")
	if err != nil {
		log.Fatalf("Error While tring to retrieve the server address: %v", err.Error())
	}
	router.Run(envAddr)
}
