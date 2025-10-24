package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/philip-edekobi/wrapped/types"
	"github.com/philip-edekobi/wrapped/util"
)

var (
	dataMap        = map[string]interface{}{}
	generalWrapped *types.ReadableGeneralInfo
	member         map[string]interface{}
)

func init() {
	data, err := os.ReadFile("./data/new_results.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &dataMap)

	general := dataMap["general_wrapped"].(map[string]interface{})
	member = dataMap["members_wrapped"].(map[string]interface{})

	jsonValue, err := json.Marshal(general)
	if err != nil {
		log.Fatalln(err)
	}
	wrapped := jsonValue
	generalWrapped = GetGeneralWrappedInfo(wrapped)
}

func GetMemberWrappedInfo(num string) (*types.WrappedInfo, bool) {
	memberData := member[num]

	if memberData == nil {
		return nil, false
	}

	var memberInfo types.WrappedInfo

	value, err := json.Marshal(memberData)
	if err != nil {
		log.Fatalln(err)
	}

	err = json.Unmarshal(value, &memberInfo)
	if err != nil {
		log.Fatalln(err)
	}

	return &memberInfo, true
}

func GetGeneralWrappedInfo(generalWrapped []byte) *types.ReadableGeneralInfo {
	var generalInfo types.GeneralInfo

	err := json.Unmarshal(generalWrapped, &generalInfo)
	if err != nil {
		log.Fatalln(err)
	}

	correctGeneralInfo := util.TransFormGeneralInfo(&generalInfo)

	return correctGeneralInfo
}

func setupRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/2023/general", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    generalWrapped,
			"error":   nil,
		})
	})

	r.GET("/2023/member/:num", func(c *gin.Context) {
		num := c.Param("num")

		member, exists := GetMemberWrappedInfo(num)
		if !exists {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"error":   "number does not exist",
				"data":    nil,
			})

			return
		}
		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    member,
			"error":   nil,
		})
	})

	return r
}

func main() {
	server := gin.Default()

	server.Use(cors.Default())

	server = setupRoutes(server)

	server.Run()
}
