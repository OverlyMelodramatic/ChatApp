package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var isLoggedIn = false

func getIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func postMessage(c *gin.Context) {
	var recieved string = c.PostForm("message")
	if recieved == "" {
		c.String(http.StatusBadRequest, "Cannot send an empty message")
	} else {
		c.HTML(http.StatusAccepted, "message.html", message{
			ID:     "10",
			Sender: "you",
			// TODO: export this layout in a constants file
			SentAt:  time.Now().Format(DatetimeLayout),
			Content: recieved,
		})
	}
}

func postLogout(c *gin.Context) {
	log.Println("a user just logout :(")
	isLoggedIn = false
	c.HTML(http.StatusAccepted, "login-modal.html", nil)
}

func getLogin(c *gin.Context) {
	if !isLoggedIn {
		c.HTML(http.StatusAccepted, "login-modal.html", nil)
	} else {
		c.Writer.WriteHeader(http.StatusAccepted)
	}
}

func postLogin(c *gin.Context) {
	identifiedUsername := c.PostForm("email")
	log.Printf("we have found ourselves a new user :) their name is %v", identifiedUsername)
	isLoggedIn = true
}

func getSignup(c *gin.Context) {
	if !isLoggedIn {
		c.HTML(http.StatusAccepted, "signup-modal.html", nil)
	} else {
		c.Writer.WriteHeader(http.StatusAccepted)
	}
}

func postSignup(c *gin.Context) {
	identifiedUsername := c.PostForm("email")
	log.Printf("we have found ourselves a new user :) their name is %v", identifiedUsername)
	isLoggedIn = true
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
	c.HTML(http.StatusOK, "conversation.html", gin.H{
		"correspondants": correspondants,
		"messages":       lastMessages,
	})
}
