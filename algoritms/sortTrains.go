package algoritms

import (
	"errors"
	"sort"
	"time"
)

type waitToTrain struct {
	train       Train
	waitingTime time.Duration
}

type SortTrains struct {
	Path *PossibleWay
	// key - departure station ID, value array of trains that departure from this station
	TravelTimeMap  map[int][]Train
	WaitingTimeMap map[int][]waitToTrain // key - TrainId, value - waitToTrain struct
}

func (d *SortTrains) SortTimeAndCost() error {
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

func (d *SortTrains) sortByCost() error {
	if len(d.Path.Way) == 0 {
		return errors.New("no data provided")
	}
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

func (d *SortTrains) sortByTime() error {
	d.copyTrainMap()
	d.WaitingTimeMap = make(map[int][]waitToTrain)

	if len(d.Path.Way) == 0 {
		return errors.New("no data provided")
	}
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
		nextStation := GetNextStation(station, d.Path)
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

func (d *SortTrains) fillWaitingTimeMap(train Train, trains []Train) error {
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

func (d *SortTrains) copyTrainMap() {
	d.TravelTimeMap = make(map[int][]Train)
	for key, value := range d.Path.TrainMap {
		d.TravelTimeMap[key] = DeepTrainSliceCopy(value)
	}
}

func newWaitToTrain(arrived Train, waitFor Train) waitToTrain {
	newWaitToTrain := waitToTrain{train: waitFor}
	newWaitToTrain.waitingTime = SmoothOutTime(waitFor.DepartureTime.Sub(arrived.ArrivalTime))
	return newWaitToTrain
}

func SmoothOutTime(t time.Duration) time.Duration {
	if t < 0 {
		return t + time.Hour*24
	}
	return t
}
