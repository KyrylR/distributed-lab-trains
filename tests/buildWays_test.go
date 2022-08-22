package tests

import (
	"testing"

	"DistributedLab_Trains/algoritms/findPath"
	"DistributedLab_Trains/utils"
)

func TestBuildWays(t *testing.T) {
	allTrains, err := utils.ParseCsvToTrainStruct("data/data_for_tests.csv")
	if err != nil {
		t.Errorf("error occurred when parsing csv file")
		return
	}
	ways := findPath.BuildPathsGo(&allTrains)
	const NumberOfWays = 10
	totalNumber := 0
	for _, path := range ways {
		totalNumber += len(path.Ways)
	}
	if totalNumber != NumberOfWays {
		t.Errorf("wrong number of paths found, given: %v, should be: %v", totalNumber, NumberOfWays)
	}
}
