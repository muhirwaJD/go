package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func singleCook(cookId int) {
	fmt.Printf("Worked %v Started cooking\n", cookId)
	time.Sleep(time.Second)
	fmt.Printf("Worked %v Finished cooking\n", cookId)
	wg.Done()
}

func main() {
	wg.Add(3)
	for i := 1; i <= 3; i++ {
		go singleCook(i)
	}

	wg.Wait()

}
