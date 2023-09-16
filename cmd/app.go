package cmd

import (
	"fmt"
	"net/http"

	"github.com/golearning/api"
	"github.com/golearning/cmd/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/go-cmp/cmp"
)

func Run() {
	fmt.Println("Hello, world1.")
	fmt.Println(cmp.Diff("Hello World", "Hello Go"))
	utils.SayHello()
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		//var arr [5]string = [5]string{"hello", "world"}
		c.Data(http.StatusOK, "text/plain", []byte("Gin API v1"))
	})

	router.GET("/albums", api.GetAlbums)

	router.Run("localhost:8080")
}
