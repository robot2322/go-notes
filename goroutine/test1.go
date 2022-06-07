package goroutine

import (
	"fmt"
)

func Test1() {

	c := make(chan bool, 100)

	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
			c <- true
		}(i)
	}

	for i := 0; i < 100; i++ {
		<-c
	}
}
