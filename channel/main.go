package main

import (
	"fmt"
	"sync"
)

var ch chan interface{}

// Example of compiler error
//writeStream := make(chan<- interface{})
//readStream := make(<-chan interface{})
//<-writeStream
//readStream <- struct{}{}


// Creating a buffered channel
// make(chan int, 4)
// why use buffered channel? to unblock the producer as soon as possible
// however you need to know the number of messages up front


//var c1, c2 <-chan interface{}
//var c3 chan<- interface{}
//select {
//case <- c1:
//// Do something
//case <- c2:
//// Do something
//case c3<- struct{}{}:
//// Do something
//}

func sender(ch chan<- interface{}, wg * sync.WaitGroup){
	defer wg.Done()
	ch <- "hello"
}


func receiver(ch <-chan interface{}, wg * sync.WaitGroup){
	defer wg.Done()
	str, ok := <-ch
	if !ok{
		fmt.Println("channel is closed")
		return
	}
	fmt.Println(str)
}

func main(){
	// bidirectional channel
	var wg sync.WaitGroup
	ch = make(chan interface{})
	defer close(ch)
	for range []int{1, 2, 3, 4}{
		wg.Add(2)
		go sender(ch, &wg)
		go receiver(ch, &wg)
	}
	wg.Wait()
}
