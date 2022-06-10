// channel控制协程

package goroutine

import (
	"fmt"
	"sync"
	"time"
)

var i int

func work() {
	time.Sleep(1000 * time.Millisecond)
	i++
	fmt.Println(i)
}

func routine(command <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	var status = "开始"
	for {
		select {
		case cmd := <-command:
			fmt.Println(cmd)
			switch cmd {
			case "Stop":
				status = "停止"
				return
			case "Pause":
				status = "暂停"
			case "Start":
				status = "开始"
			}
		default:
			if status == "开始" {
				work()
			}

		}
	}
}

func Test4() {
	var wg sync.WaitGroup
	wg.Add(1)
	command := make(chan string)
	go routine(command, &wg)

	time.Sleep(5 * time.Second)
	command <- "Pause"

	time.Sleep(5 * time.Second)
	command <- "Start"

	time.Sleep(5 * time.Second)
	command <- "Stop"

	wg.Wait()
}
