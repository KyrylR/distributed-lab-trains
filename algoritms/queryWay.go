package algoritms

import "errors"

type QueryWay struct {
	processedData *SortTrains
}

func (w *QueryWay) getLowestCost() (float64, error) {
	if w.processedData == nil {
		return 0, errors.New("uninitialised data has been provided")
	}
	cost := 0.0
	for _, trains := range w.processedData.Path.TrainMap {
		if len(trains) > 0 {
			cost += trains[0].Cost
		}
	}
	return cost, nil
}
