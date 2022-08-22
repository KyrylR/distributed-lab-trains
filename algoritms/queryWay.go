package algoritms

import (
	"errors"
	"fmt"
	"sort"
	"time"
)

type trainPath struct {
	trains     []Train
	travelTime time.Duration
}

type trainToTrain struct {
	start      Train
	next       Train
	travelTime time.Duration
}

type QueryWay struct {
	processedData    *SortTrains
	trainPathForTime trainPath
}

func (p *trainPath) getTravelTime() {
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

func getLowestTravelTimeFromPath(trainPaths []trainPath) trainPath {
	if len(trainPaths) < 1 {
		return trainPath{}
	}
	less := func(i, j int) bool {
		return trainPaths[i].travelTime < trainPaths[j].travelTime
	}
	sort.Slice(trainPaths, less)
	return trainPaths[0]
}

func (w *QueryWay) getLowestCost() (float64, []int, error) {
	if w.processedData == nil {
		return 0, nil, errors.New("uninitialised data has been provided")
	}
	cost := 0.0
	trainPath := make([]int, len(w.processedData.Path.Way)-1)
	for _, trains := range w.processedData.Path.TrainMap {
		if len(trains) > 0 {
			cost += trains[0].Cost
			if w.processedData.Path.GetLastStation() != trains[0].DepartureStationId {
				trainPath[w.processedData.Path.GetStationIndex(trains[0].DepartureStationId)] = trains[0].TrainId
			}
		}
	}
	return cost, trainPath, nil
}

func (w *QueryWay) getLowestTime() (trainPath, error) {
	if w.processedData == nil {
		return trainPath{}, errors.New("uninitialised data has been provided")
	}
	return w.trainPathForTime, nil
}

func (w *QueryWay) initLowestTime() {
	links := w.buildStationLinks()
	startLinkKey := getLinkKey(w.processedData.Path.Way[0], w.processedData.Path.Way[1])
	trainPaths := make([]trainPath, 0)
	for _, ttt := range links[startLinkKey] {
		newTrainPath := w.newTrainPath(ttt)
		ok := w.completeTrainPath(links, &newTrainPath)
		if !ok {
			newTrainPath.trains = w.buildWayFromFastestTrains()
		}
		newTrainPath.getTravelTime()
		trainPaths = append(trainPaths, newTrainPath)
	}
	w.trainPathForTime = getLowestTravelTimeFromPath(trainPaths)
}

func (w *QueryWay) completeTrainPath(links map[string][]trainToTrain, trainPath *trainPath) bool {
	for i := 2; i < len(trainPath.trains); i++ {
		lastTrainInPath := trainPath.trains[i-1]
		nextLinkKey := getLinkKey(lastTrainInPath.DepartureStationId, lastTrainInPath.ArrivalStationId)
		nextLinkedTrain := findNextLinkedTrain(links, lastTrainInPath, nextLinkKey)
		if nextLinkedTrain.TrainId != -1 {
			trainPath.trains[i] = nextLinkedTrain
		} else {
			return false
		}
	}
	return true
}

func (w *QueryWay) buildWayFromFastestTrains() []Train {
	result := make([]Train, len(w.processedData.Path.Way)-1)
	i := 0
	for _, station := range w.processedData.Path.Way {
		if len(w.processedData.TravelTimeMap[station]) > 0 && w.processedData.Path.GetLastStation() != station {
			result[i] = w.processedData.TravelTimeMap[station][0]
			i += 1
		}
	}
	return result
}

func findNextLinkedTrain(links map[string][]trainToTrain, train Train, nextLinkKey string) Train {
	for _, value := range links[nextLinkKey] {
		if value.start.TrainId == train.TrainId {
			return value.next
		}
	}
	return Train{TrainId: -1}
}

func (w *QueryWay) newTrainPath(ttt trainToTrain) trainPath {
	newTrainPath := trainPath{}
	arr := make([]Train, len(w.processedData.Path.Way)-1)
	arr[0] = ttt.start
	arr[1] = ttt.next
	newTrainPath.trains = arr
	return newTrainPath
}

func (w *QueryWay) buildStationLinks() map[string][]trainToTrain {
	links := make(map[string][]trainToTrain)
	for _, trains := range w.processedData.TravelTimeMap {
		for _, train := range trains {
			for _, waitForTrain := range w.processedData.WaitingTimeMap[train.TrainId] {
				key := getLinkKey(train.DepartureStationId, waitForTrain.train.DepartureStationId)
				newTTT := newTrainToTrainStruct(train, waitForTrain.train, waitForTrain.waitingTime)
				links[key] = append(links[key], newTTT)
			}
		}
	}
	return links
}

func (w *QueryWay) String() string {
	path := fmt.Sprint(w.processedData.Path.Way)
	cost, costTrainPath, _ := w.getLowestCost()
	timeTravelPath, _ := w.getLowestTime()
	timeTrainIds := make([]int, len(timeTravelPath.trains))
	i := 0
	for _, train := range timeTravelPath.trains {
		timeTrainIds[i] = train.TrainId
		i += 1
	}
	timeTrainIdsStr := fmt.Sprint(timeTrainIds)
	stationPath := fmt.Sprintf("Station path: %v", path)
	costPath := fmt.Sprintf("Cost: %.2f -- TrainIds: %v", cost, fmt.Sprint(costTrainPath))
	timePath := fmt.Sprintf("Time: %v -- TrainIds: %v", timeTravelPath.travelTime, timeTrainIdsStr)
	return fmt.Sprintf("%v\n%v\n%v\n", stationPath, costPath, timePath)
}

func newTrainToTrainStruct(start, next Train, waitingTime time.Duration) trainToTrain {
	startTravelTime := SmoothOutTime(start.ArrivalTime.Sub(start.DepartureTime))
	nextTravelTime := SmoothOutTime(next.ArrivalTime.Sub(next.DepartureTime))
	res := trainToTrain{
		start:      start,
		next:       next,
		travelTime: startTravelTime + nextTravelTime + waitingTime,
	}
	return res
}

func getLinkKey(startId, nextId int) string {
	return fmt.Sprintf("%v -> %v", startId, nextId)
}
