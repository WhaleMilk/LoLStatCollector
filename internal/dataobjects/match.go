package dataobjects

type Match struct {
	Metadata struct {
		DataVersion  string   `json:"dataVersion"`
		MatchID      string   `json:"matchId"`
		Participants []string `json:"participants"`
	} `json:"metadata"`
	Info struct {
		EndOfGameResult    string        `json:"endOfGameResult"`
		GameCreation       int64         `json:"gameCreation"`
		GameDuration       int           `json:"gameDuration"`
		GameEndTimestamp   int64         `json:"gameEndTimestamp"`
		GameID             int64         `json:"gameId"`
		GameMode           string        `json:"gameMode"`
		GameName           string        `json:"gameName"`
		GameStartTimestamp int64         `json:"gameStartTimestamp"`
		GameType           string        `json:"gameType"`
		GameVersion        string        `json:"gameVersion"`
		MapID              int           `json:"mapId"`
		Participants       []Participant `json:"participants"`
		PlatformID         string        `json:"platformId"`
		QueueID            int           `json:"queueId"`
		Teams              []struct {
			Bans []struct {
				ChampionID int `json:"championId"`
				PickTurn   int `json:"pickTurn"`
			} `json:"bans"`
			Objectives struct {
				Baron struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"baron"`
				Champion struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"champion"`
				Dragon struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"dragon"`
				Horde struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"horde"`
				Inhibitor struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"inhibitor"`
				RiftHerald struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"riftHerald"`
				Tower struct {
					First bool `json:"first"`
					Kills int  `json:"kills"`
				} `json:"tower"`
			} `json:"objectives"`
			TeamID int  `json:"teamId"`
			Win    bool `json:"win"`
		} `json:"teams"`
		TournamentCode string `json:"tournamentCode"`
	} `json:"info"`
}
