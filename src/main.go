package main

import (
	"os"
	"text/template"

	"github.com/gin-gonic/gin"
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
}
