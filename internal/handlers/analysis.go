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

type FinalData struct {
	GD15     int
	CSM      float32
	DPM      int
	KP       float32
	WinRate  float32
	Total_LP int
	LP_Delta int
}

func runAnalysis(data Data) (out FinalData) {
	if data.MatchTimeline.Info.EndOfGameResult != "GameComplete" {
		// do something, leave method, return nil?
	}

	var final_game_data FinalData = FinalData{
		findGD15(data),
		0.0,
		1,
		0.0,
		0.0,
		1,
		1}

	return final_game_data
}
