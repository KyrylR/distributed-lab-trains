package tests

import (
	"testing"

	"DistributedLab_Trains/algoritms"
	"DistributedLab_Trains/utils"
)

func TestParseCsvLine(t *testing.T) {
	table := []struct {
		lines []string
		ok    bool
	}{
		{[]string{"1", "1", "1", "1.0", "19:00:12", "19:00:12"}, true},
		{[]string{"12222", "34", "123123", "1.0", "19:00:12", "19:00:12"}, true},
		{[]string{"1sd", "1", "1", "1.0", "19:00:12", "19:00:12"}, false},
		{[]string{"1", "1asd", "1", "1.0", "19:00:12", "19:00:12"}, false},
		{[]string{"12", "13", "1", "1.231", "00:00:12", "00:00:22"}, true},
		{[]string{"1", "1", "1dsf", "1,0", "19:00:12", "19:00:12"}, false},
		{[]string{"132", "1", "1", "1.0", "19;00:12", "19:00:12"}, false},
		{[]string{"1", "1", "1", "1.0", "19:00:12", "19:00:312"}, false},
		{[]string{"1", "1", "1", "1.0", "19:00:12", "129:00:12"}, false},
		{[]string{"1", "1", "1", "1.0", "19:00:12", "129:00:12", "dd"}, false},
		{[]string{"1", "1", "1", "1.0", "19:00:12"}, false},
	}

	for _, data := range table {
		err := utils.ParseCsvLine(&algoritms.Train{}, data.lines)
		if data.ok && err != nil {
			t.Errorf("%v: %v, error should be nil", data.lines, err)
		}
	}
}
