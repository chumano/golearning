package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAlbums(c *gin.Context) {
	var arr [2]string = [2]string{"hello", "albums"}
	c.IndentedJSON(http.StatusOK, arr)
}
