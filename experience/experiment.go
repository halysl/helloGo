package experience

import (
	"encoding/json"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr/net"
	"github.com/shirou/gopsutil/disk"
	"io/ioutil"
	"math/rand"
	"net/url"
	"os"
	"sort"
	"strings"
	"fmt"
	"time"
)

func GenerateAddress(s string) (string, error) {
	// s := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.dhLZxNSPEZYLfQIbPAJHzrtXchRzPmY1yrsiVKU:/ip4/127.0.0.1/tcp/1234/http"
	sp := strings.SplitN(s, ":", 2)
	tok := []byte(sp[0])
	s = sp[1]

	fmt.Println(tok)
	fmt.Println(s)
	ma, err := multiaddr.NewMultiaddr(s)
	if err == nil {
		_, addr, err := manet.DialArgs(ma)
		if err != nil {
			return "", err
		}

		return "ws://" + addr + "/rpc/" + "v1", nil
	}

	_, err = url.Parse(s)
	if err != nil {
		return "", err
	}
	return s + "/rpc/" + "v1", nil
}

func LengthOfLongestSubString(s string) int {
	max := func(i, j int) int {
		if i > j {
			return i
		}
		return j
	}
	length := len(s)
	ans := 0
	cmap := make([]int, 128)
	left := 0
	for right := 0; right < length; right++ {
		index := s[right]
		left = max(left, cmap[index])
		ans = max(ans, right-left+1)
		cmap[index] = right + 1
	}
	return ans
}

type A struct {
	BList []B
}
type B struct {
	Score float64
}

func (a *A) score() {
	timeTicker := time.NewTicker(time.Second)
	for {
		select {
		case <-timeTicker.C:
			for index, _ := range a.BList {
				a.BList[index].Score = rand.Float64()
			}
			sort.Slice(a.BList, func(i, j int) bool { return a.BList[i].Score > a.BList[j].Score })
		}
	}
}

func SortStructSlice() {
	a := A{BList: []B{B{Score: 0.0}, B{Score: 0.0}, B{Score: 0.0}}}
	go a.score()
	i := 0
	for _, v := range a.BList {
		time.Sleep(2 * time.Second)
		fmt.Println(v.Score)
		i++
		if i > 10 {
			break
		}
	}
}

func SortDurationSlice() {
	start := time.Now()
	tdSlice := make([]time.Duration, 10)
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		tdSlice[rand.Intn(10)] = time.Now().Sub(start)
	}
	sort.Slice(tdSlice, func(i, j int) bool {
		return tdSlice[i] > tdSlice[j]
	})
	fmt.Println(tdSlice)
}

func IsPalindrome(x int) bool {
	origin := x
	base := 0
	for x != 0 {
		last := x % 10
		base = base*10 + last
		x = x / 10
	}
	return origin == base
}


func DelMapKey() {
	m := make(map[string]int)
	m["1"] = 1
	fmt.Println(m)
	fmt.Println(m["1"])
	fmt.Println(m["2"])
	delete(m, "1")
	fmt.Println(m)
	fmt.Println(m["1"])
	delete(m, "2")
	fmt.Println(m)
	fmt.Println(m["2"])
}

func PrintdiskUsage(p string) interface{} {
	du, _ := disk.Usage(p)
	return du
}

func JsonUnmarshalNull() {
	type SectorPath struct {
		Miner    uint64 `json:"miner"`
		Number   uint64 `json:"number"`
		Unsealed string `json:"unsealed"`
		Sealed   string `json:"sealed"`
		Cache    string `json:"cache"`
	}

	type CopyDetail struct {
		Success       bool        `json:"success"`
		Sector        *SectorPath `json:"sector"`
		FinishTime    time.Time   `json:"finish_time"`
		FailedMessage string      `json:"failed_message"`
	}

	res := make(map[uint64]CopyDetail)
	f, _ := os.Open("/tmp/a.json")
	defer f.Close()
	d, _ := ioutil.ReadAll(f)
	json.Unmarshal(d, &res)
	fmt.Printf("%+v", res)
}
