package handlers

import (
	"github.com/WhaleMilk/LoLStatCollector/internal/dataobjects"
)

func findPlayerIndex(data Data) (index int) {
	index1 := -1
	for i, player := range data.Match.Info.Participants {
		if player.Puuid == data.Me.Puuid {
			index1 = i
			break
		}
		if player.Puuid == data.Opponent.Puuid {
			index1 = (i + 5) % 9
			break
		}
	}

	return index1
}

func findGD15(data Data) (gd15 int) {
	index1 := findPlayerIndex(data)

	player1_frame, player2_frame := findPlayersGoldInFrame(data.MatchTimeline.Info.Frames[15], index1)

	return player1_frame - player2_frame
}

func findPlayersKP(data Data) (player1gold float32) {
	p_index1 := findPlayerIndex(data)
	participants := data.Match.Info.Participants
	teamKills := 0
	if p_index1 < 5 {
		for i := 0; i < 5; i++ {
			teamKills = teamKills + participants[i].Kills
		}
	} else {
		for i := 5; i < 10; i++ {
			teamKills = teamKills + participants[i].Kills
		}
	}

	return (float32(participants[p_index1].Kills+participants[p_index1].Assists) / float32(teamKills)) * 100.0
}

func findPlayerCSM(data Data) (CSM float32) {
	index := findPlayerIndex(data)
	frame_count := len(data.MatchTimeline.Info.Frames)

	cs := findPlayersCSInFrame(data.MatchTimeline.Info.Frames[frame_count-1], index)

	return cs / float32(frame_count)
}

func findPlayerDPM(data Data) (dpm float32) {
	index := findPlayerIndex(data)

	damage := findPlayersDPMInFrame(data.MatchTimeline.Info.Frames[len(data.MatchTimeline.Info.Frames)-1], index)

	return damage / float32(len(data.MatchTimeline.Info.Frames))
}

func findPlayerPos(player dataobjects.Participant) string {
	return player.IndividualPosition
}

func findPlayerWinLoss(data Data) bool {
	team := data.Me.TeamID
	if data.Match.Info.Teams[0].TeamID == team {
		return data.Match.Info.Teams[0].Win
	}
	return data.Match.Info.Teams[1].Win
}

func findPlayersGoldInFrame(frame dataobjects.Frame, p_index1 int) (player1gold int, player2gold int) {

	var frame_map = map[int]int{
		0: frame.ParticipantFrames.Num1.TotalGold,
		1: frame.ParticipantFrames.Num2.TotalGold,
		2: frame.ParticipantFrames.Num3.TotalGold,
		3: frame.ParticipantFrames.Num4.TotalGold,
		4: frame.ParticipantFrames.Num5.TotalGold,
		5: frame.ParticipantFrames.Num6.TotalGold,
		6: frame.ParticipantFrames.Num7.TotalGold,
		7: frame.ParticipantFrames.Num8.TotalGold,
		8: frame.ParticipantFrames.Num9.TotalGold,
		9: frame.ParticipantFrames.Num10.TotalGold}

	p_index2 := (p_index1 + 5) % 9
	return frame_map[p_index1], frame_map[p_index2]
}

func findPlayersCSInFrame(frame dataobjects.Frame, p_index1 int) (player1gold float32) {

	var frame_map = map[int]int{
		0: frame.ParticipantFrames.Num1.MinionsKilled + frame.ParticipantFrames.Num1.JungleMinionsKilled,
		1: frame.ParticipantFrames.Num2.MinionsKilled + frame.ParticipantFrames.Num2.JungleMinionsKilled,
		2: frame.ParticipantFrames.Num3.MinionsKilled + frame.ParticipantFrames.Num3.JungleMinionsKilled,
		3: frame.ParticipantFrames.Num4.MinionsKilled + frame.ParticipantFrames.Num4.JungleMinionsKilled,
		4: frame.ParticipantFrames.Num5.MinionsKilled + frame.ParticipantFrames.Num5.JungleMinionsKilled,
		5: frame.ParticipantFrames.Num6.MinionsKilled + frame.ParticipantFrames.Num6.JungleMinionsKilled,
		6: frame.ParticipantFrames.Num7.MinionsKilled + frame.ParticipantFrames.Num7.JungleMinionsKilled,
		7: frame.ParticipantFrames.Num8.MinionsKilled + frame.ParticipantFrames.Num8.JungleMinionsKilled,
		8: frame.ParticipantFrames.Num9.MinionsKilled + frame.ParticipantFrames.Num9.JungleMinionsKilled,
		9: frame.ParticipantFrames.Num10.MinionsKilled + frame.ParticipantFrames.Num10.JungleMinionsKilled}

	return float32(frame_map[p_index1])
}

func findPlayersDPMInFrame(frame dataobjects.Frame, p_index1 int) (player1gold float32) {

	var frame_map = map[int]int{
		0: frame.ParticipantFrames.Num1.DamageStats.TotalDamageDoneToChampions,
		1: frame.ParticipantFrames.Num2.DamageStats.TotalDamageDoneToChampions,
		2: frame.ParticipantFrames.Num3.DamageStats.TotalDamageDoneToChampions,
		3: frame.ParticipantFrames.Num4.DamageStats.TotalDamageDoneToChampions,
		4: frame.ParticipantFrames.Num5.DamageStats.TotalDamageDoneToChampions,
		5: frame.ParticipantFrames.Num6.DamageStats.TotalDamageDoneToChampions,
		6: frame.ParticipantFrames.Num7.DamageStats.TotalDamageDoneToChampions,
		7: frame.ParticipantFrames.Num8.DamageStats.TotalDamageDoneToChampions,
		8: frame.ParticipantFrames.Num9.DamageStats.TotalDamageDoneToChampions,
		9: frame.ParticipantFrames.Num10.DamageStats.TotalDamageDoneToChampions}

	return float32(frame_map[p_index1])
}
