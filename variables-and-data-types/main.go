// GO's simple data types: strings, numbers, booleans and errors
package main

import "fmt"

func main() {

	// variables()
	// arithmetic()
	//constants()
	pointers()

}



func variables() {
	var a string
	a = "foo"
	fmt.Println(a)

	var b int = 99
	fmt.Println(b)

	var c = true
	fmt.Println(c)

	d := 3.14
	fmt.Println(d)

	// type conversion
	var e int = int(d)
	fmt.Println(e)

	
}

func arithmetic() {

	a, b := 5, 2
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b) // 2, not 2.5
	fmt.Println(a % b)
	fmt.Println(float32(a) / float32(b))
	fmt.Println(a == b)
	fmt.Println(a < b)
	fmt.Println(a > b)
}


func constants() {
	const a = 42
	var i int = a

	const b float32 = 3
	// var f64 float62 = b  invalid
	var f32 float32 = b

	fmt.Println(i, f32)


	const c = iota
	fmt.Println(c) // 0

	const (
		d = 2*5
		e
		f = iota // 2 since its declared as third const
		g // 3
	)
	fmt.Println(d)
	fmt.Println(e) // 10
	fmt.Println(f)
}


func pointers() {
	s := "hello"

	p := &s
	fmt.Println((p)) // memory address
	fmt.Println(*p) // value 

	*p = "hello Go" 
	fmt.Println(s) // Hello Go
	fmt.Println(*p) // Hello Go
	
	
	p = new(string)
	fmt.Println(p, *p) // p is no longer pointing to a.

}