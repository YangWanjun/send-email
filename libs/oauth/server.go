package oauth

import (
	"github.com/gin-gonic/gin"
	"github.com/go-oauth2/oauth2/v4/server"
	"net/http"
	"sync"
)

var (
	gServer *server.Server
	once sync.Once
)

func InitServer() *server.Server {
	once.Do(func() {
		manager := InitManager()
		gServer = server.NewDefaultServer(manager)
	})
	return gServer
}

func HandleAuthorizeRequest(c *gin.Context) {
	err := gServer.HandleAuthorizeRequest(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Abort()
}

func HandleTokenRequest(c *gin.Context) {
	err := gServer.HandleTokenRequest(c.Writer, c.Request)
	if err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}
	c.Abort()
}