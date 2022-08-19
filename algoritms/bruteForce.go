package algoritms

import (
	"DistributedLab_Trains/apptypes"
	"fmt"
)

func BuildPaths(data *[]apptypes.Train) {
	//uniqueStations := getUniqueStations(data)
	//for _, station := range uniqueStations {
	//	result := buildRoutesFromDeparture(data, station)
	//	for _, posWay := range result {
	//		fmt.Println(posWay.Way)
	//	}
	//}
	buildRoutesFromDeparture(data, 3)
}

func buildRoutesFromDeparture(data *[]apptypes.Train, departure int) []*apptypes.PossibleWay {
	result := make([]*apptypes.PossibleWay, 0)
	oldResult := make([]apptypes.PossibleWay, 0)
	stationNum, pathTree := findFromDeparture(data, departure)

	pathTrees := make([]*apptypes.PathTree, 0)
	pathTrees = append(pathTrees, &pathTree)

	for i := 0; i < stationNum; i++ {
		var nextPathTrees = make([]*apptypes.PathTree, 0)
		usedTickets := make(map[string]bool)
		for _, station := range pathTrees {
			for _, train := range station.Routes {
				pathStr := getPathString(train.DepartureStationId, train.ArrivalStationId)
				if _, ok := usedTickets[pathStr]; !ok {
					addTrainToPath(data, &oldResult, &result, train, stationNum)
					usedTickets[pathStr] = true
				}
			}
			nextPathTrees = append(nextPathTrees, *station.Next...)
		}
		oldResult = make([]apptypes.PossibleWay, 0)
		for _, item := range result {
			newPosWay := *item
			newPosWay.TrainMap = make(map[int][]apptypes.Train)
			for k, v := range item.TrainMap {
				newPosWay.TrainMap[k] = v
			}

			oldResult = append(oldResult, newPosWay)
		}
		pathTrees = nextPathTrees
	}
	resultFiltered := make([]*apptypes.PossibleWay, 0)
	for _, item := range result {
		if len(item.Way) == stationNum {
			resultFiltered = append(resultFiltered, item)
		}
	}
	return resultFiltered
}

func getPathString(dep, arv int) string {
	return fmt.Sprintf("%v -> %v", dep, arv)
}

func addTrainToPath(
	data *[]apptypes.Train,
	oldPaths *[]apptypes.PossibleWay,
	newPaths *[]*apptypes.PossibleWay,
	train apptypes.Train,
	stationNum int) {
	if len(*newPaths) == 0 {
		newPosWay := buildNewPossibleWay(data, train, nil, stationNum)
		*newPaths = append(*newPaths, &newPosWay)
		return
	}
	newPathNeeded := true
	for _, path := range *newPaths {
		if path.LastArrival == train.DepartureStationId {
			if !isAlreadyPassedStation(path.Way, train.ArrivalStationId) {
				newPathNeeded = false
				allEqualTickets := findAllEqualTickets(data, train)
				path.Way = append(path.Way, train.ArrivalStationId)
				path.LastArrival = train.ArrivalStationId
				path.TrainMap[train.DepartureStationId] = append(path.TrainMap[train.DepartureStationId], allEqualTickets...)
				break
			}
		}
	}
	if newPathNeeded {
		newPosWay := buildNewPossibleWay(data, train, oldPaths, stationNum)
		if newPosWay.LastArrival != -1 {
			*newPaths = append(*newPaths, &newPosWay)
		}
	}
}

func findAllEqualTickets(data *[]apptypes.Train, train apptypes.Train) []apptypes.Train {
	result := make([]apptypes.Train, 0)
	result = append(result, train)
	for _, compareTrain := range *data {
		if compareTrain.TrainId != train.TrainId &&
			compareTrain.DepartureStationId == train.DepartureStationId &&
			compareTrain.ArrivalStationId == train.ArrivalStationId {
			result = append(result, compareTrain)
		}
	}
	return result
}

func isAlreadyPassedStation(path []int, arrival int) bool {
	for _, point := range path {
		if point == arrival {
			return true
		}
	}
	return false
}

