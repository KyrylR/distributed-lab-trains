package utils

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"DistributedLab_Trains/algoritms"
)

// ParseCsvToTrainStruct takes the path to a csv file in which entries are separated by `;` with the following
// structure: `train number; departure station; arrival station; cost; departure time; arrival time; arrival time`.
func ParseCsvToTrainStruct(path string) ([]algoritms.Train, error) {
	file, err := os.Open(path)
	if err != nil {
		log.Panicln("Open file error:", err)
		return nil, err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Panicln("Close file error:", err)
		}
	}(file)

	var re = regexp.MustCompile(`[^\d\;\:]`)

	var result []algoritms.Train
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var dataTmp algoritms.Train
		seperatedValues := strings.Split(re.ReplaceAllString(scanner.Text(), ""), ";")
		err := parseCsvLine(&dataTmp, seperatedValues)
		if err != nil {
			return nil, err
		}
		result = append(result, dataTmp)
	}
	return result, nil
}

// parseCsvLine helps to write the string of data into the Train structure.
func parseCsvLine(data *algoritms.Train, lines []string) error {
	conv, err := strconv.Atoi(lines[0])
	if err != nil {
		return err
	}
	data.TrainId = conv
	conv, err = strconv.Atoi(lines[1])
	if err != nil {
		return err
	}
	data.DepartureStationId = conv
	conv, err = strconv.Atoi(lines[2])
	if err != nil {
		return err
	}
	data.ArrivalStationId = conv
	convPrice, err := strconv.ParseFloat(lines[3], 64)
	if err != nil {
		return err
	}
	data.Cost = convPrice
	t, err := time.Parse("15:04:05", lines[4])
	if err != nil {
		return err
	}
	data.DepartureTime = t
	t, err = time.Parse("15:04:05", lines[5])
	if err != nil {
		return err
	}
	data.ArrivalTime = t
	return nil
}
