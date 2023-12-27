package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	dataMap := map[string]interface{}{}
	data, err := os.ReadFile("./data/new_results.json")
	if err != nil {
		panic(err)
	}

	json.Unmarshal(data, &dataMap)

	mems := dataMap["members_wrappped"].(map[string]interface{}) //["members_wrapped"]

	for k, v := range mems {
		fmt.Println(k, ":", v)
	}
}
