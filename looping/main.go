package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)



func main() {
	i := 99

	for {
		fmt.Println(i)
		i += 1
		break // exit loop
	}

	i = 5
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// loop through collections
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

	for _, item := range menu {
		fmt.Println(item.name)
		fmt.Println(strings.Repeat("-", 10))
		for size, price := range item.prices {
			fmt.Printf("\t%10s%10.2f\n", size, price)
		}
	}
}