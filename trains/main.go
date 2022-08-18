package main

import (
	"DistributedLab_Trains/algoritms"
	"DistributedLab_Trains/utils"
	"log"
)

const dataPath = "data/test_task_data_test.csv"

func main() {
	allTrains, err := utils.ParseCsvToDataStruct(dataPath)
	if err != nil {
		log.Panicln(allTrains)
		return
	}
	algoritms.BuildPaths(&allTrains)
}
