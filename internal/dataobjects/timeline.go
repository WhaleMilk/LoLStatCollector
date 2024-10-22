package dataobjects

type Timeline struct {
	Metadata struct {
		DataVersion  string   `json:"dataVersion"`
		MatchID      string   `json:"matchId"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		EndOfGameResult string  `json:"endOfGameResult"`
		FrameInterval   int     `json:"frameInterval"`
		Frames          []Frame `json:"frames"`
	}
	GameId       string `json:"gameId"`
	Participants []struct {
		ParticipantId int    `json:"participantId"`
		Puuid         string `json:"puuid"`
	}
}
