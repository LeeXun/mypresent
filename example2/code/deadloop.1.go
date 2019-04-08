// package deadloop

// import (
// 	"fmt"
// 	"time"
// )

// func f() {
// 	var a [1 << 18]byte
// 	a[1<<18-1] = 1
// 	// print(a[1<<18-1])
// }

// func deadloop() {
// 	for {
// 		f()
// 	}
// }

// func main() {
// 	// runtime.GOMAXPROCS(1)
// 	go deadloop()
// 	go deadloop()
// 	go deadloop()
// 	go deadloop()
// 	// go deadloop()
// 	for {
// 		// time.Sleep(time.Second * 1)
// 		fmt.Println("ASD")
// 		time.Sleep(time.Second * 1)
// 	}
// }
