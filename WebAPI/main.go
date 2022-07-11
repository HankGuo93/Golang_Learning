package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type IndexData struct {
	Title   string
	Content string
}

func test(c *gin.Context) {
	data := new(IndexData)
	data.Title = "首頁"
	data.Content = "我的第一支 gin 專案"
	c.HTML(http.StatusOK, "index.html", data)
}

func LoginPage(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func LoginAuth(c *gin.Context) {

}

func main() {
	server := gin.Default()
	server.LoadHTMLGlob("template/html/*")
	server.Static("/assets", "./template/assets")
	server.GET("/", test)
	server.GET("/login", LoginPage)
	server.Run("0.0.0.0:8080")
}
