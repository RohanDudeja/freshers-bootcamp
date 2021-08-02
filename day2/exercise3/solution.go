package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)
type account struct{
	Balance uint32
}
func main(){
	var mutex = &sync.Mutex{}
	acc := &account{500}

	for transaction := 0 ; transaction< 10 ; transaction++ {
		go func() {
			time.Sleep(time.Millisecond)
			mutex.Lock()
			amount:= uint32(rand.Intn(300))
			acc.Balance = acc.Balance + amount
			fmt.Println("amount added:",uint32(rand.Intn(100))," Balance :",acc.Balance)
			mutex.Unlock()
		}()
	}
	for transaction := 0 ; transaction< 10 ; transaction++ {
		go func() {
			time.Sleep(time.Millisecond)
			mutex.Lock()
			amount := uint32(rand.Intn(1000))
			if amount > acc.Balance {
				fmt.Println("withdrawal failed")
			}else {
				acc.Balance = acc.Balance - amount
				fmt.Println("amount removed: ",amount, "Balance: ",acc.Balance)
			}
			mutex.Unlock()
		}()
	}

	time.Sleep(time.Second)
	mutex.Lock()
	fmt.Println(acc.Balance)
	mutex.Unlock()

}

