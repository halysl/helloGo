package db

import (
	"fmt"
	"strings"

	"github.com/halysl/hellogo/getmonitorinfo/getInfo"
)

func insertSQL(instanceData []getInfo.InstanceData) {
	valueFormatSQL := "('%s', '%s', '%s', '%s', '%s', '%s', '%s', '%s')"
	execSQL := "INSERT INTO `Test`.`test`(`target`, `uptime`, `nodeload1`, `nodeload5`, `nodeload15`, `node_cpu_usage`, `node_memory_total`, `node_memory_avaibable`) VALUES %s;"
	valueSQLList := make([]string, len(instanceData))
	for k, v := range instanceData {
		valueSQLList[k] = fmt.Sprintf(valueFormatSQL, v.Target, v.Uptime, v.Nodeload1, v.Nodeload5, v.Nodeload15, v.Node_cpu_usage, v.Node_memory_total, v.Node_memory_avaibable)
	}
	ExecSQL(fmt.Sprintf(execSQL, strings.Join(valueSQLList, ",")))
}

func updateSQL(instanceData []getInfo.InstanceData) {
	execSQL := "UPDATE `Test`.`test` SET uptime='%s', nodeload1='%s', nodeload5='%s', nodeload15='%s', node_cpu_usage='%s', node_memory_total='%s', node_memory_avaibable='%s' where target='%s'"
	execsqlList := make([]string, len(instanceData))
	for k, v := range instanceData {
		execsqlList[k] = fmt.Sprintf(execSQL, v.Uptime, v.Nodeload1, v.Nodeload5, v.Nodeload15, v.Node_cpu_usage, v.Node_memory_total, v.Node_memory_avaibable, v.Target)
	}
	for _, execSQL := range execsqlList {
		ExecSQL(exec_sql)
	}
}
