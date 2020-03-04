package main

// msgpack 更高效的二进制序列化格式

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	"log"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func msgpackDemo() {
	p1 := Person{
		Name:   "Ash",
		Age:    18,
		Gender: "男",
	}
	// marshal
	b, err := msgpack.Marshal(p1)
	if err != nil {
		fmt.Printf("msgpack marshal failed,err:%v", err)
		return
	}

	// unmarshal
	var p2 Person
	err = msgpack.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("msgpack unmarshal failed,err:%v", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2) // p2:main.Person{Name:"沙河娜扎", Age:18, Gender:"男"}
}

func msgpackDemo2() {
	var a = [10]int{0,1,2,3,4,5,6,7,8,9}
	var b [10]int

	data1, err := msgpack.Marshal(a)
	if err != nil {
		log.Fatal(err)
		return
	}

	err = msgpack.Unmarshal(data1, &b)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Printf("b:%#v\n", b)
}
