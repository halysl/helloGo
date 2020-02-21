package main

import "fmt"

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
}

func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
}

func deferTest() {
	a := f1()
	b := f2()
	c := f3()
	d := f4()
	fmt.Println(a) // 5
	fmt.Println(b) // 6
	fmt.Println(c) // 5
	fmt.Println(d) // 5
}

func defer02(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}

func defer02Test() {
	x := 1
	y := 2
	defer defer02("AA", x, defer02("A", x, y))
	x = 10
	defer defer02("BB", x, defer02("B", x, y))
	y = 20
}

// A 1 2 3
// B 10 2 12
// BB 10 12 22
// AA 1 3 4
