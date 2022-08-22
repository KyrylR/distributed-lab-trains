package algoritms

import (
	"sort"
	"time"
)

// TrainPath represents the route of the trains and the timing of the route for all trains.
type TrainPath struct {
	trains     []Train
	travelTime time.Duration
}

// TrainToTrain combines trains and stores their travel time.
type TrainToTrain struct {
	start Train
	next  Train
	// travel time of the first one + wait time of the second one + travel time of the second one
	travelTime time.Duration
}

// NewTrainPath returns a new instance of the TrainPath structure with populated fields from received data.
// The first records in the instance are the two fastest trains.
func NewTrainPath(ttt TrainToTrain, way []int) TrainPath {
	newTrainPath := TrainPath{}
	arr := make([]Train, len(way)-1)
	arr[0] = ttt.start
	arr[1] = ttt.next
	newTrainPath.trains = arr
	return newTrainPath
}

// NewTrainToTrain returns a new instance of TrainToTrain structure with populated fields from received data.
func NewTrainToTrain(start, next Train, waitingTime time.Duration) TrainToTrain {
	startTravelTime := SmoothOutTime(start.ArrivalTime.Sub(start.DepartureTime))
	nextTravelTime := SmoothOutTime(next.ArrivalTime.Sub(next.DepartureTime))
	res := TrainToTrain{
		start:      start,
		next:       next,
		travelTime: startTravelTime + nextTravelTime + waitingTime,
	}
	return res
}

// initTravelTimeField goes through train by train and calculates travel time.
// At the end it initializes travelTime field in TrainPath structure.
func (p *TrainPath) initTravelTimeField() {
	var (
		travelTime      time.Duration
		lastArrivalTime time.Time
	)
	for _, train := range p.trains {
		if lastArrivalTime.Sub(time.Time{}).Seconds() > 0 {
			travelTime += SmoothOutTime(train.DepartureTime.Sub(lastArrivalTime))
		} else {
			lastArrivalTime = train.ArrivalTime
		}
		currentTrainTime := SmoothOutTime(train.ArrivalTime.Sub(train.DepartureTime))
		travelTime += currentTrainTime
	}
	p.travelTime = travelTime
}

// GetLowestTravelTimeForPath takes a slice of the TrainPath structure, sorts it and
// returns the first TrainPath instance where travelTime is the lowest.
func GetLowestTravelTimeForPath(trainPaths []TrainPath) TrainPath {
	if len(trainPaths) < 1 {
		return TrainPath{}
	}
	less := func(i, j int) bool {
		return trainPaths[i].travelTime < trainPaths[j].travelTime
	}
	sort.Slice(trainPaths, less)
	return trainPaths[0]
}
