package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/xerrors"
	"io"
	"io/ioutil"
	"os"
	"sort"
	"text/template"
)

const renderTaskTemplate = `{{- if .TotalTaskGlanceInfo -}}
--------------------------------------------------------------------------------------------------------------------------------
Task Glance Info
{{- range .TotalTaskGlanceInfo -}}
{{ . }}
{{- end }}
--------------------------------------------------------------------------------------------------------------------------------
{{- end }}

{{ if .TaskDetailMap }}
--------------------------------------------------------------------------------------------------------------------------------
Task Detail:
{{- range $hostname, $value := .TaskDetailMap }}

  hostname: {{ $hostname }}
  {{ with .Doing }}
  doing:
      ------------------------------------------------------------------
      sectorType    sectorID    startTime                   DoingTime
      ------------------------------------------------------------------
      {{- if .AP }}
      {{- range .AP }}
      AP{{ print "            "}}{{ .SectorID | printf "%-12d" }}{{ .StartTime | printf "%-28s" }}{{ .DoingTime }}
      {{- end }}
      ------------------------------------------------------------------
      {{- end }}
      {{- if .P1 }}
      {{- range .P1 }}
      P1{{ print "            "}}{{ .SectorID | printf "%-12d" }}{{ .StartTime | printf "%-28s" }}{{ .DoingTime }}
      {{- end }}
      ------------------------------------------------------------------
      {{- end }}
      {{- if .P2 }}
      {{- range .P2 }}
      P2{{ print "            "}}{{ .SectorID | printf "%-12d" }}{{ .StartTime | printf "%-28s" }}{{ .DoingTime }}
      {{- end }}
      ------------------------------------------------------------------
      {{- end }}
      {{- if .C1 }}
      {{- range .C1 }}
      C1{{ print "            "}}{{ .SectorID | printf "%-12d" }}{{ .StartTime | printf "%-28s" }}{{ .DoingTime }}
      {{- end }}
      ------------------------------------------------------------------
      {{- end }}
      {{- if .C2 }}
      {{- range .C2 }}
      C2{{ print "            "}}{{ .SectorID | printf "%-12d" }}{{ .StartTime | printf "%-28s" }}{{ .DoingTime }}
      {{- end }}
      ------------------------------------------------------------------
      {{- end }}
  {{- end }}
  {{ with .Waiting }}
  waiting task list:
    {{- if .AP }}
    AP: {{ .AP }}
    {{- end }}
    {{- if .P1 }}
    P1: {{ .P1 }}
    {{- end }}
    {{- if .P2 }}
    P2: {{ .P2 }}
    {{- end }}
    {{- if .C1 }}
    C1: {{ .C1 }}
    {{- end }}
  {{- end }}
{{- end }}
--------------------------------------------------------------------------------------------------------------------------------
{{- end }}

{{ if .C2Waiting }}
--------------------------------------------------------------------------------------------------------------------------------
C2 Waiting: {{ .C2Waiting }}
--------------------------------------------------------------------------------------------------------------------------------
{{- end }}
`


type WorkerName string

type doingTaskInfo struct {
	SectorID  int    `json:"sector_id"`
	StartTime string `json:"start_time"`
	DoingTime string `json:"doing_time"`
}

type TaskGlanceInfo struct {
	Config config `json:"config"`
	// 强制为 ap、p1、p2、c1、c2
	DoingTotal [5]int `json:"doing_total"`
	WaitingTotal [4]int `json:"waiting_total"`
}

type taskDetail struct {
	Doing struct{
		AP []doingTaskInfo `json:"AP"`
		P1 []doingTaskInfo `json:"P1"`
		P2 []doingTaskInfo `json:"P2"`
		C1 []doingTaskInfo `json:"C1"`
		C2 []doingTaskInfo `json:"C2"`
	} `json:"doing"`
	Waiting struct{
		AP []int `json:"AP"`
		P1 []int `json:"P1"`
		P2 []int `json:"P2"`
		C1 []int `json:"C1"`
	} `json:"waiting"`
}

type renderTaskInfo struct {
	TaskDetailMap map[WorkerName]taskDetail  `json:"task_details"`
	C2Waiting []int `json:"c2_waiting"`
	TotalTaskGlanceInfo []string `json:"total_info"`
}

type config struct {
	APMaxTaskCount int
	P1MaxTaskCount int
	P2MaxTaskCount int
	C1MaxTaskCount int
	C2MaxTaskCount int
}

func renderTmpl(w io.Writer, detail renderTaskInfo) error {
	renderTemp := ""
	if fileName := os.Getenv("WORKER_TASK_DETAIL_TMPL"); fileName != "" {
		f, err := os.Open(fileName)
		if err != nil {
			return xerrors.Errorf("open tempalte file:%s error:%s", fileName, err)
		}
		defer f.Close()
		d, err := ioutil.ReadAll(f)
		if err != nil {
			return xerrors.Errorf("read data from file:%s error:%s", fileName, err)
		}
		renderTemp = string(d)
	} else {
		renderTemp = renderTaskTemplate
	}

	t, err := template.New("worker task detail").Parse(renderTemp)
	if err != nil {
		return xerrors.Errorf("parse template error:%s", err)
	}
	err = t.Execute(w, detail)
	if err != nil {
		return xerrors.Errorf("template execute error:%s", err)
	}
	return nil
}

