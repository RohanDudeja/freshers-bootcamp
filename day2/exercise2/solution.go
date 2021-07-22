package main

import (
	"fmt"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var sum uint32
	for student := 1; student <= 200; student++ {
		wg.Add(1)
		go func(){
			atomic.AddUint32(&sum, uint32(rand.Intn(10)))
			time.Sleep(time.Duration(rand.Int()))
			wg.Done()
		}()
		// wg.Done() // doubt: using wg.Done() here gives random output why!?
	}
	wg.Wait()
	var average uint32
	average = sum/200
	fmt.Println(average)
}
