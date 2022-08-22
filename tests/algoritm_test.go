package tests

import (
	"testing"

	"DistributedLab_Trains/algoritms/findPath"
	"DistributedLab_Trains/utils"
)

func TestMyApproach(t *testing.T) {
	allTrains, err := utils.ParseCsvToTrainStruct("data/data_for_tests.csv")
	if err != nil {
		t.Errorf("error ocured when parsing csv file")
		return
	}
	paths := findPath.BuildPaths(&allTrains)
	sum := 0
	for _, path := range paths {
		sum += len(path.Ways)
	}
	if sum != 10 {
		t.Errorf("incorect number of founded ways, given: %v, should be: %v", sum, 10)
	}
}
