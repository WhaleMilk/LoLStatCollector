package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	//"io"
	"time"
	//"text/template"
)

type StartData struct {
	api_key string
	PUUID   string
	date    string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getRecentGames(date string, start_data StartData) {
	start_end, err := GetEpochTimes(date)
	check(err)
	var recent_games []string

	resp, err := http.Get(fmt.Sprintf("https://americas.api.riotgames.com/lol/match/v5/matches/by-puuid/%p/ids%s&api_key=%q", &start_data.PUUID, start_end, start_data.api_key))
	check(err)

	body, err := io.ReadAll(resp.Body)
	check(err)

	json.Unmarshal(body, &recent_games)
}

func GetEpochTimes(date string) (string, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return "", err
	}

	startOfDay := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, parsedDate.Location())
	endOfDay := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 23, 59, 59, 0, parsedDate.Location())

	return fmt.Sprintf("?startTime=%s&endTime=%q", startOfDay, endOfDay), nil
	// return startOfDay.Unix(), endOfDay.Unix(), nil
}
