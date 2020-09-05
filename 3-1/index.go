package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		salutation = "welcome" // <1>
	}()
	wg.Wait()
	fmt.Println(salutation)  // welcome ゴルーチンはそれらが作られたアドレス空間と同じ空間で実行する
}
