package main

import (
	//"fmt"
	"encoding/json"
	//"net/http"
	"os"

	"github.com/WhaleMilk/LoLStatCollector/internal/handlers"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var start_info handlers.StartData
	content, e := os.ReadFile("/assets/start_info.json")
	check(e)
	err := json.Unmarshal(content, &start_info)
	check(err)

}
