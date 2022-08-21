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
