package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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

func GetMemberWrappedInfo(num string) *WrappedInfo {
	memberData := member[num]
	var memberInfo WrappedInfo

	value, err := json.Marshal(memberData)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(value, &memberInfo)
	if err != nil {
		log.Fatalln(err)
	}

	return &memberInfo
}

func GetGeneralWrappedInfo() []byte {
	return generalWrapped
}

func setupRoutes() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/2023/general", func(c *gin.Context) {
		generalInfo := GetGeneralWrappedInfo()

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    string(generalInfo),
		})
	})

	r.GET("/2023/member/:num", func(c *gin.Context) {
		num := c.Param("num")

		fmt.Println("num", num)

		member := GetMemberWrappedInfo(num)
		c.JSON(http.StatusOK, member)
	})

	return r
}

func main() {
	server := setupRoutes()

	server.Run()
}
