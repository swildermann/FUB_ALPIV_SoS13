package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	rwm     sync.RWMutex
	balance int
)

// "reads" the current balance and prints it
// to the sdtout
func get(balance *int) {
	rwm.RLock()
	fmt.Println(*balance)
	rwm.RUnlock()
}

// "writes" the current balance: increases the
// balance by the given amount
func put(balance *int, amount int) {
	rwm.Lock()
	*balance += amount
	rwm.Unlock()
}

func main() {
	balance = 0
	go get(&balance)
	go put(&balance, 100)
	go get(&balance)
	go put(&balance, 100)
	go get(&balance)
	time.Sleep(1000)
}
