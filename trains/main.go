package main

import (
	"log"

	"DistributedLab_Trains/algoritms"
	"DistributedLab_Trains/utils"
)

const dataPath = "data/test_task_data.csv"

func main() {
	allTrains, err := utils.ParseCsvToTrainStruct(dataPath)
	if err != nil {
		log.Panicln(allTrains)
		return
	}
	algoritms.BuildPathsGo(&allTrains)
}
