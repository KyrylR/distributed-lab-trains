package apptypes

import "time"

type Data struct {
	TrainId            int
	DepartureStationId int
	ArrivalStationId   int
	Cost               float64
	DepartureTime      time.Time
	ArrivalTime        time.Time
}
