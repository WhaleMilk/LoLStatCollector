package handlers

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
	"strconv"
)

type CSVAverages struct {
	GamesPlayed [2]int
	GamesWon    [2]int
	GD_15       [2]int
	CSM         [2]float32
	DPM         [2]float32
	KP          [2]float32
	WinRate     [2]float32
	Total_LP    int
	LP_Delta    int
}

func GetCSVData(filename string) CSVAverages {
	datastream := readCSVData(filename)
	reader, err := parseCSV(datastream)
	check(err)

	MostRecentData := processCSVData(reader)
	return MostRecentData
}

func readCSVData(filename string) []byte {
	f, err := os.Open(filename)
	check(err)
	defer f.Close()

	data, err := io.ReadAll(f)
	check(err)

	return data
}

func parseCSV(data []byte) (*csv.Reader, error) {
	reader := csv.NewReader(bytes.NewReader(data))
	return reader, nil
}

func processCSVData(reader *csv.Reader) CSVAverages {
	file_data, err := reader.ReadAll()
	check(err)

	last_line := file_data[len(file_data)-1]
	return populateAverages(last_line)
}

func populateAverages(data []string) CSVAverages {

	var averages CSVAverages = CSVAverages{
		[2]int{int(getPairVal(strconv.ParseInt(data[0], 10, 64))), int(getPairVal(strconv.ParseInt(data[1], 10, 64)))},
		[2]int{int(getPairVal(strconv.ParseInt(data[3], 10, 64))), int(getPairVal(strconv.ParseInt(data[4], 10, 64)))},
		[2]int{int(getPairVal(strconv.ParseInt(data[5], 10, 64))), int(getPairVal(strconv.ParseInt(data[6], 10, 64)))},
		[2]float32{float32(getPairValFloat(strconv.ParseFloat(data[7], 32))), float32(getPairValFloat(strconv.ParseFloat(data[8], 32)))},
		[2]float32{float32(getPairValFloat(strconv.ParseFloat(data[9], 32))), float32(getPairValFloat(strconv.ParseFloat(data[10], 32)))},
		[2]float32{float32(getPairValFloat(strconv.ParseFloat(data[11], 32))), float32(getPairValFloat(strconv.ParseFloat(data[12], 32)))},
		[2]float32{float32(getPairValFloat(strconv.ParseFloat(data[13], 32))), float32(getPairValFloat(strconv.ParseFloat(data[14], 32)))},
		int(getPairVal(strconv.ParseInt(data[15], 10, 64))),
		int(getPairVal(strconv.ParseInt(data[16], 10, 64)))}

	return averages
}

func getPairVal(i int64, err error) (out int64) {
	check(err)
	return i
}

func getPairValFloat(f float64, err error) (out float64) {
	check(err)
	return f
}
