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
	for _, paths := range ways {
		for _, path := range paths.Ways {
			path.DeleteRouteDuplicates()
		}
	}
	query := algoritms.Query{}
	err = query.Initialize(ways)
	if err != nil {
		log.Panicln(allTrains)
		return
	}
	// fmt.Println("-------------------------Time-------------------------------------")
	// result := query.QueryByTime(2)
	// for _, res := range result {
	// 	fmt.Println(res.String())
	// }
	fmt.Println("-------------------------Cost-------------------------------------")
	result := query.QueryByCost(2)
	for _, res := range result {
		fmt.Println(res.String())
	}
}
