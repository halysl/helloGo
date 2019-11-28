package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
	"regexp"
	"strconv"
)

type Vmstat struct {
	Wired_Memory    int
	Active_Memory   int
	Inactive_Memory int
	Free_Memory     int
	//Real_Mem_Total int
}

func RunCmd(cmd string) string {
	exec_res := exec.Command(cmd)
	stdout, err := exec_res.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	defer stdout.Close()
	if err := exec_res.Start(); err != nil {
		log.Fatal(err)
	}
	opBytes, err := ioutil.ReadAll(stdout)
	if err != nil {
		log.Fatal(err)
	}
	return string(opBytes)
}

func parseInfo(info string) map[string]int64 {
	vm_info_map := make(map[string]int64)
	vm_info_regexp := regexp.MustCompile(`([\S ]+):\s+(\d+)`)
	all_res := vm_info_regexp.FindAllStringSubmatch(info, 100)
	for _, one := range all_res {
		value, err := strconv.ParseInt(one[2], 10, 32)
		if err != nil {
			log.Fatal(err)
			value = 0
		}
		vm_info_map[one[1]] = value
	}
	return vm_info_map
}

func (v *Vmstat) joinInfo(parseinfo map[string]int64) {
	var total int64
	for _, value := range parseinfo {
		total += value * 1024
	}
	//v.Real_Mem_Total = int(total/1024/1024)
	v.Active_Memory = int(parseinfo["Pages active"] / 1024)
	v.Free_Memory = int(parseinfo["Free Memory"] / 1024)
	v.Wired_Memory = int(parseinfo["Pages wired down"] / 1024)
	v.Inactive_Memory = int(parseinfo["Pages inactive"] / 1024)
}

func main() {
	cmd := "vm_stat"
	vm_info := RunCmd(cmd)
	parse_info := parseInfo(vm_info)
	v := Vmstat{}
	v.joinInfo(parse_info)
	fmt.Printf("Active_Memory:\t\t%d MB\nInactive_Memory:\t%d MB\nFree_Memory:\t\t%d MB\nWired_Memory:\t\t%d MB\n", v.Active_Memory, v.Inactive_Memory, v.Free_Memory, v.Wired_Memory)
}
