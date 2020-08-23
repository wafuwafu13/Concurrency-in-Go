package main

import (
	"fmt"
	"time"
)

func main() {
	var data int
	go func() {
		data++
	}()
	time.Sleep(1*time.Second) // 表示させない手軽な方法ではあるが解決策ではない
	if data == 0 {
		fmt.Printf("the value is %v.\n", data)
	}
}