func buildNewPossibleWay(data *[]apptypes.Train,
	train apptypes.Train,
	oldPaths *[]apptypes.PossibleWay,
	stationNum int) apptypes.PossibleWay {
	if oldPaths == nil || (oldPaths != nil && len(*oldPaths) == 0) {
		newPosWay := apptypes.PossibleWay{TrainMap: make(map[int][]apptypes.Train)}
		newPosWay.LastArrival = train.ArrivalStationId
		newPosWay.Way = append(newPosWay.Way, train.DepartureStationId, train.ArrivalStationId)
		allEqualTickets := findAllEqualTickets(data, train)
		newPosWay.TrainMap[train.DepartureStationId] = append(newPosWay.TrainMap[train.DepartureStationId], allEqualTickets...)
		return newPosWay
	}
	for _, path := range *oldPaths {
		if path.LastArrival == train.DepartureStationId && stationNum != len(path.Way) {
			if !isAlreadyPassedStation(path.Way, train.ArrivalStationId) {
				path.Way = append(path.Way, train.ArrivalStationId)
				path.LastArrival = train.ArrivalStationId
				allEqualTickets := findAllEqualTickets(data, train)
				path.TrainMap[train.DepartureStationId] = append(path.TrainMap[train.DepartureStationId], allEqualTickets...)
				return path
			}
		}
	}
	return apptypes.PossibleWay{LastArrival: -1}
}

func findFromDeparture(data *[]apptypes.Train, departure int) (int, apptypes.PathTree) {
	uniqueStations := getUniqueStations(data)
	destinationMap := buildMappedData(data)

	var (
		pathTree         = apptypes.PathTree{Id: -1, ArrivalStations: make(map[int]bool)}
		previousStations = make(map[int]*apptypes.PathTree)
		nextStations     = make(map[int]*apptypes.PathTree)

		countUniqueStations = len(uniqueStations)
		departures          = []int{departure}
	)

	// To DO if there is no depId in uniqueStations

	for i := 0; i <= countUniqueStations; i++ {
		visited := make(map[int]bool)
		for _, dep := range departures {
			if dep == -1 {
				continue
			}
			visited[dep] = true
			for _, train := range destinationMap[dep] {
				possibleNewStation := train.ArrivalStationId
				if _, ok := visited[possibleNewStation]; !ok && possibleNewStation != departure {
					visited[possibleNewStation] = true
				}

				if possibleNewStation != departure {
					addTrainToPathTree(&pathTree, &nextStations, &previousStations, train)
				}
			}
		}

		departures = []int{}
		for k := range visited {
			if k != departure {
				departures = append(departures, k)
			}
		}
		previousStations = nextStations
		nextStations = make(map[int]*apptypes.PathTree)
	}
	return countUniqueStations, pathTree
}

func addTrainToPathTree(pathTree *apptypes.PathTree,
	newLastPoint *map[int]*apptypes.PathTree,
	oldLastPoint *map[int]*apptypes.PathTree,
	train apptypes.Train) {

	if pathTree.Id == -1 {
		pathTree.Id = 1
		pathTree.DepartureId = train.DepartureStationId
		pathTree.ArrivalStations[train.ArrivalStationId] = true
		pathTree.Routes = append(pathTree.Routes, train)
		pathTreeArr := make([]*apptypes.PathTree, 0)
		pathTree.Next = &pathTreeArr
		(*newLastPoint)[train.DepartureStationId] = pathTree
		return
	}
	newPathTree := apptypes.PathTree{Id: -1, ArrivalStations: make(map[int]bool)}
	if lastPathTree, ok := (*newLastPoint)[train.DepartureStationId]; ok {
		lastPathTree.Routes = append(lastPathTree.Routes, train)
		lastPathTree.ArrivalStations[train.ArrivalStationId] = true
	} else {
		newPathTree.Id = 1
		newPathTree.DepartureId = train.DepartureStationId
		pathTreeArr := make([]*apptypes.PathTree, 0)
		newPathTree.Next = &pathTreeArr

		for _, prevPathTree := range *oldLastPoint {
			if _, ok := prevPathTree.ArrivalStations[train.DepartureStationId]; ok {
				if prevPathTree.DepartureId != train.ArrivalStationId {
					newPathTree.Routes = append(newPathTree.Routes, train)
					newPathTree.ArrivalStations[train.ArrivalStationId] = true
				}
				*(*prevPathTree).Next = append(*(*prevPathTree).Next, &newPathTree)
			}
		}
		(*newLastPoint)[train.DepartureStationId] = &newPathTree
	}
}

func buildMappedData(data *[]apptypes.Train) map[int][]apptypes.Train {
	result := make(map[int][]apptypes.Train)
	for _, train := range *data {
		result[train.DepartureStationId] = append(result[train.DepartureStationId], train)
	}
	return result
}

func getUniqueStations(data *[]apptypes.Train) []int {
	unique := map[int]bool{}
	for _, train := range *data {
		unique[train.ArrivalStationId] = true
	}

	result := make([]int, len(unique))
	i := 0

	for k := range unique {
		result[i] = k
		i++
	}

	return result
}
