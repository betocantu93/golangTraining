package main

import (
	"sync"
	"fmt"
	"time"
	"math/rand"
)

var wg sync.WaitGroup
var counter int
var mutex sync.Mutex



func main(){
	wg.Add(2)
	go incrementor("Foo:")
	go incrementor("Bar:")
	wg.Wait()
	fmt.Println("Final Counter: ",counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		time.Sleep( time.Duration(rand.Intn(3)) * time.Microsecond)
		mutex.Lock()
		counter++
		fmt.Println(s, i, "counter: ", counter)
		mutex.Unlock()
	}
	wg.Done()
}
