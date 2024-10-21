package main

import (
	//"fmt"
	//"encoding/json"
	//"net/http"
	//"io"
	"time"
)

func main() {
	const PUUID string = "etqNgHQY0OaE_LSnmSZZiYPBYNkOZQbO31cNpDsbmzTx36--Xjx7C99RgxIqBWggeaqq1o6FBNTz5Q"
	// var api_key string = ""

}

func GetEpochTimes(date string) (int64, int64, error) {
	parsedDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return 0, 0, err
	}

	startOfDay := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 0, 0, 0, 0, parsedDate.Location())
	endOfDay := time.Date(parsedDate.Year(), parsedDate.Month(), parsedDate.Day(), 23, 59, 59, 0, parsedDate.Location())

	return startOfDay.Unix(), endOfDay.Unix(), nil
}
