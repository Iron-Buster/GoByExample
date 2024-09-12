package base

import "fmt"

func Vals() (int, int) {
	return 3, 7
}

func TestVals() {
	a, b := Vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := Vals()
	fmt.Println(c)
}
