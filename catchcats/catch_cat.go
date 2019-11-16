package catchcats

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
	"time"
)

type cat struct {
	// 一个结构体代表一个爬取项目
	name, url string
}

var base_url = "https://http.cat"

func Run() {
	data, statusCode := httpGet(base_url)
	if statusCode == 200 {
		res := regexParse(string(data))
		pic_path := getPicPath()
		for _, c := range res {
			data, statusCode = httpGet(c.url)
			if statusCode == 200 {
				write_pic(pic_path, c.name, data)
			} else {
				fmt.Printf("url: %s\tstatus code: %d\n", c.url, statusCode)
			}
		}
	} else {
		fmt.Printf("url: %s\tstatus code: %d\n", base_url, statusCode)
	}
}

func httpGet(url string) (data []byte, statusCode int) {
	// get 请求从 url 拿数据，返回 []byte, status code
	fmt.Printf("date:%s\turl:%s\n", time.Now(), url)
	resp, err := http.Get(url)
	if err != nil {
		statusCode = -100
		return data, statusCode
	}
	defer resp.Body.Close()
	data, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		statusCode = -200
		return data, statusCode
	}
	statusCode = resp.StatusCode
	return data, statusCode
}

func regexParse(data string) []cat {
	// 从爬虫结果中通过正则提取所有目标数据的 name 和 url
	imageExp := regexp.MustCompile(`\((/images/\d{3}.jpg)\)"></div><div class="Thumbnail_content__2eR9q"><div class="Thumbnail_title__2iqYK">(\d{3})</div><p>([\w-´ ]*)`)
	res := make([]cat, 0)
	tds := imageExp.FindAllStringSubmatch(data, 100)
	for _, value := range tds {
		var buffer1 bytes.Buffer
		var buffer2 bytes.Buffer
		buffer1.WriteString(value[2])
		buffer1.WriteString("_")
		buffer1.WriteString(value[3])
		name := buffer1.String()
		buffer2.WriteString(base_url)
		buffer2.WriteString(value[1])
		url := buffer2.String()
		res = append(res, cat{name, url})
	}
	return res
}

func getPicPath() string {
	// 获得图片保存目录
	pwd,  _ := os.Getwd()
	pic_dir := path.Join(pwd, "catchcats", "pic")
	return pic_dir
}

func write_pic(base_path, name string, data []byte) {
	// 写入文件，形成图片
	file, err := os.Create(base_path + "/" + name + ".jpg")
	if err != nil{
		log.Fatalln(err)
	}
	defer file.Close()
	_, err = file.Write(data)
	if err != nil {
		log.Fatalln(err)
	}
}
