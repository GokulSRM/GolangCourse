package main

import "fmt"

func double(a int) { //function
	fmt.Println(a * 2)
}

type T struct { //struct
	val int
}

func (p *T) a() { //pointer function
	p.val += 1
}

func (p T) b() {
	p.val += 2
}

func sum(nums ...int){  // variadic function
	total := 0
	for _, v:= range nums{
		total+=v
	}
	fmt.Println(total)
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

	a := [5]int{0, 2, 4, 6, 8} //arrays
	var s []int = a[1:3]       //slices
	s[0] = 8
	fmt.Println(a)
	fmt.Println(s[0]) 
	
	/*A slice does not store any data;
	it just describes a section of an underlying array.
	Changing the elements of a slice modifies the corresponding
	elements of its underlying array*/

	a1 := make([]int, 5) // make methode for slice
	a1 = append(a1, 8)   //append method to add values to slice
	fmt.Println(a1)

	/*
	   	m := make(map[string]int) //map using make()method { m := make(map[string]int) }
	       m["James"] = 42
	       m["Amy"] = 24

	       fmt.Println(m["James"])
	*/
	m := map[string]int{
		"James": 42,
		"Amy":   24}

	fmt.Println(m["Amy"])

	delete(m, "James") //delete in map

	m["Bob"] = 8

	fmt.Println(m)

	sum (2,4,6) // call variadic function
	sum (42,8)
	sum(1,2,3,4,5,6,7,8,9)

}
