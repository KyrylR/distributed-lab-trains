package algoritms

import "time"

// Train structure represents a single entry in the CSV data file.
type Train struct {
	TrainId            int
	DepartureStationId int
	ArrivalStationId   int
	Cost               float64
	DepartureTime      time.Time
	ArrivalTime        time.Time
}

// DeepTrainSliceCopy makes a full copy of the train slice.
func DeepTrainSliceCopy(arr []Train) []Train {
	newArr := make([]Train, len(arr))
	for i := 0; i < len(arr); i++ {
		newArr[i] = arr[i]
	}
	return newArr
}

// RemoveDuplicates takes a Train slice and produces an output containing
// all unique stations that are compared by TrainId.
func RemoveDuplicates(arr []Train) []Train {
	unique := map[int]Train{}
	for _, train := range arr {
		unique[train.TrainId] = train
	}

	result := make([]Train, len(unique))
	i := 0

	for _, v := range unique {
		result[i] = v
		i++
	}

	return result
}
