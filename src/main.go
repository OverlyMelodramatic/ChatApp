package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
	"time"

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

// *************************** Formatter ***************************

func fomartAsTime(t string) string {
	layout := "2006-01-02T15:04:05.000Z"
	var ti, err = time.Parse(layout, t)
	if err != nil {
		log.Panicln("found uncompliant date" + err.Error())
	}
	return fmt.Sprintf("%02d:%02d", ti.Local().Hour(), ti.Minute())
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
			SentAt:  "2024-02-02T10:21:46.453Z",
			Content: "Salut ! Comment ça va ?",
		},
		{
			ID:      "2",
			Sender:  "you",
			SentAt:  "2024-02-02T10:24:57.045Z",
			Content: "Bah ecoute tranquille et toi ?",
		},
		{
			ID:      "3",
			Sender:  "them",
			SentAt:  "2024-02-02T10:25:21.765Z",
			Content: "Je passe ma vie sur Rocket League, donc mon cycle de sommeil est horrible mais je suis heureux haha",
		},
		{
			ID:      "4",
			Sender:  "you",
			SentAt:  "2024-02-02T10:25:58.385Z",
			Content: "La belle vie alors !",
		},
		{
			ID:      "5",
			Sender:  "you",
			SentAt:  "2024-02-02T10:34:57.205Z",
			Content: "Je sens que je vais te suivre, on se fait une partie aujourd'hui ?",
		},
		{
			ID:      "6",
			Sender:  "them",
			SentAt:  "2024-02-02T10:36:23.320Z",
			Content: "Ok, j'suis libre pendant une heure, allons-y :)",
		},
		{
			ID:      "7",
			Sender:  "you",
			SentAt:  "2024-02-02T10:47:57.587Z",
			Content: "Je me connecte, je te dis quand j'arrive",
		},
		{
			ID:      "8",
			Sender:  "you",
			SentAt:  "2024-02-02T10:48:57.238Z",
			Content: "C'est bon je suis là, je t'attends",
		},
		{
			ID:      "9",
			Sender:  "you",
			SentAt:  "2024-02-02T14:32:34.320Z",
			Content: "C'était cool, faut qu'on se refasse ça!",
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
	router.SetFuncMap(template.FuncMap{
		"formatAsTime": fomartAsTime,
	})

	router.GET("/", getIndex)
	router.GET("/messages", getMessages)
	router.POST("/messages", postMessage)

	router.Static("/css", "../templates/css")
	router.Static("/assets", "../assets")
	router.StaticFile("favicon.ico", "../favicon.ico")
	router.LoadHTMLGlob("../templates/html/*")

	router.Run("localhost:80")
}
