package utils

import (
	"DistributedLab_Trains/apptypes"
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func ParseCsvToDataStruct(path string) ([]apptypes.Train, error) {
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

	var result []apptypes.Train
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var dataTmp apptypes.Train
		seperatedValues := strings.Split(re.ReplaceAllString(scanner.Text(), ""), ";")
		err := parseCsvLine(&dataTmp, seperatedValues)
		if err != nil {
			return nil, err
		}
		result = append(result, dataTmp)
	}
	return result, nil
}

func parseCsvLine(data *apptypes.Train, lines []string) error {
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
