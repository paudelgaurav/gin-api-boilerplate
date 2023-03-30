package internal

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type Router struct {
	*gin.Engine
}

// NewRouter : all the routes are defined here
func NewRouter() Router {
	httpRouter := gin.Default()

	httpRouter.GET("/health-check", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "All is well",
		})
	})
	TYPE := os.Getenv("ENV")
	if TYPE == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	return Router{httpRouter}
}

func (r *Router) RunServer() {
	port := Config("SERVER_PORT")
	if port == "" {
		r.Run()
	} else {
		r.Run(port)
	}
}
