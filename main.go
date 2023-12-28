package main

import (
	"encoding/json"
	"os"

	"github.com/gin-gonic/gin"
)

var dataMap = map[string]interface{}{}

func init() {
	data, err := os.ReadFile("./data/new_results.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &dataMap)
}

func setupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	return r
}

func main() {
	// mems := dataMap["members_wrappped"].(map[string]interface{}) //["members_wrapped"]

	server := setupRoutes()

	server.Run(":4000")
}
