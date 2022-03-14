package playground

import (
	"fmt"
	"time"
)

func sum(x, y int, intChan chan int) {
	time.Sleep(3 * time.Second)
	intChan <- x + y
	close(intChan)
}

func GoroutineGetFromChan() {
	intChan := make(chan int, 1)
	go sum(12, 13, intChan)
	fmt.Println("got from chan:", <-intChan)
}
