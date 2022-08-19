package algoritms

import "DistributedLab_Trains/apptypes"

// IsStationAlreadyPassed takes a slice int' and a value to check if this value belongs to the slice
func IsStationAlreadyPassed(path []int, arrival int) bool {
	for _, point := range path {
		if point == arrival {
			return true
		}
	}
	return false
}

// BuildMappedData takes all train data and creates a map, where key is the
// departure station ID and value is a slice of all available trains for departure from that station.
func BuildMappedData(data *[]apptypes.Train) map[int][]apptypes.Train {
	result := make(map[int][]apptypes.Train)
	for _, train := range *data {
		result[train.DepartureStationId] = append(result[train.DepartureStationId], train)
	}
	return result
}

// GetUniqueStations accepts all data about trains and produce an
// output that contains all stations that is available from departure or in arrival.
// If `onlyArrival` is specified, the output will only contain possible arrival stations.
func GetUniqueStations(data *[]apptypes.Train, onlyArrival bool) []int {
	unique := map[int]bool{}
	for _, train := range *data {
		unique[train.ArrivalStationId] = true
		if !onlyArrival {
			unique[train.DepartureStationId] = true
		}
	}

	result := make([]int, len(unique))
	i := 0

	for k := range unique {
		result[i] = k
		i++
	}

	return result
}
