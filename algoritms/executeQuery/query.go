package executeQuery

import (
	"sort"

	"DistributedLab_Trains/algoritms/findPath"
)

// Query - stores a slice of the QueryWay structure.
type Query struct {
	allWays []QueryWay
}

// Initialize - takes a slice of Ways and executes all necessary functions
// to prepare the data for the desired queries.
func (q *Query) Initialize(ways []findPath.Ways) error {
	q.allWays = make([]QueryWay, 0)
	for _, possibleWays := range ways {
		for _, way := range possibleWays.Ways {
			newProcessedData := SortTrains{Path: way}
			err := newProcessedData.SortTimeAndCost()
			if err != nil {
				return err
			}
			newQueryWay := QueryWay{processedData: &newProcessedData}
			q.allWays = append(q.allWays, newQueryWay)
		}
	}
	for i := 0; i < len(q.allWays); i++ {
		q.allWays[i].initTrainPathForTime()
	}
	return nil
}

// sortByTime - uses the built-in sort function to order the slices by lowest time.
func (q *Query) sortByTime() {
	less := func(i, j int) bool {
		first, _ := q.allWays[i].getLowestTime()
		second, _ := q.allWays[j].getLowestTime()
		return first.travelTime < second.travelTime
	}
	sort.Slice(q.allWays, less)
}

// sortByCost - uses the built-in sort function to order the slices by lowest cost.
func (q *Query) sortByCost() {
	less := func(i, j int) bool {
		first, _, _ := q.allWays[i].getLowestCost()
		second, _, _ := q.allWays[j].getLowestCost()
		return first < second
	}
	sort.Slice(q.allWays, less)
}

// QueryByTime sorts paths by time and calls returnWays to return them.
func (q *Query) QueryByTime(recordsNum int) []QueryWay {
	q.sortByTime()
	return q.returnWays(recordsNum)
}

// QueryByCost sorts paths by cost and calls returnWays to return them.
func (q *Query) QueryByCost(recordsNum int) []QueryWay {
	q.sortByCost()
	return q.returnWays(recordsNum)
}

// returnWays - returns a given number of processed trains.
func (q *Query) returnWays(recordsNum int) []QueryWay {
	if recordsNum > len(q.allWays) {
		recordsNum = len(q.allWays)
	}
	result := make([]QueryWay, recordsNum)
	for i := 0; i < recordsNum; i++ {
		result[i] = q.allWays[i]
	}
	return result
}
