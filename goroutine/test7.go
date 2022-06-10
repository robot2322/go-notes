// 协程监听channel数据进行转账

package goroutine

import (
	"fmt"
	"net/http"
	"time"
)

type User struct {
	Cash int
}

type Transfer struct {
	Sender    *User
	Recipient *User
	Amount    int
}

func sendCashHandler(transferchan chan Transfer) {
	var val Transfer
	for {
		val = <-transferchan
		val.Sender.sendCash(val.Recipient, val.Amount)
	}
}

func (u *User) sendCash(to *User, amount int) bool {
	if u.Cash < amount {
		return false
	}
	/* 设置延迟Sleep，当多个goroutines并行执行时,便于进行数据安全分析 */
	time.Sleep(500 * time.Millisecond)
	u.Cash = u.Cash - amount
	to.Cash = to.Cash + amount
	return true
}

func Test7() {
	me := User{Cash: 500}
	you := User{Cash: 500}
	transferchan := make(chan Transfer)
	go sendCashHandler(transferchan)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		transfer := Transfer{Sender: &me, Recipient: &you, Amount: 50}
		/*转账*/
		result := make(chan int)
		go func(transferchan chan<- Transfer, transfer Transfer, result chan<- int) {
			transferchan <- transfer
			result <- 1
		}(transferchan, transfer, result)

		/*用select来处理超时机制*/
		select {
		case <-result:
			fmt.Fprintf(w, "I have $%d\n", me.Cash)
			fmt.Fprintf(w, "You have $%d\n", you.Cash)
			fmt.Fprintf(w, "Total transferred: $%d\n", (you.Cash - 500))
		case <-time.After(time.Second * 10): //超时处理
			fmt.Fprintf(w, "Your request has been received, but is processing slowly")
		}
	})
	http.ListenAndServe(":8080", nil)
}
