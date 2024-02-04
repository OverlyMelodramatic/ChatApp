package main

import (
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.SetFuncMap(template.FuncMap{
		"formatAsTime": fomartAsTime,
	})

	router.GET("/", getIndex)
	router.GET("/messages", getMessages)
	router.POST("/message", postMessage)

	router.Static("/css", "../templates/css")
	router.Static("/assets", "../assets")
	router.StaticFile("favicon.ico", "../favicon.ico")
	router.LoadHTMLGlob("../templates/html/*")

	router.Run("localhost:80")
}
