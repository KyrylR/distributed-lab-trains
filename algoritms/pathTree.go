package algoritms

import "DistributedLab_Trains/apptypes"

type PathTree struct {
	DepartureId int
	Routes      []apptypes.Train
	Next        []*PathTree
}

func (tree *PathTree) buildPathTree(uniqueStations []int, mappedData map[int][]apptypes.Train) {
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
}

func (tree *PathTree) fillPathTreeNext(treeMap *map[int]*PathTree) {
	for _, arvStation := range getArrivalStations(&tree.Routes) {
		tree.Next = append(tree.Next, (*treeMap)[arvStation])
	}
}

func newPathTree(departure int, trains []apptypes.Train) *PathTree {
	newPathTree := PathTree{DepartureId: departure}
	newPathTree.Routes = trains
	return &newPathTree
}

func getArrivalStations(routes *[]apptypes.Train) []int {
	return GetUniqueStations(routes, true)
}
