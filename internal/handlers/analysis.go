package handlers

import (
	"github.com/WhaleMilk/LoLStatCollector/internal/dataobjects"
)

type data struct {
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

func runAnalysis(match dataobjects.Match, timeline dataobjects.Timeline, me dataobjects.Participant, opp dataobjects.Participant) {

}
