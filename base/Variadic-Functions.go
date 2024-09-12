package base

import "fmt"

// 变参函数
func Sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, x := range nums {
		total += x
	}
	fmt.Println(total)
}

func TestVariadicFunctions() {
	Sum(1, 2)
	Sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	Sum(nums...)
}