func main() {
	path := "./worker_task_detail.json"
	f, err := os.Open(path)
	if err != nil {
		fmt.Printf("open file:%s error:%s", path, err)
		return
	}
	data, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Printf("read file:%s error:%s", path, err)
		return
	}
	var td renderTaskInfo
	if err := json.Unmarshal(data, &td); err != nil {
		fmt.Printf("json unmarshal error:%s", err)
		return
	}

	t := make(map[WorkerName]TaskGlanceInfo)
	c := config{APMaxTaskCount: 10, P1MaxTaskCount: 10, P2MaxTaskCount: 10, C1MaxTaskCount: 10, C2MaxTaskCount: 10}

	t["host1"] = TaskGlanceInfo{
		Config:       c,
		DoingTotal:   [5]int{0,2,3,4,5},
		WaitingTotal: [4]int{6,7,8,9},
	}
	t["host2"] = TaskGlanceInfo{
		Config:       c,
		DoingTotal:   [5]int{6,7,8,9,11},
		WaitingTotal: [4]int{6,7,8,9},
	}
	t["host3"] = TaskGlanceInfo{
		Config:       config{APMaxTaskCount: 0, P1MaxTaskCount: 0, P2MaxTaskCount: 0, C1MaxTaskCount: 0, C2MaxTaskCount: 3},
		DoingTotal:   [5]int{0,0,0,0,1},
		WaitingTotal: [4]int{0,0,0,0},
	}

	var result bytes.Buffer
	td.TotalTaskGlanceInfo = getTaskGlanceInfoRender(t)
	err = renderTmpl(&result, td)
	if err != nil {
		fmt.Printf("render template error:%s", err)
		return
	}
	fmt.Print(result.String())

	//renderManual(os.Stdout, td)
}

func getTaskGlanceInfoRender(taskInfo map[WorkerName]TaskGlanceInfo) []string {
	res := make([]string, len(taskInfo))
	baseFormat := "\n%s\n  doing:\n    ap: %s\n    p1: %s\n    p2: %s\n    c1: %s\n    c2: %s\n  waiting task count:\n    ap: %d\n    p1: %d\n    p2: %d\n    c1: %d\n"
	c2Format := "\n%s(maybe a c2 worker?)\n  doing:\n    c2: %s\n"
	writeProcess := func (count int, max int) string {
		if count == 0 && max == 0 {
			return ""
		}
		if max <= 0 {
			return fmt.Sprintf("最大运行数量为 %d，小于等于零，不予展示（实际运行数量为 %d）",count, max)
		}
		if count > max {
			return fmt.Sprintf("最大运行数量为 %d，实际运行数量 为 %d，大于最大运行数量，无法展示",count, max)
		}
		var b bytes.Buffer
		b.WriteRune('[')
		for c := 0; c < 100; c++ {
			if float64(c) < float64(count)/float64(max) * 100 {
				b.WriteRune('>')
			} else {
				b.WriteRune(' ')
			}
		}
		b.WriteRune(']')
		b.WriteString(fmt.Sprintf("(实际任务：%d/设计任务：%d)", count, max))
		return b.String()
	}

	hostnameList := make([]WorkerName, 0)
	for hostname, _ := range taskInfo {
		hostnameList = append(hostnameList, hostname)
	}
	sort.Slice(hostnameList, func(i int, j int) bool {
		return hostnameList[i] < hostnameList[j]
	})
	for _, hostname := range hostnameList {
		info := taskInfo[hostname]
		dAP := writeProcess(info.DoingTotal[0], info.Config.APMaxTaskCount)
		dP1 := writeProcess(info.DoingTotal[1], info.Config.P1MaxTaskCount)
		dP2 := writeProcess(info.DoingTotal[2], info.Config.P2MaxTaskCount)
		dC1 := writeProcess(info.DoingTotal[3], info.Config.C1MaxTaskCount)
		dC2 := writeProcess(info.DoingTotal[4], info.Config.C2MaxTaskCount)
		// 对于c2worker的展示内容会少很多
		if dAP == "" && dP1 == "" && dP2 == "" && dC1 == "" && dC2 != "" {
			res = append(res, fmt.Sprintf(c2Format, hostname, dC2))
		} else {
			res = append(res, fmt.Sprintf(baseFormat,
				hostname, dAP, dP1, dP2, dC1, dC2,
				info.WaitingTotal[0], info.WaitingTotal[1],info.WaitingTotal[2], info.WaitingTotal[3]))
		}

	}
	return res
}
