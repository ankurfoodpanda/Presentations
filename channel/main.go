package main

import (
	"fmt"
	"sync"
)

var ch chan interface{}

func sender(ch chan<- interface{}, wg * sync.WaitGroup){
	defer wg.Done()
	ch <- "hello"
}


func receiver(ch <-chan interface{}, wg * sync.WaitGroup){
	defer wg.Done()
	fmt.Println(<-ch)
}

func main(){
	// bidirectional channel
	var wg sync.WaitGroup
	ch = make(chan interface{})
	for range []int{1, 2, 3, 4}{
		wg.Add(2)
		go sender(ch, &wg)
		go receiver(ch, &wg)
	}
	wg.Wait()
}