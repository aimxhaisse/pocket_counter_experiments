package main

import (
	"github.com/jinzhu/now"
	"log"
	"time"
)

var LastHeight int64 = 0
var LastTime time.Time
var FirstTime time.Time

func ComputeDataForYear(year int, config Config) {
	log.Printf("computing data for year %d", year)

	for i := 1; i <= 12; i += 1 {
		first_day_of_month := time.Date(year, time.Month(i), 1, 0, 0, 0, 0, time.UTC)
		last_day_of_month := now.With(first_day_of_month).EndOfMonth()

		log.Printf("Looking for blocks at [%s] -> [%s]", first_day_of_month, last_day_of_month)

		if first_day_of_month.Before(FirstTime) {
			log.Printf("skipping as there was no pocket at the time")
			continue
		}

		if time.Now().Before(last_day_of_month) {
			log.Printf("leaving now as we can't go into the future, yet")
			break
		}

		block_start, block_end := GetClosestHeights(LastHeight, first_day_of_month, LastTime, last_day_of_month, config)
		log.Printf("For month %s: %d -> %d", first_day_of_month, block_start, block_end)
	}
}

func main() {
	config := getConfig("config/config.json")

	log.Printf("fetching last block height")

	var err error
	LastHeight, err = GetLatestHeight(config)
	if err != nil {
		log.Fatal("can't get last block height: %v", err)
	}
	log.Printf("last block height is %d", LastHeight)

	current, err := GetBlock(LastHeight, config)
	LastTime = current.Block.Time
	first, err := GetBlock(1, config)
	FirstTime = first.Block.Time

	ComputeDataForYear(2020, config)
	ComputeDataForYear(2021, config)
}	
