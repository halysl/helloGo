package usechannel

import "fmt"

func sum(nl []int, ch chan int) {
	total := 0
	for _, v := range nl {
		total += v
	}
	ch <- total
}

func main() {
	a := []int{1, 2, 3, 4, 5, 6, 7}
	ch := make(chan int)
	lenA := len(a)
	go sum(a[:lenA/2], ch)
	go sum(a[lenA/2:], ch)
	x := <-ch
	y := <-ch
	fmt.Printf("a=%v\nlen(a)=%d\nsum(a[%d:])=%d\nsum(a[:%d])=%d\n", a, lenA, lenA/2, x, lenA/2, y)

	useCacheChan()

	//go say1("t1")
	//go say1("t2")
	//go say1("t3")
	//go say1("t4")
	//go say1("t5")
	//go say1("t6")

	c := make(chan int, 10)
	go fib(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}

	funcfunc()
}

func useCacheChan() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}

func say1(s string) {
	fmt.Println(s)
}

func fib(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func funcfunc() {
	printA := func(s string) string {
		fmt.Print(s)
		return s
	}

	s := printA("hemmo")
	fmt.Println(s)

}
