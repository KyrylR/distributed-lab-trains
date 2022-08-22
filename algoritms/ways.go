package algoritms

import (
	"reflect"
)

// PossibleWay represents the path from station to station and the associated
// trains that can be used to get to the desired stations.
type PossibleWay struct {
	Way      []int
	TrainMap map[int][]Train // key - station identifier, value - array of similar trains.
}

// Ways represents an array of PossibleWay references.
type Ways struct {
	Ways []*PossibleWay
}

// proposeWay as function is called, it `proposes` a given PathTree object to every PossibleWay in the Ways structure.
// Returns FALSE if there are no possible paths in which the ID of the last station matches the DepartureId field of
// the PathTree object, otherwise it tries to add a new path the required number of times and returns TRUE at the end.
func (ways *Ways) proposeWay(pathNode *PathTree, maxStations int) bool {
	addedWay, justAdded := true, false
	if len(ways.Ways) == 0 {
		for addedWay {
			addedWay = ways.addWay(PossibleWay{}, pathNode, maxStations)
		}
		return true
	}
	for _, way := range ways.Ways {
		tempWey := way.fullCopy()
		addedWay = way.proposeNewWay(pathNode, ways, maxStations)
		for addedWay {
			justAdded = true
			addedWay = ways.addWay(tempWey, pathNode, maxStations)
		}
	}
	return justAdded
}

// addWay takes an instance of PossibleWay, a reference to pathTree and the number of unique stations in the data.
// Calls proposeNewWay function, and if it returns true, we add new path to the array and return true,
// otherwise we return false.
func (ways *Ways) addWay(newWay PossibleWay, pathNode *PathTree, maxStations int) bool {
	if newWay.proposeNewWay(pathNode, ways, maxStations) {
		ways.Ways = append(ways.Ways, &newWay)
		return true
	}
	return false
}

// clear - removes any way that do not contain all stations
func (ways *Ways) clear(maxStations int) {
	for i := 0; i < len(ways.Ways); i++ {
		if len(ways.Ways[i].Way) != maxStations {
			ways.Ways = append(ways.Ways[:i], ways.Ways[i+1:]...)
			i-- // Since we just deleted a[i], we must redo that index
		}
	}
}

// isWayNeeded returns true if the provided array of station IDs is unique in the current PossibleWay array
// of the Ways structure, otherwise it returns false.
func (ways *Ways) isWayNeeded(newWay []int) bool {
	for _, way := range ways.Ways {
		if way.isEqual(newWay) {
			return false
		}
	}
	return true
}

// proposeNewWay checks for path constraints, if all is OK, calls addNewWay function, otherwise returns false.
func (w *PossibleWay) proposeNewWay(pathNode *PathTree, ways *Ways, maxStations int) bool {
	if maxStations <= len(w.Way) || maxStations <= 0 {
		return false
	}
	if len(w.Way) == 0 {
		w.Way = append(w.Way, pathNode.DepartureId)
		return w.addNewWay(pathNode, ways)
	}
	if w.GetLastStation() == pathNode.DepartureId {
		return w.addNewWay(pathNode, ways)
	}
	return false
}

// addNewWay takes reference to PathTree, where field DepartureId corresponds to the last station in the path;
// and a reference to the actual Ways. It checks if the station has already been passed
// and if its path in Ways is unique. If the new path is unique in the current array of stations
// and unique in the PossibleWay array in the Ways structure, it is added to the current array and returns true,
// otherwise false is returned.
func (w *PossibleWay) addNewWay(pathNode *PathTree, ways *Ways) bool {
	for _, arrival := range pathNode.Next {
		if !IsStationAlreadyPassed(w.Way, arrival.DepartureId) {
			newWay := append(w.copy(), arrival.DepartureId)
			if ways.isWayNeeded(newWay) {
				w.Way = newWay
				w.addRoutes(pathNode, arrival.DepartureId)
				return true
			}
		}
	}
	return false
}

// addRoutes searches all trains that are in the Routes' field of the PathTree structure;
// fills TrainMap field in PossibleWay with those trains that have the appropriate arrival station ID (toStation).
func (w *PossibleWay) addRoutes(pathNode *PathTree, toStation int) {
	for _, route := range pathNode.Routes {
		if route.ArrivalStationId == toStation && pathNode.DepartureId == route.DepartureStationId {
			if w.TrainMap == nil {
				w.TrainMap = make(map[int][]Train)
			}
			w.TrainMap[route.DepartureStationId] = append(w.TrainMap[route.DepartureStationId], route)
		}
	}
}

// DeleteRouteDuplicates uses RemoveDuplicates function to clear TrainMap of duplicates.
func (w *PossibleWay) DeleteRouteDuplicates() {
	for _, station := range w.Way {
		w.TrainMap[station] = RemoveDuplicates(w.TrainMap[station])
	}
}

// copy - makes a copy of Way array
func (w *PossibleWay) copy() []int {
	copiedWay := make([]int, len(w.Way))
	for idx, item := range w.Way {
		copiedWay[idx] = item
	}
	return copiedWay
}

// fullCopy - makes full a copy of PossibleWay structure
func (w *PossibleWay) fullCopy() PossibleWay {
	var copiedPossibleWay PossibleWay
	copiedWay := make([]int, len(w.Way))
	copiedTrainMap := make(map[int][]Train)
	for idx, item := range w.Way {
		copiedWay[idx] = item
	}

	for key, value := range w.TrainMap {
		copiedTrainMap[key] = DeepTrainSliceCopy(value)
	}

	copiedPossibleWay.Way = copiedWay
	copiedPossibleWay.TrainMap = copiedTrainMap
	return copiedPossibleWay
}

// isEqual returns true if the station path is the same, otherwise returns false
func (w *PossibleWay) isEqual(other []int) bool {
	return reflect.DeepEqual(w.Way, other)
}

// GetLastStation returns the ID of the last station in the Way
func (w *PossibleWay) GetLastStation() int {
	return w.Way[len(w.Way)-1]
}

// GetNextStation returns the next station after the given one,
// returns -1 if the given station is the last one on the path.
func GetNextStation(currentStation int, way *PossibleWay) int {
	if currentStation == way.GetLastStation() {
		return -1
	}
	for idx, station := range way.Way {
		if station == currentStation {
			return way.Way[idx+1]
		}
	}
	return -1
}
