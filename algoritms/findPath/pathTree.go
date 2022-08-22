package findPath

import (
	"DistributedLab_Trains/algoritms"
)

// PathTree represents a unique station in the data
type PathTree struct {
	DepartureId int
	Routes      []algoritms.Train // Routes contain information about each Train departing from this station
	Next        []*PathTree       // Next - slice of stations that can be reached from the current one.
}

// BuildPathTree fills each field in the PathTree structure with received information about unique stations,
// and all stations in map format, where key is the source station ID and value is a slice of
// the corresponding `Train` structures.
func (tree *PathTree) BuildPathTree(uniqueStations []int, mappedData map[int][]algoritms.Train) map[int]*PathTree {
	tree.Routes = mappedData[tree.DepartureId]

	treeMap := make(map[int]*PathTree)
	for _, station := range uniqueStations {
		if station != tree.DepartureId {
			treeMap[station] = newPathTree(station, mappedData[station])
		} else {
			treeMap[station] = tree
		}
	}

	for _, pathTree := range treeMap {
		pathTree.fillPathTreeNext(&treeMap)
	}
	return treeMap
}

// fillPathTreeNext takes treeMap where key is the id of the departure station and value is the tree itself.
// Uses treeMap to fill the `Next` field in the pathTree structure.
func (tree *PathTree) fillPathTreeNext(treeMap *map[int]*PathTree) {
	for _, arvStation := range getArrivalStations(&tree.Routes) {
		tree.Next = append(tree.Next, (*treeMap)[arvStation])
	}
}

// newPathTree returns a new PathTree structure element with the given departure ID and associated trains.
func newPathTree(departure int, trains []algoritms.Train) *PathTree {
	newPathTree := PathTree{DepartureId: departure}
	newPathTree.Routes = trains
	return &newPathTree
}

// getArrivalStations returns a slice of station IDs to which provided trains can reach.
func getArrivalStations(routes *[]algoritms.Train) []int {
	return GetUniqueStations(routes, true)
}
