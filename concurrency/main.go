package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// async1()
	// async2()
	// async3()
	async4()
}

func async1() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		fmt.Println("this happens async")
		wg.Done()
	}()

	fmt.Println("THis is synchronous")
	wg.Wait()
}

func async2() {
	var wg sync.WaitGroup
	ch := make(chan string)

	wg.Add(1)

	go func() {
		ch <- "the message"

	}()

	go func() {
		fmt.Println(<-ch)
		wg.Done()
	}()

	wg.Wait()
}

func async3() {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		ch1 <- "message to ch1"
	}()

	go func() {
		ch2 <- "message to ch2"
	}()

	time.Sleep(40 * time.Microsecond)

	select {
	case msg := <-ch1:
		fmt.Println(msg)
	case msg := <-ch2:
		fmt.Println(msg)
	default:
		fmt.Println(("no msg available"))
	}

}

func async4() {
	ch := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
	}()

	for msg := range ch {
		fmt.Println(msg)

	}
}
