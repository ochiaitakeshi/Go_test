package main

import (
	"fmt"
	"time"
)

func main() {
	t1 := time.NewTicker(1 * time.Second) // 1秒おきに通知
	t2 := time.NewTicker(3 * time.Second) // 3秒おきに通知

	for {
		select {
		case <-t1.C:
			// 1秒経過した。ここで何かを行う。
			fmt.Printf("hello, world 1s\n")
		case <-t2.C:
			// 3秒経過した。ここで何かを行う。
			fmt.Printf("hello, world 3s\n")
		}
	}
	t1.Stop() // タイマを止める。
	t2.Stop() // タイマを止める。
}
