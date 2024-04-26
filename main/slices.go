package main

import "fmt"

func slices() {

	s := make([]string, 3)
	fmt.Println("emp: ", s)

	s[0], s[1], s[2] = "a", "b", "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])
	fmt.Println("len: ", len(s))

	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy: ", c)

	l := s[:5]
	println("sl2: ", l)

}

func main() { slices() }
