package goroutine

import (
	"context"
	"fmt"
	"time"
)

func Test5() {
	ctx, cancel := context.WithCancel(context.Background())
	go watch(ctx, "watch1")
	go watch(ctx, "watch2")
	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(2 * time.Second)

}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "is Done")
			return
		default:
			fmt.Println(name, "is running")
			time.Sleep(2 * time.Second)
		}
	}
}
