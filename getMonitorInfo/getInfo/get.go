package getInfo

import (
	"strconv"
	"time"
)

type InstanceData struct {
	Target                string
	Uptime                string
	Nodeload1             string
	Nodeload5             string
	Nodeload15            string
	Node_cpu_usage        string
	Node_memory_total     string
	Node_memory_avaibable string
}

func GetOriginInfo() []InstanceData {
	var metric = make(map[string]string)
	metric["node_boot_time"] = "node_boot_time_seconds"
	metric["node_load1"] = "node_load1"
	metric["node_load5"] = "node_load5"
	metric["node_load15"] = "node_load15"
	metric["node_cpu"] = "round((1%20-%20avg(irate(node_cpu_seconds_total%7Bmode%3D%22idle%22%7D%5B30m%5D))%20by%20(instance))*100)"
	metric["node_memory_total"] = "node_memory_MemTotal_bytes"
	metric["node_memory_available"] = "node_memory_MemAvailable_bytes"

	nowTimestamp := strconv.FormatInt(time.Now().Unix(), 10)

	all_target := ProcessTargetInfo()
	lenTarget := len(all_target)
	instanceList := make([]InstanceData, lenTarget)

	all_uptime := ProcessUptime(metric["node_boot_time"], "node_boot_time", nowTimestamp)
	all_nodeload1 := ProcessQuery(metric["node_load1"], "node_load1", nowTimestamp)
	all_nodeload5 := ProcessQuery(metric["node_load5"], "node_load5", nowTimestamp)
	all_nodeload15 := ProcessQuery(metric["node_load15"], "node_load15", nowTimestamp)
	all_nodeCpu := ProcessQuery(metric["node_cpu"], "node_cpu", nowTimestamp)
	all_nodeMemoryTotal := ProcessNodeMemory(metric["node_memory_total"], "node_memory_total", nowTimestamp)
	all_nodeMemoryAvaiblable := ProcessNodeMemory(metric["node_memory_available"], "node_memory_available", nowTimestamp)

	//var metricResultMap map[string]MetricResult
	//metricResultMap = make(map[string]MetricResult)
	//metricResultMap["node_boot_time"] = ProcessUptime(metric["node_boot_time"], "node_boot_time", nowTimestamp)
	//metricResultMap["node_load1"] = ProcessQuery(metric["node_load1"], "node_load1", nowTimestamp)
	//metricResultMap["node_load5"] = ProcessQuery(metric["node_load5"], "node_load5", nowTimestamp)
	//metricResultMap["node_load15"] = ProcessQuery(metric["node_load15"], "node_load15", nowTimestamp)
	//metricResultMap["node_cpu"] = ProcessQuery(metric["node_cpu"], "node_cpu", nowTimestamp))
	//metricResultMap["node_memory_total"] = ProcessNodeMemory(metric["node_memory_total"], "node_memory_total", nowTimestamp)
	//metricResultMap["node_memory_available"] = ProcessNodeMemory(metric["node_memory_available"], "node_memory_available", nowTimestamp)




	for k, v := range all_target {
		instanceList[k].Target = v.Instance
	}

	for k, v := range instanceList {
		for _, temp := range all_uptime.MetricInfoList {
			if temp.Instance == v.Target {
				v.Uptime = temp.Value
				instanceList[k] = v
				break
			}

		}
		for _, temp := range all_nodeload1.MetricInfoList {
			if temp.Instance == v.Target {
				v.Nodeload1 = temp.Value
				instanceList[k] = v
				break
			}
		}
		for _, temp := range all_nodeload5.MetricInfoList {
			if temp.Instance == v.Target {
				v.Nodeload5 = temp.Value
				instanceList[k] = v
				break
			}
		}
		for _, temp := range all_nodeload15.MetricInfoList {
			if temp.Instance == v.Target {
				v.Nodeload15 = temp.Value
				instanceList[k] = v
				break
			}
		}
		for _, temp := range all_nodeCpu.MetricInfoList {
			if temp.Instance == v.Target {
				v.Node_cpu_usage = temp.Value
				instanceList[k] = v
				break
			}
		}
		for _, temp := range all_nodeMemoryTotal.MetricInfoList {
			if temp.Instance == v.Target {
				v.Node_memory_total = temp.Value
				instanceList[k] = v
				break
			}
		}
		for _, temp := range all_nodeMemoryAvaiblable.MetricInfoList {
			if temp.Instance == v.Target {
				v.Node_memory_avaibable = temp.Value
				instanceList[k] = v
				break
			}
		}
	}
	return instanceList
}
