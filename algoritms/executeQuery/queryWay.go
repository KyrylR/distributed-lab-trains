package executeQuery

import (
	"errors"
	"fmt"

	"DistributedLab_Trains/algoritms"
)

// QueryWay is used to make the last preparations before running queries.
type QueryWay struct {
	ProcessedData    *SortTrains
	TrainPathForTime TrainPath
}

// getLowestTime returns path and cost, with the lowest price to pass all stations.
// returns an error if uninitialised data was provided.
func (w *QueryWay) getLowestCost() (float64, []int, error) {
	if w.ProcessedData == nil {
		return 0, nil, errors.New("uninitialised data has been provided")
	}
	cost := 0.0
	trainPath := make([]int, len(w.ProcessedData.Path.Way)-1)
	for _, trains := range w.ProcessedData.Path.TrainMap {
		if len(trains) > 0 {
			cost += trains[0].Cost
			if w.ProcessedData.Path.GetLastStation() != trains[0].DepartureStationId {
				trainPath[w.ProcessedData.Path.GetStationIndex(trains[0].DepartureStationId)] = trains[0].TrainId
			}
		}
	}
	return cost, trainPath, nil
}

// getLowestTime returns the path that takes the shortest time to traverse all stations.
// returns an error if uninitialized data was provided.
func (w *QueryWay) getLowestTime() (TrainPath, error) {
	if w.ProcessedData == nil {
		return TrainPath{}, errors.New("uninitialised data has been provided")
	}
	return w.TrainPathForTime, nil
}

// initTrainPathForTime - finds the lowest path time for Path from the given processed data in the structure.
// Use the buildStationLinks function and link features to speed up the process.
// Finally, initialised the TrainPathForTime field in the QueryWay structure.
func (w *QueryWay) initTrainPathForTime() {
	links := w.buildStationLinks()
	startLinkKey := getLinkKey(w.ProcessedData.Path.Way[0], w.ProcessedData.Path.Way[1])
	trainPaths := make([]TrainPath, 0)
	for _, ttt := range links[startLinkKey] {
		newTrainPath := NewTrainPath(ttt, w.ProcessedData.Path.Way)
		ok := w.completeTrainPath(links, &newTrainPath)
		if !ok {
			newTrainPath.trains = w.buildWayFromFastestTrains()
		}
		newTrainPath.initTravelTimeField()
		trainPaths = append(trainPaths, newTrainPath)
	}
	w.TrainPathForTime = GetLowestTravelTimeForPath(trainPaths)
}

// completeTrainPath uses the reference structure to create a complete TrainPath,
// returns true if everything was found, otherwise returns false.
func (w *QueryWay) completeTrainPath(links map[string][]TrainToTrain, trainPath *TrainPath) bool {
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

// buildWayFromFastestTrains simply takes the fastest trains from station to station
// and creates a slice of such trains.
func (w *QueryWay) buildWayFromFastestTrains() []algoritms.Train {
	result := make([]algoritms.Train, len(w.ProcessedData.Path.Way)-1)
	i := 0
	for _, station := range w.ProcessedData.Path.Way {
		if len(w.ProcessedData.TravelTimeMap[station]) > 0 && w.ProcessedData.Path.GetLastStation() != station {
			result[i] = w.ProcessedData.TravelTimeMap[station][0]
			i += 1
		}
	}
	return result
}

// findNextLinkedTrain returns the next linked train in the link structure for the given key,
// otherwise it returns an empty train.
func findNextLinkedTrain(links map[string][]TrainToTrain, train algoritms.Train, nextLinkKey string) algoritms.Train {
	for _, value := range links[nextLinkKey] {
		if value.start.TrainId == train.TrainId {
			return value.next
		}
	}
	return algoritms.Train{TrainId: -1}
}

// buildStationLinks creates links between stations using the TrainToTrain structure.
// The key is provided by getLinkKey function; for example: `1929 -> 1902` - string value, where first value is
// departure station and second value is arrival station.
func (w *QueryWay) buildStationLinks() map[string][]TrainToTrain {
	links := make(map[string][]TrainToTrain)
	for _, trains := range w.ProcessedData.TravelTimeMap {
		for _, train := range trains {
			for _, waitForTrain := range w.ProcessedData.WaitingTimeMap[train.TrainId] {
				key := getLinkKey(train.DepartureStationId, waitForTrain.train.DepartureStationId)
				newTTT := NewTrainToTrain(train, waitForTrain.train, waitForTrain.waitingTime)
				links[key] = append(links[key], newTTT)
			}
		}
	}
	return links
}

// getLinkKey returns the key value for the following construct `links := make(map[string][]TrainToTrain)`,
// which is used in the buildStationLinks function.
func getLinkKey(startId, nextId int) string {
	return fmt.Sprintf("%v -> %v", startId, nextId)
}

// String implements Stringer interface. Used to display query information about Path nicely.
func (w *QueryWay) String() string {
	path := fmt.Sprint(w.ProcessedData.Path.Way)
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
	timePath := fmt.Sprintf("Time: %v -- TrainIds: %v", timeTravelPath.TravelTime, timeTrainIdsStr)
	return fmt.Sprintf("%v\n%v\n%v\n", stationPath, costPath, timePath)
}
