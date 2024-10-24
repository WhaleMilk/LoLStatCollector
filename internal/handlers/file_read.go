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
	//find last row of data with total games played in each position
	//populate with that info into CSVAverages
	file_data, err := reader.ReadAll()
	check(err)

	last_line := file_data[len(file_data)-1]

}

func populateAverages(data []string) CSVAverages {
	gamespair := [2]int64{getPairVal(strconv.ParseInt(data[0], 10, 64)), getPairVal(strconv.ParseInt(data[1], 10, 64))}
	wonpair := [2]int64{getPairVal(strconv.ParseInt(data[3], 10, 64)), getPairVal(strconv.ParseInt(data[4], 10, 64))}
	gdpair := [2]int64{getPairVal(strconv.ParseInt(data[5], 10, 64)), getPairVal(strconv.ParseInt(data[6], 10, 64))}
	cspair := [2]float32{getPairValFloat(strconv.ParseFloat(data[0], 10, 64)), getPairVal(strconv.Float(data[1], 10, 64))}
	var averages CSVAverages = CSVAverages{}

	return averages
}

func getPairVal(i int64, err error) (out int64) {
	check(err)
	return i
}

func getPairValFloat(f float32, err error) (out float32) {
	check(err)
	return f
}
