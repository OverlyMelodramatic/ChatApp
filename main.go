package main

import (
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

func getMessages(c *gin.Context) {
	var alex = user{
		ID:   "123456",
		Name: "Alexandre",
	}
	var lastMessages = []message{
		{
			ID:      "1",
			Sender:  "123456",
			SentAt:  "2024-02-02T10:21:46Z",
			Content: "Salut ! Comment ça va ?",
		},
		{
			ID:      "2",
			Sender:  "654321",
			SentAt:  "2024-02-02T10:24:57Z",
			Content: "Bah ecoute tranquille et toi ?",
		},
		{
			ID:      "3",
			Sender:  "123456",
			SentAt:  "2024-02-02T10:25:21Z",
			Content: "Je passe ma vie sur Rocket League, donc mon cycle de sommeil est horrible mais je suis heureux haha",
		},
		{
			ID:      "4",
			Sender:  "654321",
			SentAt:  "2024-02-02T10:25:58Z",
			Content: "La belle vie alors !",
		},
		{
			ID:      "5",
			Sender:  "654321",
			SentAt:  "2024-02-02T10:34:57Z",
			Content: "Je sens que je vais te suivre, on se fait une partie aujourd'hui ?",
		},
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"correspondants": alex,
		"messages":       lastMessages,
	})
}

func main() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.GET("/messages", getMessages)
	router.Static("/css", "./templates/css")
	router.LoadHTMLGlob("templates/*")
	router.Run("localhost:8080")
}
