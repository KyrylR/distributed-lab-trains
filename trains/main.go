package main

import (
	"fmt"
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
	ways := algoritms.BuildPathsGo(&allTrains)
	query := algoritms.Query{}
	err = query.Initialize(ways)
	if err != nil {
		log.Panicln(allTrains)
		return
	}
	result := query.QueryByTime(10)
	for _, res := range result {
		fmt.Println(res.String())
	}
	result = query.QueryByCost(10)
	for _, res := range result {
		fmt.Println(res.String())
	}
}
