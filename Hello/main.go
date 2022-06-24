package main

import "fmt"

func double(a int) {
	fmt.Println(a * 2)
}
type T struct {
    val int
}
func (p *T) a() {
    p.val += 1
}
func (p T) b() {
    p.val += 2
}
func main() {
	// fmt.Println("Hello world")
	for i := 4; i > 0; i-- {
		defer double(i)
	}
	x := T{5}
    x.a()
    x.b()
    fmt.Println(x.val)
}
