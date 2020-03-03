package main

import (
	"fmt"
	"strconv"
)

func strcovDemo() {
	var (
		a=1000
		b=true
		c=1.0
		d="0"
		e="1"
		f="3.14"
		g="0xff"

		)
	res0, _ := strconv.Atoi(d)
	fmt.Println(res0, strconv.Itoa(a))

	res1, _ := strconv.ParseInt(d, 10, 64)
	res2, _ := strconv.ParseBool(e)
	res3, _ := strconv.ParseFloat(f, 64)
	res4, _ := strconv.ParseInt(g, 16, 32)
	fmt.Println(res1, res2, res3, res4)

    res5 := strconv.FormatFloat(c, 'f', 10, 64)
    res6 := strconv.FormatBool(b)
    res7 := strconv.FormatInt(int64(a), 10)
    res8 := strconv.FormatInt(int64(a), 16)
    fmt.Println(res5, res6, res7, res8)
}
