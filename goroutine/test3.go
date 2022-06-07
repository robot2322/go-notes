package goroutine

import (
	"fmt"
	"sync"
)

func Test3() {

	num := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(1)

	go func(x int, num chan int) {
		defer wg.Done()
		z := x + 100
		// 传递值到channel
		num <- z
	}(10, num)

	// 取出channel中的值
	fmt.Println(<-num)

	wg.Wait()

}
