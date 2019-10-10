package main

import (
	"fmt"
)

func main() {

	intc := make(chan int)
	go todo(intc)
	for {
		value, ok := <-intc
		if !ok {
			fmt.Println("get ", value, ok)
			break
		}
		fmt.Println("get ", value, ok)
	}

	intc1 := make(chan int)
	go todo(intc1)

	for v := range intc1 {
		fmt.Println("have ", v)
	}

}

func todo(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}
