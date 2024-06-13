package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// branching()
	//deferredFunction()
	run()

}

func branching() {
	i := 8
	if i < 5 {
		fmt.Println("less than 5")
	} else if i < 10 {
		fmt.Println("less than 10")
	} else {
		fmt.Println("at least 10")
	}

	switch i {
	case 1:
		fmt.Println("first case")
	case 2 + 2:
		fmt.Println("second case")
	default:
		fmt.Println("default case")
	}

	// logical switch
	switch j := 8; {
	case j < 5:
		fmt.Println("less than 5")
	case j < 10:
		fmt.Println("less than 10")
	default:
		fmt.Println("greater than 10")
	}

	type menuItem struct {
		name   string
		prices map[string]float64
	}

	menu := []menuItem{
		{name: "Coffee", prices: map[string]float64{"small": 1.5, "medium": 1.85, "large": 2.2}},
		{name: "Espresso", prices: map[string]float64{"single": 2.2, "double": 2.55, "triple": 2.82}},
	}
loop: //label
	for {
		in := bufio.NewReader(os.Stdin)
		fmt.Println("Please select an option")
		fmt.Println("1) Print menu")
		fmt.Println("2) Add item")
		fmt.Println("q) quit")

		choice, _ := in.ReadString('\n')
		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			for _, item := range menu {
				fmt.Println(item.name)
				fmt.Println(strings.Repeat("-", 10))
				for size, price := range item.prices {
					fmt.Printf("\t%10s%10.2f\n", size, price)
				}
			}
		case "2":
			fmt.Println("Enter name of the new item")
			name, _ := in.ReadString('\n')
			menu = append(menu, menuItem{name: strings.TrimSpace(name), prices: make(map[string]float64)})

		case "q":
			break loop
		default:
			fmt.Println("Unknown option")
		}

	}
}

func deferredFunction() {
	fmt.Println("main1")
	defer fmt.Println("defer1")
	fmt.Println("main2")
	defer fmt.Println("defer2")

}

func run() {
	dividend, divisor := 10, 5
	fmt.Printf("%v divided by %v is %v\n", dividend, divisor, divide(dividend, divisor))

	dividend2, divisor2 := 10, 0
	fmt.Printf("%v divided by %v is %v\n", dividend2, divisor2, divide(dividend2, divisor2))
}

func divide(dividend, divisor int) int {
	defer func() {

		if msg := recover(); msg != nil {
			fmt.Println(msg)
		}
	}()
	return dividend / divisor
}
