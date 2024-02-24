package controllers

import (
	"chat_app/db"
	"chat_app/models"
	"chat_app/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MessageController struct {
	messageService *services.MessageService
}

func NewMessageController(DatabaseClient *db.DbClient) *MessageController {
	return &MessageController{
		messageService: services.NewMessageService(DatabaseClient),
	}
}

func (mc *MessageController) PostMessage(c *gin.Context) {
	var newMessage models.Message
	c.Request.ParseForm()
	if err := c.BindJSON(newMessage); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	newM, err := mc.messageService.CreateMessage(c, &newMessage)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.HTML(http.StatusAccepted, "message.html", newM)
}

func (mc *MessageController) GetMessages(c *gin.Context) {
	corrs := c.QueryArray("correspondants-id")
	lastMessages, err := mc.messageService.GetAllMessages(c, &corrs)
	if err != nil {
		log.Fatalln(err)
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	c.HTML(http.StatusOK, "conversation.html", gin.H{
		"correspondants": corrs,
		"messages":       lastMessages,
	})
}
