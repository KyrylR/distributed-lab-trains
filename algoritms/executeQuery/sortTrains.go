package executeQuery

import (
	"errors"
	"sort"
	"time"

	"DistributedLab_Trains/algoritms"
	"DistributedLab_Trains/algoritms/findPath"
)

// waitToTrain is used to combine the train and waitingTime fields.
type waitToTrain struct {
	// The next train
	train algoritms.Train
	// the waitingTime field answers the question of how long to wait for the next train
	waitingTime time.Duration
}

// SortTrains stores the data used for queries
type SortTrains struct {
	Path *findPath.PossibleWay
	// key - departure station ID, value array of trains that departure from this station
	TravelTimeMap  map[int][]algoritms.Train
	WaitingTimeMap map[int][]waitToTrain // key - TrainId, value - waitToTrain struct
}

// SortTimeAndCost calls sortByCost and sortByTime functions.
// The goroutines were used to speed up the function.
func (d *SortTrains) SortTimeAndCost() error {
	if len(d.Path.Way) == 0 {
		return errors.New("no data provided")
	}
	err := d.sortByTime()
	if err != nil {
		return err
	}
	err = d.sortByCost()
	if err != nil {
		return err
	}
	return nil
}

// sortByCost - sorts Path.TrainMap by difference in cost for different trains.
func (d *SortTrains) sortByCost() error {
	for _, trains := range d.Path.TrainMap {
		less := func(i, j int) bool {
			return trains[i].Cost < trains[j].Cost
		}
		sort.Slice(trains, less)
		if !sort.SliceIsSorted(trains, less) {
			return errors.New("error occurred, while sorting by cost")
		}
	}

	return nil
}

// sortByTime - sorts TravelTimeMap by difference in arrival and departure times and fills WaitingTimeMap
// field with fillWaitingTimeMap function.
func (d *SortTrains) sortByTime() error {
	d.copyTrainMap()
	d.WaitingTimeMap = make(map[int][]waitToTrain)

	for station, trains := range d.TravelTimeMap {
		less := func(i, j int) bool {
			first := SmoothOutTime(trains[i].ArrivalTime.Sub(trains[i].DepartureTime))
			second := SmoothOutTime(trains[j].ArrivalTime.Sub(trains[j].DepartureTime))
			return first < second
		}
		sort.Slice(trains, less)
		if !sort.SliceIsSorted(trains, less) {
			return errors.New("error occurred, while sorting by time")
		}
		nextStation := findPath.GetNextStation(station, d.Path)
		if nextStation != -1 {
			trainsFromNextStation := d.TravelTimeMap[nextStation]
			for _, train := range trains {
				err := d.fillWaitingTimeMap(train, trainsFromNextStation)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}

// fillWaitingTimeMap takes Train instance and Trains slice, where train.ArrivalId is equal to trains[...].DepartureId;
// creates slice of waitToTrain structures and sorts it by waitingTime field.
// Finally, it initializes WaitingTimeMap field of SortTrains structure.
func (d *SortTrains) fillWaitingTimeMap(train algoritms.Train, trains []algoritms.Train) error {
	result := make([]waitToTrain, len(trains))
	for i := 0; i < len(result); i++ {
		result[i] = newWaitToTrain(train, trains[i])
	}
	less := func(i, j int) bool {
		return result[i].waitingTime < result[j].waitingTime
	}
	sort.Slice(result, less)
	if !sort.SliceIsSorted(result, less) {
		return errors.New("error occurred, while sorting by waiting time")
	}
	d.WaitingTimeMap[train.TrainId] = result
	return nil
}

// copyTrainMap - creates a copy of TrainMap, and initializes the TravelTimeMap field.
func (d *SortTrains) copyTrainMap() {
	d.TravelTimeMap = make(map[int][]algoritms.Train)
	for key, value := range d.Path.TrainMap {
		d.TravelTimeMap[key] = algoritms.DeepTrainSliceCopy(value)
	}
}

// newWaitToTrain returns a new instance of the waitToTrain structure that has been populated with the given data.
func newWaitToTrain(arrived algoritms.Train, waitFor algoritms.Train) waitToTrain {
	newWaitToTrain := waitToTrain{train: waitFor}
	newWaitToTrain.waitingTime = SmoothOutTime(waitFor.DepartureTime.Sub(arrived.ArrivalTime))
	return newWaitToTrain
}

// SmoothOutTime takes time.Duration and, if the value is less than 0, adds 24 hours,
// turning the value positive, and then returns it.
func SmoothOutTime(t time.Duration) time.Duration {
	if t < 0 {
		return t + time.Hour*24
	}
	return t
}
