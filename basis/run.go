package basis

import (
	"errors"
	"fmt"
	"math/rand"
	"runtime"
	"strconv"
	"time"
)

func Run_basis() {
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
	time.Sleep(10*time.Microsecond)
	fmt.Println("hello")

	fmt.Println(use_complex())
	fmt.Println(modified_str("test"))
	fmt.Println(modified_str("wqers"))
	create_err()
	use_slice()
	use_map()
	use_if()
	use_for_1()
	use_break_or_continue()
	use_break_more_depth()
	use_for_enum_dict()
	use_switch()
	fmt.Printf("\n a = %d b = %d max is %d", 1, 2, use_func_get_max(1, 2))
	fmt.Printf("\n a = %d b = %d max is %d", 93, 31, use_func_get_max(93, 31))

	use_func_more_arg(1, 2, 3)

	test_a := 1
	fmt.Printf("\n test_a is: %d, now add 1 is: %d", test_a, use_func_point_arg(&test_a))

	fmt.Println()
	abb := []int{1, 2, 3}
	slice_abb := abb[1:]
	use_func_point_args2(slice_abb)

	use_defer()

	go say("hello")
	say("world")
}

func use_complex() complex64 {
	return 1 + 1i
}

func modified_str(s string) string {
	c := []byte(s)
	c[0] = 'z'
	s = string(c)
	return s
}

func create_err() {
	err := errors.New("this is a new error")
	if err != nil {
		fmt.Println(err)
	}
}

func use_slice() {
	arr := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	slice1 := arr[1:3]
	fmt.Println()
	fmt.Printf("arr value:%d\n", arr)
	fmt.Printf("slice value:%d\nslice length:%d\nslice cap:%d", slice1, len(slice1), cap(slice1))
	fmt.Println("now i will take slice append some element")
	slice1 = append(slice1, 1)
	fmt.Printf("now arr value:%d\nnow slice value:%d", arr, slice1)
}

func use_map() {
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

func use_random() int {
	res := rand.Int()
	return res
}

func use_if() {
	if a := use_random(); a > 10 {
		fmt.Printf("a is %d, more than 10", a)
	} else if a == 10 {
		fmt.Printf("a is %d, equal 10")
	} else {
		fmt.Printf("a is %d, less than 10")
	}
}

func use_for_1() {
	for i := 0; i < 10; i++ {
		fmt.Printf("\nnow number is %d", i)
	}
}

func use_break_or_continue() {
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

func use_break_more_depth() {
	fmt.Println()
	for i := 0; i < 5 ; i++ {
		for j := 0; j < 5; j++ {
			if i * j == 3 {
				break
			}
			fmt.Printf("\n %d * %d = %d", i, j, i * j)
		}
	}

	fmt.Println()
test:
	for i := 0; i < 5 ; i++ {
		for j := 0; j < 5; j++ {
			if i * j == 3 {
				break test
			}
			fmt.Printf("\n %d * %d = %d", i, j, i * j)
		}
	}
}

func use_for_enum_dict() {
	a_map := make(map[string]int)
	a_map["first"] = 1
	a_map["second"] = 2
	a_map["third"] = 3
	for k, v := range a_map {
		fmt.Printf("\na_map['%s'] = %d", k, v)
	}
	list := [...]int{1, 2, 3, 4, 5}
	slice := list[1:]
	for index, value := range slice {
		fmt.Printf("\nindex:%d\tvalue:%d", index, value)
	}
}

func use_switch() {
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

func use_func_get_max(a, b int) (max int) {
	max = a
	if a < b {
		max = b
	}
	return
}

func use_func_more_arg(arg ...int) {
	for i, v := range arg {
		fmt.Printf("\n now arg[%d]: %d", i, v)
	}
}

func use_func_point_arg(a *int) int {
	*a = *a + 1
	return *a
}

func use_func_point_args2(a []int) {
	fmt.Println(a[0])
	a[0] = - a[0]
	fmt.Println(a[0])
}

func use_defer() {
	for i:=0;i<5;i++ {
		defer fmt.Println(i)
	}
}


func say(s string) {
	for i := 0; i < 5; i++ {
		runtime.Gosched()
		fmt.Println(s)
	}
}