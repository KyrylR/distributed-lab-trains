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
	if len(uniqueStations) != 4 {
		t.Errorf("incorect number of unique station found, given: %v, should be: %v",
			len(uniqueStations), 4)
	}
	NumberOfTrainsThatDepFromStation := []int{2, 3, 3, 1}
	mappedData := findPath.BuildMappedData(&allTrains)
	i := 0
	for _, value := range mappedData {
		if len(value) != NumberOfTrainsThatDepFromStation[i] {
			t.Errorf("incorect number of trains fro station calculated, given: %v, should be: %v",
				len(value), NumberOfTrainsThatDepFromStation[i])
		}
		i += 1
	}
	pathTree := findPath.PathTree{DepartureId: 1}
	treeMap := pathTree.BuildPathTree(uniqueStations, mappedData)
	pathTrees := findPath.GetAllPathTrees(&pathTree, &treeMap)
	if len(pathTrees) != 4 {
		t.Errorf("incorectly determinid home many unique stations in the data, given: %v, should be: %v",
			len(pathTrees), 4)
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
