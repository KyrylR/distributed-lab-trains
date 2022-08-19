package algoritms

import "reflect"

type PossibleWay struct {
	Way      []int
	TrainMap map[int][]Train
}

type Ways struct {
	Ways []PossibleWay
}

func (ways *Ways) isWayNeeded(newWay []int) bool {
	for _, way := range ways.Ways {
		if way.isEqual(newWay) {
			return false
		}
	}
	return true
}

func (w *PossibleWay) proposeNewWay(pathNode *PathTree, ways Ways, maxStations int) bool {
	if maxStations >= len(w.Way) || maxStations <= 0 {
		return false
	}
	if len(w.Way) == 0 {
		w.Way = append(w.Way, pathNode.DepartureId)
		return w.addWay(pathNode, ways)
	}
	if w.getLastStation() == pathNode.DepartureId {
		return w.addWay(pathNode, ways)
	}
	return false
}

func (w *PossibleWay) addWay(pathNode *PathTree, ways Ways) bool {
	for _, arrv := range pathNode.Next {
		if !IsStationAlreadyPassed(w.Way, arrv.DepartureId) {
			newWay := append(w.Way, pathNode.DepartureId)
			if ways.isWayNeeded(newWay) {
				w.Way = newWay
				w.addRoutes(pathNode, arrv.DepartureId)
				return true
			}
		}
	}
	return false
}

func (w *PossibleWay) addRoutes(pathNode *PathTree, toStation int) {
	for _, route := range pathNode.Routes {
		if w.getLastStation() == toStation {
			w.TrainMap[route.DepartureStationId] = append(w.TrainMap[route.DepartureStationId], route)
		}
	}
}

func (w *PossibleWay) getLastStation() int {
	return w.Way[len(w.Way)-1]
}

func (w *PossibleWay) isEqual(other []int) bool {
	return reflect.DeepEqual(w.Way, other)
}
