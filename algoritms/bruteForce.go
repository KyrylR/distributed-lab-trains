package algoritms

import (
	"fmt"
)

func BuildPaths(data *[]Train) []Ways {
	uniqueStations := GetUniqueStations(data, false)
	mappedData := BuildMappedData(data)
	result := make([]Ways, 0)
	for _, station := range uniqueStations {
		pathTree := PathTree{DepartureId: station}
		treeMap := pathTree.buildPathTree(uniqueStations, mappedData)
		pathTrees := getAllPathTrees(&pathTree, &treeMap)
		ways := buildNewWaysFromPathTree(pathTrees, len(uniqueStations))
		for _, way := range ways.Ways {
			fmt.Println(way.Way)
		}
		result = append(result, ways)
	}
	return result
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
