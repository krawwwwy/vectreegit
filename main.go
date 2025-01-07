package main

import "fmt"

func zero(xPtr *int) {
	*xPtr = 0
}
func main() {
	x := 5
	zero(&x)
	fmt.Println(x) // x is 0
}

func sum(sl []int) int {
	ans := 0
	for _, v := range sl {
		ans += v
	}
	panic("asd")
	return ans

}

func half(x int) (int, bool) {
	return x / 2, x%2 == 0
}

func maxing(args ...int) int {
	ans := args[0]
	for _, v := range args {
		if v > ans {
			ans = v
		}
	}
	return ans
}

func makeOddGenerator() func() int {
	i := 1
	return func() int {
		i += 2
		return i
	}
}

func fib(n int) int {
	if n < 2 {
		return n
	}
	return fib(n-1) + fib(n-2)
}
