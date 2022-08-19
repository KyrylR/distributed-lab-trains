package algoritms

import (
	"fmt"
)

func BuildPaths(data *[]Train) {
	uniqueStations := GetUniqueStations(data, false)
	mappedData := BuildMappedData(data)
	//for _, station := range uniqueStations {
	for i := 0; i < 10000; i++ {

		pathTree := PathTree{DepartureId: 1902}
		treeMap := pathTree.buildPathTree(uniqueStations, mappedData)
		pathTrees := getAllPathTrees(&pathTree, &treeMap)
		ways := buildNewWaysFromPathTree(pathTrees, len(uniqueStations))
		//for _, way := range ways.Ways {
		//	fmt.Println(way.Way)
		//}
		fmt.Println(len(ways.Ways))
	}
	//}
}

func buildNewWaysFromPathTree(pathTrees []*PathTree, maxStations int) Ways {
	ways := Ways{}

	for i := 0; i < maxStations+1; i++ {
		for _, offer := range pathTrees {
			ways.proposeWay(offer, maxStations)
		}
	}

	ways.clear(maxStations)
	return ways
}

func getAllPathTrees(tree *PathTree, treeMap *map[int]*PathTree) []*PathTree {
	start := tree.DepartureId

	result := []*PathTree{tree}
	i := 1
	for k, v := range *treeMap {
		if k != start {
			result = append(result, v)
			i++
		}
	}

	return result
}
