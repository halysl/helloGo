package main

// 通过 gob 模块对数据序列化的同时保存数据类型

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

type s struct {
	data map[string]interface{}
}

func gobDemo() {
	var s1 = s{
		data: make(map[string]interface{}, 8),
	}
	s1.data["count"] = 1

	// encode
	buf := new(bytes.Buffer)
	enc := gob.NewEncoder(buf)
	err := enc.Encode(s1.data)
	if err != nil {
		fmt.Println("gob encode failed, err:", err)
		return
	}
	b := buf.Bytes()
	fmt.Println(b)
	var s2 = s{
		data: make(map[string]interface{}, 8),
	}

	// decode
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(&s2.data)
	if err != nil {
		fmt.Println("gob decode failed, err", err)
		return
	}
	fmt.Println(s2.data)
	for _, v := range s2.data {
		fmt.Printf("value:%v, type:%T\n", v, v)
	}
}
