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
		data.Me, data.Opponent = GetPlayers(start_info.PUUID, data.Match)
		game_sets = append(game_sets, runForEachGame(data))
	}
	/* TODO:
	* Change Final Data around to just be a CSVAverages struct
	* Move code in main.runForEachMatch() to here so that main only has to call this method
	* That allows for the creation of struct we populate with data for each game we play
	* then we can send that to a file_write.go file to spit the averaged data out into a csv
	* Need to keep on hand a count of how many jungle and mid games are played as we iterate
	* Possibly a data struct that stores the individual scores for each data which we can then add onto corresponding fields in the CSV struct?
	 */

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
