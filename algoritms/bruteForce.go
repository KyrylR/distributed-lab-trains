package algoritms

import (
	"fmt"
	"sync"
)

// BuildPaths accepts the given data as an array of Train.
// Builds the corresponding PathTree structure and uses it to find all paths from station to station
// that can be reached with the corresponding trains.
// Returns a slice of Ways, containing all `possible paths'.
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

// BuildPathsGo is the same function as BuildPaths, but uses goroutines to speed up the search for all paths.
func BuildPathsGo(data *[]Train) []Ways {
	uniqueStations := GetUniqueStations(data, false)
	mappedData := BuildMappedData(data)
	result := make([]Ways, 0)

	var mutex sync.Mutex
	var wg sync.WaitGroup

	worker := func(station int) {
		wg.Add(1)
		defer wg.Done()

		pathTree := PathTree{DepartureId: station}
		treeMap := pathTree.buildPathTree(uniqueStations, mappedData)
		pathTrees := getAllPathTrees(&pathTree, &treeMap)
		ways := buildNewWaysFromPathTree(pathTrees, len(uniqueStations))

		mutex.Lock()
		result = append(result, ways)
		mutex.Unlock()
	}

	for _, station := range uniqueStations {
		go worker(station)
	}
	wg.Wait()
	return result
}

// buildNewWaysFromPathTree takes a PathTree slice and the number of unique stations in the data.
// Creates an instance of Ways and starts `offering' it each PathTree as many times as the number
// of unique stations in the data, starting with the specified one. At the end, clears each `incomplete' path and
// returns the Ways object.
func buildNewWaysFromPathTree(pathTrees []*PathTree, maxStations int) Ways {
	ways := Ways{}

	for i := 0; i < maxStations; i++ {
		for _, offer := range pathTrees {
			ways.proposeWay(offer, maxStations)
		}
	}

	ways.clear(maxStations)
	return ways
}

// getAllPathTrees takes a completed tree and its associated treeMap, returns a PathTree slice.
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
