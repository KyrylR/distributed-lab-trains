package tests

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"DistributedLab_Trains/algoritms/executeQuery"
	"DistributedLab_Trains/algoritms/findPath"
	"DistributedLab_Trains/utils"
)

const dataPath = "data/data_for_tests.csv"

func TestQuery(t *testing.T) {
	allTrains, err := utils.ParseCsvToTrainStruct(dataPath)
	if err != nil {
		t.Errorf("error occurred when parsing csv file")
		return
	}
	ways := findPath.BuildPathsGo(&allTrains)
	for _, paths := range ways {
		for _, path := range paths.Ways {
			path.DeleteRouteDuplicates()
		}
	}
	query := executeQuery.Query{}
	err = query.Initialize(ways)
	if err != nil {
		t.Errorf("error occurred during data processing")
		return
	}
	timeQWay := [][]int{
		{1, 2, 3, 4},
		{4, 3, 2, 1},
		{2, 4, 3, 1},
	}
	timeQValue := make([]time.Duration, 3)
	timeQValue[0] += 47*time.Hour + 40*time.Minute
	timeQValue[1] += 48*time.Hour + 10*time.Minute
	timeQValue[2] += 49*time.Hour + 20*time.Minute
	result := query.QueryByTime(3)
	for idx, res := range result {
		if !reflect.DeepEqual(timeQWay[idx], res.ProcessedData.Path.Way) {
			t.Errorf("wrong path found, given: %v, should be: %v",
				fmt.Sprint(res.ProcessedData.Path.Way), fmt.Sprint(timeQWay[idx]))
		}

		if timeQValue[idx] != res.TrainPathForTime.TravelTime {
			t.Errorf("wrong path found by travel time, given: %v, should be: %v",
				res.TrainPathForTime.TravelTime, timeQValue[idx])
		}
	}
	costQWay := [][]int{
		{2, 4, 3, 1},
		{2, 3, 1, 4},
		{2, 1, 4, 3},
	}
	costs := []float64{53.00, 98.00, 100.00}
	result = query.QueryByCost(3)
	for idx, res := range result {
		currentCost, _, err := res.GetLowestCost()
		if err != nil {
			t.Errorf("error occurred when trying to get the lowest cost")
			return
		}
		if !reflect.DeepEqual(costQWay[idx], res.ProcessedData.Path.Way) {
			t.Errorf("wrong path found, given: %v, should be: %v",
				fmt.Sprint(res.ProcessedData.Path.Way), fmt.Sprint(timeQWay[idx]))
		}

		if costs[idx] != currentCost {
			t.Errorf("wrong path found by cost, given: %v, should be: %v",
				currentCost, costs[idx])
		}
	}
}
