package base

import "fmt"

// 闭包

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func TestClosures() {

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
