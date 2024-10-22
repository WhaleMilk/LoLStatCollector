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
	GD15     float32
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

	return
}

func findGD15(data Data) (gd15 int32) {
	index1 := -1
	index2 := -1
	for i, player := range data.MatchTimeline.Participants {
		if player.Puuid == data.Me.Puuid {
			index1 = i
			index2 = (i + 5) % 9
			break
		}
		if player.Puuid == data.Opponent.Puuid {
			index2 = i
			index1 = (i + 5) % 9
			break
		}
	}

	//min15frame := data.MatchTimeline.Info.Frames[15]
	//possibly make a dictionary of sorts between each num struct from the Frame struct and the indexes?
}
