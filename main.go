package main

import (
	"fmt"

	"github.com/halysl/hellogo/basis"
	"github.com/halysl/hellogo/experience"
	"github.com/halysl/hellogo/tlotus"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func initViper() {
	configPath := "conf/config.toml"
	viper.SetConfigType("toml")
	viper.SetConfigFile(configPath)
	if err := viper.ReadInConfig(); err != nil {
		log.Warnf("read config err %+v,try to reset config file with default value", err)
		panic(err)
	}
}

func main() {
	res, _ := experience.GenerateAddress("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJBbGxvdyI6WyJyZWFkIiwid3JpdGUiLCJzaWduIiwiYWRtaW4iXX0.dhLZxNSPEZYLfQIbPAJHzrtXchRzPmY1yrsiVKU:/ip4/127.0.0.1/tcp/1234/http")
	show(res)

	basis.UseContext()

	tlotus.TestTicket()
	tlotus.TryDeclare()

	show(experience.LengthOfLongestSubString(""))
	show(experience.LengthOfLongestSubString("abcabcbb"))
	show(experience.LengthOfLongestSubString("bbbbbbb"))
	show(experience.LengthOfLongestSubString("pwwkew"))

	experience.SortStructSlice()

	experience.SortDurationSlice()

	experience.DelMapKey()

	show(experience.PrintdiskUsage("/Users/light/media"))

	experience.JsonUnmarshalNull()
}

func show(s interface{}) {
	fmt.Println(s)
}
