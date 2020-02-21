package main

import "fmt"

type calc func(int) int

func closure01(base int) (calc, calc) {
	adder := func(i int) int {
		base += i
		return base
	}
	suber := func(i int) int {
		base -= i
		return base
	}
	return adder, suber
}

func closure01Test() {
	f1, f2 := closure01(100)
	fmt.Println(f1(1), f2(2))
	fmt.Println(f1(3), f2(4))
	fmt.Println(f1(5), f2(6))
}
