package spider_usda_nal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

/* 用于爬取 美国农业部 公布的 自然农作物 的水彩画
*/
type AllInfo []Urls

type Urls struct {
	page int
	urlStruct []Url
}

type Url struct {
	url string
	number int
	artist string
	year int
	scientificName string
	commonName string
	country string
	specimen string
	picUrl string
}

const base_url = "https://usdawatercolors.nal.usda.gov/pom/search.xhtml?start=0"

func Run() {
	info := AllInfo{}
	data, statusCode := httpGet(base_url)
	if statusCode == 200 {
		urls := Urls{page: getPage(data)}
		urls.urlStruct = getItemsInfo(data)
		info = append(info, urls)
		for next_url, have_next := getNextUrl(data);have_next {
			data, statusCode = httpGet(next_url)
			if statusCode == 200 {
				urls := Urls{page: getPage(data)}
				urls.urlStruct = getItemsInfo(data)
				info = append(info, urls)
			}
		}
	}
}

func httpGet(url string) (content string, statusCode int) {
	// get 请求从 url 拿数据，返回 []byte, status code
	fmt.Printf("date:%s\turl:%s\n", time.Now(), url)
	resp, err := http.Get(url)
	if err != nil {
		statusCode = -100
		return
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	content = string(data)
	if err != nil {
		statusCode = -200
		return
	}
	statusCode = resp.StatusCode
	return
}

func getNextUrl(data string) (string, bool) {}

func getItemsInfo(data string) []Url {
	return []Url
}

func getPage(data string) int {

}
