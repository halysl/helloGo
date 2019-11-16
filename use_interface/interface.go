package use_interface

import (
	"fmt"
	"reflect"
	"strconv"
)

type Stringer interface {
	String() string
}

type Uinterface interface {
	Print() string
}

type Binary uint64
func (i Binary) String() string {
	return strconv.FormatUint(i.Get(), 10)
}
func (i Binary) Get() uint64 {
	return uint64(i)
}

func Run_interface() {
	b := Binary(200)
	s := Stringer(b)
	fmt.Println(s.String())
	fmt.Println("s 的类型：", reflect.TypeOf(s))
}
