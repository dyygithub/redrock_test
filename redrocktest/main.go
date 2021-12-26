package main

import (
	"fmt"
	"sync"
)

func main(){
	var mu sync.Mutex
	go func(){
		fmt.Println("有点强人锁男")
		mu.Lock()
	}()
	mu.Unlock()
}
//出现错误的原因：没有相应的互斥锁。


