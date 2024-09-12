package base

import "fmt"

// 结构体
type person struct {
	name string
	age  int
}

// 构造一个新的person结构体
func newPerson(name string) *person {
	p := person{name: name}
	p.age = 42
	return &p
}

func TestStructs() {
	fmt.Println(person{"Bob", 20})
	fmt.Println(person{name: "Alice", age: 30})
	fmt.Println(person{name: "Fred"})
	// &前缀生成一个结构体的指针
	fmt.Println(&person{name: "Ann", age: 40})

	s := person{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
