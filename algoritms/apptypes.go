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
