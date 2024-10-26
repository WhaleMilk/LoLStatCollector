package handlers

import (
	"encoding/csv"
	"os"
	"strconv"
)

func CalcNewLine(games []GameSetData, lastLine CSVAverages, start StartData) []string {
	var games_count [2]uint8
	var pos uint8

	for _, match := range games {
		if match.Position == "JUNGLE" {
			pos = 0
		} else {
			pos = 1
		}

		games_count[pos]++
		lastLine.GD_15[pos] += match.GD15
		lastLine.CSM[pos] += match.CSM
		lastLine.DPM[pos] += match.DPM
		lastLine.KP[pos] += match.KP

		if match.WinLoss {
			lastLine.GamesWon[pos]++
		}
	}

	lastLine.GamesPlayed[0] += int(games_count[0])
	lastLine.GamesPlayed[1] += int(games_count[1])

	calcNewAverages(games_count, &lastLine)
	calcNewWR(&lastLine)

	currentLP := GetRankedData(start.SummonerID, start.ApiKey)
	lastLine.LP_Delta = currentLP - lastLine.Total_LP
	lastLine.Total_LP = currentLP

	return convertToSlice(lastLine)
}

func calcNewAverages(games_count [2]uint8, lastLine *CSVAverages) {
	for i := 0; i < 2; i++ {
		if games_count[i] != 0 {
			lastLine.GD_15[i] /= int(games_count[i]) + 1
			lastLine.CSM[i] /= float32(games_count[i]) + 1
			lastLine.DPM[i] /= float32(games_count[i]) + 1
			lastLine.KP[i] /= float32(games_count[i]) + 1
		}
	}
}

func calcNewWR(lastLine *CSVAverages) {
	lastLine.WinRate[0] = 100.0 * (float32(lastLine.GamesWon[0]) / float32(lastLine.GamesPlayed[0]))
	lastLine.WinRate[1] = 100.0 * (float32(lastLine.GamesWon[1]) / float32(lastLine.GamesPlayed[1]))
}

func convertToSlice(data CSVAverages) []string {
	var out []string = []string{
		strconv.Itoa(data.GamesPlayed[0]),
		strconv.Itoa(data.GamesPlayed[1]),
		strconv.Itoa(data.GamesWon[0]),
		strconv.Itoa(data.GamesWon[1]),
		strconv.Itoa(data.GD_15[0]),
		strconv.Itoa(data.GD_15[1]),
		strconv.FormatFloat(float64(data.CSM[0]), 'f', 2, 64),
		strconv.FormatFloat(float64(data.CSM[1]), 'f', 2, 64),
		strconv.FormatFloat(float64(data.DPM[0]), 'f', 2, 64),
		strconv.FormatFloat(float64(data.DPM[1]), 'f', 2, 64),
		strconv.FormatFloat(float64(data.KP[0]), 'f', 2, 64),
		strconv.FormatFloat(float64(data.KP[1]), 'f', 2, 64),
		strconv.FormatFloat(float64(data.WinRate[0]), 'f', 2, 64),
		strconv.FormatFloat(float64(data.WinRate[1]), 'f', 2, 64),
		strconv.Itoa(data.Total_LP),
		strconv.Itoa(data.LP_Delta)}

	return out
}

func WriteNewLine(filename string, record []string) {
	writer, file, err := createCSVWriter(filename)
	check(err)
	defer file.Close()

	writeCSVRecord(writer, record)

	writer.Flush()
}

func createCSVWriter(filename string) (*csv.Writer, *os.File, error) {
	f, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	check(err)

	writer := csv.NewWriter(f)
	return writer, f, nil
}

func writeCSVRecord(writer *csv.Writer, record []string) {
	err := writer.Write(record)
	check(err)
}
