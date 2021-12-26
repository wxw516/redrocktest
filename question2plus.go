package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func choice1(){
	mu.Lock()
	for a := 3;a <= 123456;a++{
		m := 0
		for i := 2;i < a;i++{
			if a%i == 0{
				m += 1
			}
		}
		if m == 0{
			fmt.Println(a)
		}
	}
	mu.Unlock()
}

func main()  {
	wg.Add(2)
	go choice1()
	go choice1()
	wg.Wait()
}

