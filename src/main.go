package main

import (
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	gin.SetMode(gin.DebugMode)
	router.SetFuncMap(template.FuncMap{
		"formatAsTime": fomartAsTime,
	})
	router.SetTrustedProxies(nil)

	router.GET("/", getIndex)
	router.GET("/messages", getMessages)
	router.POST("/message", postMessage)
	router.POST("/login", postLogin)
	router.GET("/login", getLogin)
	router.POST("/logout", postLogout)

	router.Static("/css", "../templates/css")
	router.Static("/assets", "../assets")
	router.LoadHTMLGlob("../templates/html/*")
	router.Run(":" + os.Getenv("PORT"))
}
