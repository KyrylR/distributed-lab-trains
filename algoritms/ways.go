package algoritms

import (
	"reflect"
)

type PossibleWay struct {
	Way      []int
	TrainMap map[int][]Train
}

type Ways struct {
	Ways []*PossibleWay
}

func (ways *Ways) proposeWay(pathNode *PathTree, maxStations int) bool {
	addedWay, justAdded := true, false
	if len(ways.Ways) == 0 {
		for addedWay {
			addedWay = ways.addWay(PossibleWay{}, pathNode, maxStations)
		}
		return true
	}
	for _, way := range ways.Ways {
		tempWey := *way
		addedWay = way.proposeNewWay(pathNode, ways, maxStations)
		for addedWay {
			justAdded = true
			addedWay = ways.addWay(tempWey, pathNode, maxStations)
		}
	}
	return justAdded
}

func (ways *Ways) addWay(newWay PossibleWay, pathNode *PathTree, maxStations int) bool {
	if newWay.proposeNewWay(pathNode, ways, maxStations) {
		ways.Ways = append(ways.Ways, &newWay)
		return true
	}
	return false
}

func (ways Ways) isWayNeeded(newWay []int) bool {
	for _, way := range ways.Ways {
		if way.isEqual(newWay) {
			return false
		}
	}
	return true
}

func (ways *Ways) clear(maxStations int) {
	for i := 0; i < len(ways.Ways); i++ {
		if len(ways.Ways[i].Way) != maxStations {
			ways.Ways = append(ways.Ways[:i], ways.Ways[i+1:]...)
			i-- // Since we just deleted a[i], we must redo that index
		}
	}
}

func (w *PossibleWay) proposeNewWay(pathNode *PathTree, ways *Ways, maxStations int) bool {
	if maxStations <= len(w.Way) || maxStations <= 0 {
		return false
	}
	if len(w.Way) == 0 {
		w.Way = append(w.Way, pathNode.DepartureId)
		return w.addNewWay(pathNode, ways)
	}
	if w.getLastStation() == pathNode.DepartureId {
		return w.addNewWay(pathNode, ways)
	}
	return false
}

func (w *PossibleWay) addNewWay(pathNode *PathTree, ways *Ways) bool {
	for _, arrv := range pathNode.Next {
		if !IsStationAlreadyPassed(w.Way, arrv.DepartureId) {
			newWay := append(w.copy(), arrv.DepartureId)
			if ways.isWayNeeded(newWay) {
				w.Way = newWay
				w.addRoutes(pathNode, arrv.DepartureId)
				return true
			}
		}
	}
	return false
}

func (w PossibleWay) copy() []int {
	copiedWay := make([]int, len(w.Way))
	for idx, item := range w.Way {
		copiedWay[idx] = item
	}
	return copiedWay
}

func (w *PossibleWay) addRoutes(pathNode *PathTree, toStation int) {
	for _, route := range pathNode.Routes {
		if route.ArrivalStationId == toStation {
			if w.TrainMap == nil {
				w.TrainMap = make(map[int][]Train)
			}
			w.TrainMap[route.DepartureStationId] = append(w.TrainMap[route.DepartureStationId], route)
		}
	}
}

func (w PossibleWay) getLastStation() int {
	return w.Way[len(w.Way)-1]
}

func (w PossibleWay) isEqual(other []int) bool {
	return reflect.DeepEqual(w.Way, other)
}
