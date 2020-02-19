package basis

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func runBasis() {
	fmt.Printf("Hello, world or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n")
	// fmt.Printf(stringutil.Reverse("\nHellow, world."))
	a := "1213133122"
	b, c := strconv.ParseFloat(a, 64)
	if c != nil {
		b = -1.0
	}
	fmt.Printf("type(a)=%T,a=%s\n", a, a)
	fmt.Printf("type(a)=%T,a=%f\n", b, b)
	fmt.Printf("type(a)=%T,a=%s\n", c, c)
	fmt.Printf("type(a)=%T,a=%s\n", a, a)
	fmt.Println("i will sleep")
	time.Sleep(10 * time.Microsecond)
	fmt.Println("hello")

	fmt.Println(useComplex())
	fmt.Println(modifiedStr("test"))
	fmt.Println(modifiedStr("wqers"))
	createErr()
	useSlice()
	useMap()
	useIf()
	useFor1()
	useBreakOrContinue()
	useBreakMoreDepth()
	useForEnumDict()
	useSwitch()
	fmt.Printf("\n a = %d b = %d max is %d", 1, 2, useFuncGetMax(1, 2))
	fmt.Printf("\n a = %d b = %d max is %d", 93, 31, useFuncGetMax(93, 31))

	useFuncMoreArg(1, 2, 3)

	testA := 1
	fmt.Printf("\n test_a is: %d, now add 1 is: %d", testA, useFuncPointArg(&testA))

	fmt.Println()
	abb := []int{1, 2, 3}
	sliceAbb := abb[1:]
	useFuncPointArgs2(sliceAbb)

	useDefer()

	go say("hello")
	say("world")
}

func useComplex() complex64 {
	return 1 + 1i
}

func modifiedStr(s string) string {
	c := []byte(s)
	c[0] = 'z'
	s = string(c)
	return s
}

func createErr() {
	err := errors.New("this is a new error")
	if err != nil {
		fmt.Println(err)
	}
}

func useSlice() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice1 := arr[1:3]
	fmt.Println()
	fmt.Printf("arr value:%d\n", arr)
	fmt.Printf("slice value:%d\nslice length:%d\nslice cap:%d", slice1, len(slice1), cap(slice1))
	fmt.Println("now i will take slice append some element")
	slice1 = append(slice1, 1)
	fmt.Printf("now arr value:%d\nnow slice value:%d", arr, slice1)
}

func useMap() {
	map1 := make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3
	fmt.Printf("\n\nthe map[three] is:%d\n", map1["three"])
	res, ok := map1["four"]
	if ok {
		fmt.Printf("map1[four] is %d\n", res)
	} else {
		fmt.Printf("map1[four] is not found\n")
	}
}

func useRandom() int {
	res := rand.Int()
	return res
}

func useIf() {
	if a := useRandom(); a > 10 {
		fmt.Printf("a is %d, more than 10", a)
	} else if a == 10 {
		fmt.Printf("a is %d, equal 10", a)
	} else {
		fmt.Printf("a is %d, less than 10", a)
	}
}

func useFor1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\nnow number is %d", i)
	}
}

func useBreakOrContinue() {
	fmt.Println()
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Printf("\nnow number is %d", i)
	}

	fmt.Println()
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		fmt.Printf("\n now number is %d", i)
	}
}

func useBreakMoreDepth() {
	fmt.Println()
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i*j == 3 {
				break
			}
			fmt.Printf("\n %d * %d = %d", i, j, i*j)
		}
	}

	fmt.Println()
test:
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if i*j == 3 {
				break test
			}
			fmt.Printf("\n %d * %d = %d", i, j, i*j)
		}
	}
}

func useForEnumDict() {
	aMap := make(map[string]int)
	aMap["first"] = 1
	aMap["second"] = 2
	aMap["third"] = 3
	for k, v := range aMap {
		fmt.Printf("\na_map['%s'] = %d", k, v)
	}
	list := [...]int{1, 2, 3, 4, 5}
	slice := list[1:]
	for index, value := range slice {
		fmt.Printf("\nindex:%d\tvalue:%d", index, value)
	}
}

func useSwitch() {
	fmt.Println()
	i := 3
	switch i {
	case 1:
		fmt.Println("i < 2")
	case 2:
		fmt.Println("i < 3")
	case 3:
		fmt.Println("i < 4")
	case 4:
		fmt.Println("i < 5")
	default:
		fmt.Println("i < 999999")
	}

	fmt.Println()
	i = 3
	switch i {
	case 1:
		fmt.Println("i < 2")
	case 2:
		fmt.Println("i < 3")
	case 3:
		fmt.Println("i < 4")
		fallthrough
	case 4:
		fmt.Println("i < 5")
		fallthrough
	default:
		fmt.Println("i < 999999")
	}
}

func useFuncGetMax(a, b int) (max int) {
	max = a
	if a < b {
		max = b
	}
	return
}

func useFuncMoreArg(arg ...int) {
	for i, v := range arg {
		fmt.Printf("\n now arg[%d]: %d", i, v)
	}
}

func useFuncPointArg(a *int) int {
	*a = *a + 1
	return *a
}

func useFuncPointArgs2(a []int) {
	fmt.Println(a[0])
	a[0] = -a[0]
	fmt.Println(a[0])
}

func useDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}

func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}
