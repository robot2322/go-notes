// WaitGroup主进程阻塞

package goroutine

import (
	"fmt"
	"sync"
)

func Test2() {

	wg := sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			wg.Done()
		}(i)
	}

	wg.Wait()

}
