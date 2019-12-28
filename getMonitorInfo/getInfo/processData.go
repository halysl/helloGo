package getInfo

import (
	"encoding/json"
	"log"
)

type targetResult struct {
	Instance string
	Job      string
	Health   string
}

type target struct {
	Labels struct {
		Instance string
		Job      string
	}
	DiscoveredLabels struct {
		Address      string `json:"__address__"`
		Metrics_path string `json:"__metrics_path__"`
		Schema       string `json:"__scheme__"`
		Job          string
	}
	Scrape     string `json:"scrapeUrl"`
	LastScrape string
	Health     string
	LastError  string
}

type originTargetInfo struct {
	Status string
	Data   struct {
		ActiveTargets  []target
		DroppedTargets []target
	}
}

func ProcessTargetInfo() []targetResult {
	var targetResultList []targetResult

	var tI originTargetInfo
	targetInfo, _ := GetTargetInfo()
	err := json.Unmarshal(targetInfo, &tI)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range tI.Data.ActiveTargets {
		if v.Labels.Instance == "localhost:9090" {
			continue
		} else {
			targetResultList = append(targetResultList, targetResult{
				Instance: v.Labels.Instance,
				Job:      v.Labels.Job,
				Health:   v.Health,
			})
		}
	}

	return targetResultList
}

type MetricResult struct {
	Name           string
	MetricInfoList []metricInfo
}

type metricInfo struct {
	Instance string
	Job      string
	Value    string
}

type originMetricInfo struct {
	Status string
	Data   struct {
		ResultType string
		Result     []struct {
			Metric struct {
				Instance string
				Job      string
			}
			Value []interface{}
		}
	}
}

func ProcessQuery(query, name, t string) (mr MetricResult) {
	mr.Name = name

	metricData, _ := GetQueryInfo(query, t)

	var mI originMetricInfo
	err := json.Unmarshal(metricData, &mI)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range mI.Data.Result {
		queryValue := ""
		for _, val := range v.Value {
			switch val.(type) {
			case string:
				queryValue = val.(string)
				break
			}
		}
		mr.MetricInfoList = append(mr.MetricInfoList, metricInfo{
			Instance: v.Metric.Instance,
			Job:      v.Metric.Job,
			Value:    queryValue,
		})
	}
	return
}

func ProcessUptime(query, name, t string) (mr MetricResult) {
	mr = ProcessQuery(query, name, t)
	for k, v := range mr.MetricInfoList {
		mr.MetricInfoList[k].Value = parseTimeByHour(t, v)
	}
	return
}

func ProcessNodeMemory(query, name, t string) (mr MetricResult) {
	mr = ProcessQuery(query, name, t)

	for k, v := range mr.MetricInfoList {
		mr.MetricInfoList[k].Value = parseUnitToGB(v)
	}
	return
}
