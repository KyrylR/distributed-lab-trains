package apptypes

import "time"

// Train structure represent CSV from data file
type Train struct {
	TrainId            int
	DepartureStationId int
	ArrivalStationId   int
	Cost               float64
	DepartureTime      time.Time
	ArrivalTime        time.Time
}

type PossibleWay struct {
	Way      []int
	TrainMap map[int][]Train
}
