package base

import "fmt"

// 方法
type rect struct {
	w, h int
}

/**

方法的定义
>>>>> func (接收者) 函数名 (参数列表) (返回值列表) {}
>>>>> 调用方法时，Go 会自动处理值和指针之间的转换
*/

func (r *rect) area() int {
	return r.w * r.h
}

func (r rect) perim() int {
	return 2*r.w + 2*r.h
}

func TestMethods() {
	r := rect{w: 10, h: 5}

	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim: ", rp.perim())
}
