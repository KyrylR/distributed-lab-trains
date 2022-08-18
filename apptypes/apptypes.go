package apptypes

import "time"

type Train struct {
	TrainId            int
	DepartureStationId int
	ArrivalStationId   int
	Cost               float64
	DepartureTime      time.Time
	ArrivalTime        time.Time
}

type PathTree struct {
	Id              int
	DepartureId     int
	ArrivalStations map[int]bool
	Routes          []Train
	Next            *[]*PathTree
}

type PossibleWay struct {
	LastArrival int
	Way         []int
	TrainMap    map[int][]Train
}

type Path struct {
	Trains        []Train
	TotalCost     int
	TotalDuration time.Time
}
