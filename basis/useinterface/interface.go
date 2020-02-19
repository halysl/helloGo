package useinterface

import (
	"fmt"
	"reflect"
	"strconv"
)

type stringer interface {
	String() string
}

type uinterface interface {
	Print() string
}

type binary uint64

func (i binary) String() string {
	return strconv.FormatUint(i.get(), 10)
}

func (i binary) get() uint64 {
	return uint64(i)
}

func runInterface() {
	b := binary(200)
	s := stringer(b)
	fmt.Println(s.String())
	fmt.Println("s 的类型：", reflect.TypeOf(s))
}
