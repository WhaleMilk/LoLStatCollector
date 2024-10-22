package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	//"io"
	"time"

	"github.com/WhaleMilk/LoLStatCollector/internal/dataobjects"
	//"text/template"
)

type StartData struct {
	ApiKey string `json:"apiKey"`
	PUUID  string `json:"PUUID"`
	Date   string `json:"date"`
}

type GamesPlayed struct {
	Game_ids []string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func GetRecentGames(start_data StartData) (games []string) {
	start_end, err := GetEpochTimes(start_data.Date)
	check(err)
	var Recent_games GamesPlayed

	requestString := fmt.Sprint("https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/", start_data.PUUID, "/ids", start_end, "&api_key=", start_data.ApiKey)

	resp, err := http.Get(requestString)
	check(err)

	body, err := io.ReadAll(resp.Body)
	check(err)

	enc := fmt.Sprint(`{"Game_ids":`, string(body), `}`)
	json.Unmarshal([]byte(enc), &Recent_games)
	return Recent_games.Game_ids
}

func GetGameData(game_id string, start_data StartData) (match_data dataobjects.Match) {
	var match dataobjects.Match

	resp, err := http.Get(fmt.Sprint("https://americas.api.riotgames.com/lol/match/v5/matches/", game_id, "?api_key=", start_data.ApiKey))
	check(err)

	body, err := io.ReadAll(resp.Body)
	check(err)

	json.Unmarshal(body, &match)
	return match
}

func GetPlayers(PUUID string, match dataobjects.Match) (me dataobjects.Participant, opponent dataobjects.Participant) {
	index := 0
	for i, partic := range match.Info.Participants {
		if partic.Puuid == PUUID {
			index = i
			break
		}
	}

	return match.Info.Participants[index], match.Info.Participants[(index+5)%9]
}

func GetGameTimeline(game_id string, start_data StartData) (match_timeline dataobjects.Timeline) {
	var timeline dataobjects.Timeline

	resp, err := http.Get(fmt.Sprint("https://americas.api.riotgames.com/lol/match/v5/matches/", game_id, "/timeline?api_key=", start_data.ApiKey))
	check(err)

	body, err := io.ReadAll(resp.Body)
	check(err)

	json.Unmarshal(body, &timeline)
	return timeline
}

func GetEpochTimes(date string) (string, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}

	startOfDay := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, parsedDate.Location())
	endOfDay := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 23, 59, 59, 0, parsedDate.Location())

	return fmt.Sprint("?startTime=", startOfDay.Unix(), "&endTime=", endOfDay.Unix()), nil
}
