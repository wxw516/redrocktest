package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func choice(){
	mu.Lock()
	for a := 6;a <= 123456;a++{
		m := 0
		for i := 1;i <= a/2;i++{
			if a%i == 1{
				m += i
			}
		}
		if m == a{
			fmt.Println(a)
		}
	}
	mu.Unlock()
}

func main()  {
	wg.Add(2)
	go choice()
	go choice()
	wg.Wait()
}
