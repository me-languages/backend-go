package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	app *gin.Engine
)

func HomePage(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"items": []map[string]string{
			{
				"version_number": "v1.0.1",
			},
		},
	})
}

func registerRouter(r *gin.RouterGroup) {
	r.GET("/home", HomePage)
}

// init gin app
func init() {
	app = gin.New()
	r := app.Group("/api")
	registerRouter(r)
}

// entrypoint
func Handler(w http.ResponseWriter, r *http.Request) {
	app.ServeHTTP(w, r)
}
