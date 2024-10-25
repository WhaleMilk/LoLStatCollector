package handlers

import (
	"encoding/csv"
	"os"
)

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
