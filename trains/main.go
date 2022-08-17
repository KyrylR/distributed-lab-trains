package main

import (
	"DistributedLab_Trains/utils"
	"fmt"
)

const dataPath = "data/test_task_data.csv"

func main() {
	fmt.Println(utils.ParseCsvToDataStruct(dataPath))
}
