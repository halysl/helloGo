package getInfo

import (
	"log"
	"strconv"
	"time"
)

func parseTimeByHour(t string, v metricInfo) string {
	now64, err := strconv.ParseInt(t, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	nowts := time.Unix(now64, 0)
	start64, err := strconv.ParseInt(v.Value, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	startts := time.Unix(start64, 0)
	d := nowts.Sub(startts)
	return d.String()
}


func parseUnitToGB(v metricInfo) string {
	temp, _ := strconv.ParseFloat(v.Value, 64)
	temp = temp / 1024 / 1024 / 1024
	temp2str := strconv.FormatFloat(temp, 'f', 2, 64)
	return temp2str
}

