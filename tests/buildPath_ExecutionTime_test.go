package tests

import (
	"fmt"
	"testing"
	"time"

	"DistributedLab_Trains/algoritms"
	"DistributedLab_Trains/utils"
)

const dataPath = "test_task_data.csv"

func TestBuildPathExecutionTime(t *testing.T) {
	allTrains, err := utils.ParseCsvToTrainStruct(dataPath)
	if err != nil {
		t.Errorf("error ocured when parsing csv file")
		return
	}
	var (
		withGoTime    time.Duration
		withoutGoTime time.Duration
	)
	start := time.Now()
	algoritms.BuildPaths(&allTrains)
	withoutGoTime = utils.TimeTrack(start, "BuildPaths")

	start = time.Now()
	algoritms.BuildPathsGo(&allTrains)
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
