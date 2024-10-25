package main

import (
	//"fmt"
	"encoding/json"
	"fmt"

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
	start_info := handlers.StartData{}
	content, e := os.ReadFile("F:/Skillcapped/Code/LoLStatsCollector/assets/start_info.json")
	check(e)
	json.Unmarshal(content, &start_info)

	matches := handlers.GetRecentGames(start_info)
	//runForEachGame(matches, start_info)
	ProcessedGames := handlers.RunAnalysis(matches, start_info)
	MostRecentData := handlers.GetCSVData("F:/Skillcapped/Code/LoLStatsCollector/assets/processed_data.csv")

	fmt.Println(handlers.GetGameData(matches[0], start_info))
}
