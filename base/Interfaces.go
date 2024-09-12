package base

import (
	"fmt"
	"math"
)

/*
*
接口
实现Go中的接口，只需要实现接口中的所有方法
*/
type geometry interface {
	area() float64
	perim() float64
}

type rectX struct {
	w, h float64
}

type circle struct {
	radius float64
}

func (r rectX) area() float64 {
	return r.w * r.h
}

func (r rectX) perim() float64 {
	return 2*r.w + 2*r.h
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 通用的measure函数，将geometry 接口的实例传递过来。
func measure(g geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perim())
}

func TestInterfaces() {

	r := rectX{w: 3, h: 4}
	c := circle{radius: 5}

	measure(r)
	measure(c)
}
