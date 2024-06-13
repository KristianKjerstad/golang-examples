package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)



func main() {
	// arrays
	var arr [3]string
	fmt.Println(arr)

	arr = [3]string{"coffee", "espresso", "cappuccino"}

	fmt.Println(arr)
	arr[1] = "chai tea"
	fmt.Println(arr)
	arr2 := arr

	arr2[2] = "Chai latte"
	fmt.Println(arr, arr2)

	// slices

	s := []int{1,2,3}
	fmt.Println(s[1])
	s[1] = 99
	fmt.Println(s)


	s = append(s, 5, 10, 15)
	fmt.Println(s)

	s = slices.Delete(s, 1, 3) // remove indices from slice - up to but not including the last
	fmt.Println(s)


	// maps
	var m map[string][]string 
	fmt.Println(m)

	m = map[string][]string{
		"coffee":{"Coffee", "Espresso"}, // []string is optional
		"tea":{"Hot tea", "Chai tea"},
	}
	fmt.Println(m["coffee"])
	m["other"] = []string{"hot chocolate"}
	fmt.Println(m)

	delete(m, "tea")
	fmt.Println(m)
	v, ok := m["tea"]
	fmt.Println(v, ok)

	m2 := m
	m2["coffee"] = []string{"coffee"}
	m["tea"] = []string{"hot tea"}
	fmt.Println(m)
	fmt.Println(m2) 


	// structs
	testStructs()

}


func testStructs() {
	fmt.Println("Please select an option")
	fmt.Println("1) Print menu")
	in := bufio.NewReader(os.Stdin)
	choice, _ := in.ReadString('\n')
	choice = strings.TrimSpace(choice)

	type menuItem struct{
		name string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.5, "medium": 1.85, "large": 2.2}},
		{name: "Espresso", prices: map[string]float64{"single": 2.2, "double": 2.55, "triple": 2.82}},
	}

	fmt.Println(menu)
}