package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

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

	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "ok",
		})
	})

	r.GET("/readyz", func(c *gin.Context) {
		// Check that generalWrapped is loaded
		if generalWrapped == nil {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "not ready",
			})
			return
		}

		// Optionally, check if member is not empty
		if len(member) == 0 {
			c.JSON(http.StatusServiceUnavailable, gin.H{
				"status": "not ready",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status": "ready",
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
	// Load .env file if it exists
	if err := godotenv.Load(); err != nil {
		log.Printf("No .env file found or error loading .env file: %v", err)
	}

	server := gin.Default()

	// Configure CORS from ALLOWED_ORIGINS environment variable (from .env file).
	// ALLOWED_ORIGINS should be a comma-separated list of allowed origins.
	// Example in .env file: ALLOWED_ORIGINS=http://localhost:3000,https://example.com
	allowedEnv := os.Getenv("ALLOWED_ORIGINS")
	var allowOrigins []string
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
	}

	if allowedEnv != "" {
		parts := strings.Split(allowedEnv, ",")
		for _, p := range parts {
			p = strings.TrimSpace(p)
			if p == "" {
				continue
			}
			// If a wildcard is provided, enable AllowAllOrigins
			if p == "*" {
				corsConfig.AllowAllOrigins = true
				allowOrigins = nil
				break
			}
			allowOrigins = append(allowOrigins, p)
		}
	}

	if !corsConfig.AllowAllOrigins {
		corsConfig.AllowOrigins = allowOrigins
	}

	log.Printf("CORS allowlist: %v", corsConfig.AllowOrigins)
	server.Use(cors.New(corsConfig))

	server = setupRoutes(server)

	server.Run()
}
