package handlers

import (
	"encoding/csv"
	"os"
)

func WriteNewLine(filename string, records [][]string) {
	writer, file, err := createCSVWriter(filename)
	check(err)
	defer file.Close()

	for _, record := range records {
		writeCSVRecord(writer, record)
	}

	writer.Flush()
}

func createCSVWriter(filename string) (*csv.Writer, *os.File, error) {
	f, err := os.Create(filename)
	check(err)

	writer := csv.NewWriter(f)
	return writer, f, nil
}

func writeCSVRecord(writer *csv.Writer, record []string) {
	err := writer.Write(record)
	check(err)
}
