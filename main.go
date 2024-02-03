package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// *************************** Structs ***************************

type message struct {
	ID      string `json:"id"`
	Content string `json:"content"`
	SentAt  string `json:"sent_at"`
	Sender  string `json:"sender"`
}

type user struct {
	ID   string `json:"id"`
	Name string `json:"name"`
}

// *************************** Controller ***************************

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func postMessage(c *gin.Context) {
	log.Println(c.Request.Body)
}

func getMessages(c *gin.Context) {
	var correspondants = []user{
		{
			ID:   "123456",
			Name: "Alexandre",
		},
	}
	var lastMessages = []message{
		{
			ID:      "1",
			Sender:  "them",
			SentAt:  "2024-02-02T10:21:46Z",
			Content: "Salut ! Comment ça va ?",
		},
		{
			ID:      "2",
			Sender:  "you",
			SentAt:  "2024-02-02T10:24:57Z",
			Content: "Bah ecoute tranquille et toi ?",
		},
		{
			ID:      "3",
			Sender:  "them",
			SentAt:  "2024-02-02T10:25:21Z",
			Content: "Je passe ma vie sur Rocket League, donc mon cycle de sommeil est horrible mais je suis heureux haha",
		},
		{
			ID:      "4",
			Sender:  "you",
			SentAt:  "2024-02-02T10:25:58Z",
			Content: "La belle vie alors !",
		},
		{
			ID:      "5",
			Sender:  "you",
			SentAt:  "2024-02-02T10:34:57Z",
			Content: "Je sens que je vais te suivre, on se fait une partie aujourd'hui ?",
		},
	}
	c.HTML(http.StatusOK, "messages.html", gin.H{
		"correspondants": correspondants,
		"messages":       lastMessages,
	})
}

func main() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)

	router.GET("/", getIndex)
	router.GET("/messages", getMessages)
	router.POST("/messages", postMessage)

	router.Static("/css", "./templates/css")
	router.Static("/assets", "./assets")
	router.StaticFile("favicon.ico", "./favicon.ico")
	router.LoadHTMLGlob("templates/html/*")

	router.Run("localhost:8080")
}
