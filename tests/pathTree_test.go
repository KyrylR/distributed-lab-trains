package tests

import (
	"testing"

	"DistributedLab_Trains/algoritms/findPath"
	"DistributedLab_Trains/utils"
)

func TestPathTree(t *testing.T) {
	allTrains, err := utils.ParseCsvToTrainStruct("data/data_for_tests.csv")
	if err != nil {
		t.Errorf("error occurred when parsing csv file")
		return
	}
	uniqueStations := findPath.GetUniqueStations(&allTrains, false)
	mappedData := findPath.BuildMappedData(&allTrains)
	pathTree := findPath.PathTree{DepartureId: 1}
	treeMap := pathTree.BuildPathTree(uniqueStations, mappedData)
	pathTrees := findPath.GetAllPathTrees(&pathTree, &treeMap)
	if len(pathTrees) != 4 {
		t.Errorf("incorectly determinid home many unique sities in the data, given: %v, should be: %v", len(pathTrees), 4)
	}

	for _, tree := range pathTrees {
		for _, route := range tree.Routes {
			if route.DepartureStationId != tree.DepartureId {
				t.Errorf("incorectly added route to path, given: %v, should be: %v",
					route.DepartureStationId, tree.DepartureId)
			}
		}
	}
}
