package main

import (
	"fmt"
	"time"
)

func main() {
	a := makeTimestamp()

	fmt.Printf("%d \n", a)
}

func makeTimestamp() int64 {
	//return time.Now().UnixNano() / int64(time.Millisecond)
	return time.Now().UnixNano() / int64(time.Second)
}
