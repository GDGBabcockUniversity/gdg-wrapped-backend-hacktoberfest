package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

var (
	dataMap        = map[string]interface{}{}
	generalWrapped []byte
	member         map[string]interface{}
)

func init() {
	data, err := os.ReadFile("./data/new_results.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &dataMap)

	general := dataMap["general_wrapped"].(map[string]interface{})
	member = dataMap["members_wrappped"].(map[string]interface{})

	jsonValue, err := json.Marshal(general)
	if err != nil {
		log.Fatalln(err)
	}
	generalWrapped = jsonValue
}

func GetMemberWrappedInfo(num string) []byte {
	memberData := member[num]

	value, err := json.Marshal(memberData)
	if err != nil {
		log.Fatalln(err)
	}

	return value
}

func GetGeneralWrappedInfo() []byte {
	return generalWrapped
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
	server := setupRoutes()

	server.Run(":4000")
}
