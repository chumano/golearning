package cmd

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/golearning/api"
	"github.com/golearning/utils"

	"github.com/gin-gonic/gin"

	"github.com/gin-contrib/static"
)

type App struct {
	router *gin.Engine
}

func (app *App) Setup() {
	router := gin.Default()

	app.router = router

	app.serveWebStatic()

	router.GET("/api", func(c *gin.Context) {
		//var arr [5]string = [5]string{"hello", "world"}
		c.Data(http.StatusOK, "text/plain", []byte("Gin API v1"))
	})

	router.GET("/api/albums", api.GetAlbums)
}

func (app *App) Run(addr string) {
	fmt.Println("=======================")
	utils.SayHello()
	fmt.Println("=======================")

	err := app.router.Run(addr)
	if err != nil {
		log.Fatal(fmt.Errorf("CHUNO: %v", err))
		panic(1)
	}
}

func (app *App) serveWebStatic() {
	router := app.router

	webDir := utils.GetWebDirPath()
	router.Use(static.Serve("/", static.LocalFile(webDir, true)))
	router.NoRoute(func(c *gin.Context) {
		if !strings.HasPrefix(c.Request.RequestURI, "/api") {
			c.File(webDir + "/index.html")
		}
		//default 404 page not found
	})
}
