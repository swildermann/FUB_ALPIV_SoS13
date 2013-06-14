package main

import (
	"fmt"; "sync"; "time"; "runtime"
)

var (
	rwm     sync.RWMutex
	balance int
)

// "reads" the current balance and prints it
// to the sdtout
func get(n int, balance *int) {
	rwm.RLock()
	fmt.Println("Reader",n,*balance)
	rwm.RUnlock()
}

// "writes" the current balance: increases the
// balance by the given amount
func put(n int,balance *int, amount int) {
	rwm.Lock()
	*balance += amount
	fmt.Println("Writer",n,"added",amount)
	rwm.Unlock()
}

func main() {
	balance = 0			// set initial value
	runtime.GOMAXPROCS(5)
	go get(1,&balance)		// start thread 1 of 5...	
	go put(1,&balance, 100)
	go get(2,&balance)
	go put(2,&balance, 100)
	go get(3,&balance)
	time.Sleep(10e5)		// wait until threads are supposedly done
	fmt.Println("balance:",balance)	// print balance 
}
