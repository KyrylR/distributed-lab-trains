package algoritms

import (
	"DistributedLab_Trains/apptypes"
	"fmt"
)

func BuildPaths(data *[]apptypes.Train) {
	uniqueStations := GetUniqueStations(data, false)
	mappedData := BuildMappedData(data)
	pathTree := PathTree{DepartureId: 1}
	pathTree.buildPathTree(uniqueStations, mappedData)
	fmt.Println(pathTree)
	//for _, station := range uniqueStations {
	//	result := buildRoutesFromDeparture(data, station)
	//	for _, posWay := range result {
	//		fmt.Println(posWay.Way)
	//	}
	//}
	//buildRoutesFromDeparture(data, 3)
}
