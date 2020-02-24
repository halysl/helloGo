package main

// import "fmt"

// type student struct {
// 	name string
// 	age  int
// }

// func structTest() {
// 	m := make(map[string]*student)
// 	fmt.Printf("m=%T\tid(m)=%p\n", m, m)
// 	stus := []student{
// 		{name: "小王子", age: 18},
// 		{name: "娜扎", age: 23},
// 		{name: "大王八", age: 9000},
// 	}
// 	fmt.Printf("%+v\n", stus)
// 	fmt.Printf("stus[0]:%p\nstus[1]:%p\nstus[2]:%p\n\n", &stus[0], &stus[1], &stus[2])

// 	for _, stu := range stus {
// 		fmt.Printf("id(stu):%p\tstu.name=%s\tstu.age=%d\n", &stu, stu.name, stu.age)
// 		m[stu.name] = &stu
// 	}
// 	for k, v := range m {
// 		fmt.Println(k, "=>", v.name)
// 	}
// }
