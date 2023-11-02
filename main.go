package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

// init gin app
func Init() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.GET("/", HomePage)
	r.Group("/")
	return r
}

// entrypoint
func main() {
	routersInit := Init()
	endPoint := fmt.Sprintf(":%d", 4000)
	server := &http.Server{
		Addr:    endPoint,
		Handler: routersInit,
	}
	server.ListenAndServe()
}
