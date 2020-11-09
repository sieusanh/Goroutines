package main

import (
	"fmt"
	"sync"
	//"time"
	//"runtime"
)

func g1() {
	fmt.Println("g1")
	wg.Done()
}

func g2() {
	fmt.Println("g2")
	wg.Done()
}


var wg sync.WaitGroup	

func main() {
	// Example 1
	/*
	// go func()

	go g1()

	ng := runtime.NumGoroutine()
	fmt.Println(ng)

	time.Sleep(time.Second) // pause 1 second
	 */

	//------------------------------- 

	// Example 2:  using WaitGroup to control goroutines
	// Synchronized goroutines
	/**
		logic 1

		go g1()
		go g2()

		logic 2
	*/

	fmt.Println("begin") // logic 1

	wg.Add(2)

	go g1()
	go g2()

	wg.Wait()

	fmt.Println("end") // logic 2
}