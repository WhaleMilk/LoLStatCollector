package handlers

import (
	"github.com/WhaleMilk/LoLStatCollector/internal/dataobjects"
)

type Data struct {
	Match         dataobjects.Match
	MatchTimeline dataobjects.Timeline
	Me            dataobjects.Participant
	Opponent      dataobjects.Participant
}

type GameSetData struct {
	GD15     int
	CSM      float32
	DPM      float32
	KP       float32
	Position string
	WinLoss  bool
}

func RunAnalysis(games []string, start_info StartData) (out []GameSetData) {
	var game_sets []GameSetData
	var data Data

	for _, match := range games {
		data.Match = GetGameData(match, start_info)
		data.MatchTimeline = GetGameTimeline(match, start_info)
		if len(data.MatchTimeline.Info.Frames) < 15 {
			continue
		}
		data.Me, data.Opponent = GetPlayers(start_info.PUUID, data.Match)
		game_sets = append(game_sets, runForEachGame(data))
	}

	return game_sets
}

func runForEachGame(data Data) GameSetData {

	if data.MatchTimeline.Info.EndOfGameResult != "GameComplete" {
		return GameSetData{}
	}

	var final_game_data GameSetData = GameSetData{
		findGD15(data),
		findPlayerCSM(data),
		findPlayerDPM(data),
		findPlayersKP(data),
		findPlayerPos(data.Me),
		findPlayerWinLoss(data)}

	return final_game_data
}
