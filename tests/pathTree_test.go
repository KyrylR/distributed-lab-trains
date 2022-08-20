package tests

import (
	"testing"

	"DistributedLab_Trains/algoritms"
	"DistributedLab_Trains/utils"
)

func TestPathTree(t *testing.T) {
	allTrains, err := utils.ParseCsvToTrainStruct("data/data_for_tests.csv")
	if err != nil {
		t.Errorf("error ocured when parsing csv file")
		return
	}
	uniqueStations := algoritms.GetUniqueStations(&allTrains, false)
	mappedData := algoritms.BuildMappedData(&allTrains)
	pathTree := algoritms.PathTree{DepartureId: 1}
	treeMap := pathTree.BuildPathTree(uniqueStations, mappedData)
	pathTrees := algoritms.GetAllPathTrees(&pathTree, &treeMap)
	if len(pathTrees) != 4 {
		t.Errorf("incorectly determinid home many unique sities in the data, given: %v, should be: %v", len(pathTrees), 4)
	}
}
