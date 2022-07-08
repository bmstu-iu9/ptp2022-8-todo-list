package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

type entry struct {
	ID    string `json:"id"`
	Value string `json:"value"`
}

var entries = []entry{
	{ID: "1", Value: "12345"},
	{ID: "2", Value: "Vyacheslav, где фронт?"},
	{ID: "351", Value: "pukpuk"},
}

func main() {
	router := gin.Default()
	router.GET("/entries", getEntries)
	router.GET("/ping", ping)

	err := router.SetTrustedProxies(nil)
	if err != nil {
		os.Exit(1)
	}

	port := os.Getenv("HTTP_PORT")
	if port == "" {
		port = "8080"
	}

	err = router.Run("0.0.0.0:" + port)
	if err != nil {
		os.Exit(1)
	}
}

func getEntries(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, entries)
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, `{"Status" : "OK"}`)
}
