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
	handlers.RunAnalysis(matches, start_info)

	fmt.Println(handlers.GetGameData(matches[0], start_info))
}

// func runForEachGame(games []string, start_info handlers.StartData) (out FinalData) {
// 	data := handlers.Data{}
// 	for _, match := range games {
// 		data.Match = handlers.GetGameData(match, start_info)
// 		data.MatchTimeline = handlers.GetGameTimeline(match, start_info)
// 		data.Me, data.Opponent = handlers.GetPlayers(start_info.PUUID, data.Match)
// 		//handlers.runAnalysis(data)
// 	}

// 	return
// }
