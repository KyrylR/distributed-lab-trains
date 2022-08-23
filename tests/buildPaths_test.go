package tests

import (
	"fmt"
	"testing"
	"time"

	"DistributedLab_Trains/algoritms/findPath"
	"DistributedLab_Trains/utils"
)

func TestBuildPaths(t *testing.T) {
	allTrains, err := utils.ParseCsvToTrainStruct("data/data_for_tests.csv")
	if err != nil {
		t.Errorf("error occurred when parsing csv file")
		return
	}
	const NumberOfWays = 10
	paths := findPath.BuildPathsGo(&allTrains)
	sum := 0
	for _, path := range paths {
		sum += len(path.Ways)
	}
	if sum != NumberOfWays {
		t.Errorf("incorect number of founded ways, given: %v, should be: %v", sum, NumberOfWays)
	}
	for _, path := range paths {
		for _, way := range path.Ways {
			for _, station := range way.Way {
				if station != way.GetLastStation() {
					if len(way.TrainMap[station]) != 1 {
						t.Errorf("incorect number of trains in TrainMap for station: %v, given: %v, should be: %v",
							station, len(way.TrainMap[station]), 1)
					}
				}
			}
			if len(way.Way) != 4 {
				t.Errorf("incorect number of station in way, given: %v, should be: %v",
					len(way.Way), 4)
			}
		}
	}
}

func TestBuildPathExecutionTime(t *testing.T) {
	allTrains, err := utils.ParseCsvToTrainStruct("data/test_task_data.csv")
	if err != nil {
		t.Errorf("error ocured when parsing csv file")
		return
	}
	var (
		withGoTime    time.Duration
		withoutGoTime time.Duration
	)
	start := time.Now()
	findPath.BuildPaths(&allTrains)
	withoutGoTime = utils.TimeTrack(start, "BuildPaths")

	start = time.Now()
	findPath.BuildPathsGo(&allTrains)
	withGoTime = utils.TimeTrack(start, "BuildPathsGo")

	fmt.Println(getFasterResult(withoutGoTime, withGoTime))
}

func getFasterResult(withoutGoTime, withGoTime time.Duration) string {
	if withoutGoTime-withGoTime < 0 {
		res := float64(withoutGoTime.Nanoseconds()) / float64(withGoTime.Nanoseconds())
		return fmt.Sprintf("BuildPaths is faster than BuildPathsGo about %v times\n", res)
	}
	res := float64(withoutGoTime.Nanoseconds()) / float64(withGoTime.Nanoseconds())
	return fmt.Sprintf("BuildPathsGo is faster than BuildPaths about %v times\n", res)
}
