package getinfo

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"getMonitorInfo/static"
	"net/http"
)

func httpGet(url string) (httpGetData []byte, statusCode int) {
	// get 请求从 url 拿数据，返回 []byte, status code
	log.Printf("url:%s\n", url)
	resp, err := http.Get(url)
	if err != nil {
		statusCode = -100
		log.Fatal(err)
		return
	}
	defer resp.Body.Close()

	httpGetData, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	return
}

func GetTargetInfo() (targetInfo []byte, err error) {
	targetInfo, statusCode := httpGet(static.BaseTargetUrl)
	if statusCode != 200 {
		log.Printf("url: %s\tstatus code: %d\n", static.BaseTargetUrl, statusCode)
		err = errors.New("No Data")
	}
	return
}

func GetQueryInfo(query, t string) (queryInfo []byte, err error) {
	queryInfo, statusCode := httpGet(fmt.Sprintf(static.BaseQueryUrl, query, t))
	if statusCode != 200 {
		log.Printf("url: %s\tstatus code: %d\n", static.BaseQueryUrl, statusCode)
		err = errors.New("No Data")
	}
	return
}
