// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	Timer(mathBeMathin)
// }

// func mathBeMathin() {
// 	for i := 0; i < 150_000; i++ {
// 		x := (i*i + 1) / (i + 2)
// 		fmt.Println(x)
// 	}
// }

//	func Timer(f func()) {
//		start := time.Now()
//		f()
//		total := time.Since(start)
//		fmt.Println(total)
//	}
package main

import (
	"fmt"
	"time"
)

func main() {
	Timer(mathBeMathin)
}

func mathBeMathin() {
	chunkSize := 1000
	numChunks := 150_000 / chunkSize
	done := make(chan bool)

	for i := 0; i < numChunks; i++ {
		go func(start int) {
			for j := start; j < start+chunkSize; j++ {
				_ = (j*j + 1) / (j + 2)
			}
			done <- true
		}(i * chunkSize)
	}

	for i := 0; i < numChunks; i++ {
		<-done
	}
}

func Timer(f func()) {
	start := time.Now()
	f()
	total := time.Since(start)
	fmt.Println(total)
}
