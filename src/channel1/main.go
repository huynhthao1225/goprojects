package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	ch := make(chan int, 50)
	wg.Add(2)
	go receiver(ch)
	go sender(ch, 110)
	wg.Wait()

}

func receiver(ch <-chan int) {

	for v := range ch {
		fmt.Println("have ", v)

	}
	wg.Done()
}
func sender(ch chan<- int, k int) {
	for i := 1; i < k; i++ {
		ch <- i
	}
	close(ch)
	wg.Done()

}
