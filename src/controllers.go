package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func postMessage(c *gin.Context) {
	var message string = c.PostForm("message")
	var templateArgs gin.H = nil
	if message == "" {
		templateArgs = gin.H{
			"error": "Cannot send an empty message",
		}
	} else {
		log.Printf("New Message posted by user: %v", message)
	}
	c.HTML(http.StatusOK, "user-input-form.html", templateArgs)
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
