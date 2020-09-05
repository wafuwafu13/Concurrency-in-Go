package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(salutation) // <1> 開始する前にループが終了、参照を保持
		}()
	}
	wg.Wait()
}

// good day
// good day
// good day
