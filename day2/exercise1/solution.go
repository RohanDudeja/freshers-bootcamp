package main

import (
	"encoding/json"
	"fmt"
)
//for json output
type Ans struct {
	Arr map[string]int
}

func main() {
	arr := make(map[string]int)
	val := []string{"quick", "brown", "fox", "lazy", "dog"}
	str := make(chan string, len(val))
	done := make(chan bool)
	go func() {
		for {
			word, more := <-str
			if more {
				fmt.Print(word," ")
				for _,c := range word {
					arr[string(c)]=arr[string(c)]+1
				}
			} else {
				fmt.Println("\nreceived all words")
				done <- true
				return
			}
		}
	}()
	for j := 0; j < len(val); j++ {
		str <- val[j]
	}
	close(str)
	<-done
	fmt.Println(arr)

	// for json print format
	answer := &Ans{arr}
	json,err := json.Marshal(answer)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(json))
}
