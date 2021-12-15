package oauth

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"send-email/constants"
)

type (
	// ErrorHandleFunc error handling function
	ErrorHandleFunc func(*gin.Context, error)
	// Config defines the config for Session middleware
	Config struct {
		// error handling when starting the session
		ErrorHandleFunc ErrorHandleFunc
		// keys stored in the context
		TokenKey string
		// defines a function to skip middleware.Returning true skips processing
		// the middleware.
		Skipper func(*gin.Context) bool
	}
)

var (
	// DefaultConfig is the default middleware config.
	DefaultConfig = Config{
		ErrorHandleFunc: func(ctx *gin.Context, err error) {
			ctx.AbortWithError(500, err)
		},
		TokenKey: constants.TokenKey,
		Skipper: func(_ *gin.Context) bool {
			return false
		},
	}
)

// HandleTokenVerify Verify the access token of the middleware
func HandleTokenVerify(config ...Config) gin.HandlerFunc {
	cfg := DefaultConfig
	if len(config) > 0 {
		cfg = config[0]
	}

	if cfg.ErrorHandleFunc == nil {
		cfg.ErrorHandleFunc = DefaultConfig.ErrorHandleFunc
	}

	tokenKey := cfg.TokenKey
	clientDomain := constants.ClientDomain
	if tokenKey == "" {
		tokenKey = DefaultConfig.TokenKey
	}

	return func(c *gin.Context) {
		if cfg.Skipper != nil && cfg.Skipper(c) {
			c.Next()
			return
		}
		ti, err := gServer.ValidationBearerToken(c.Request)
		if err != nil {
			cfg.ErrorHandleFunc(c, err)
			c.JSON(http.StatusUnauthorized, gin.H{"detail": err.Error()})
			return
		}
		client, err := gServer.Manager.GetClient(c, ti.GetClientID())
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"detail": "Cannot get client domain"})
			return
		}

		c.Set(tokenKey, ti)
		c.Set(clientDomain, client.GetDomain())
		c.Next()
	}
}