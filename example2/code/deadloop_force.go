package main

import (
	"fmt"
	"time"
)

func f() {
	var a [1 << 18]byte
	a[1<<18-1] = 1
}

func deadloop() {
	for {
		f()
	}
}

func main() {
	go deadloop()
	go deadloop()
	go deadloop()
	go deadloop()
	for {
		fmt.Println("Main Loop")
		time.Sleep(time.Second * 1)
	}
}
