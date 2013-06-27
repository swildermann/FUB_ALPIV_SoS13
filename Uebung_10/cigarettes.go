package main

import (
	"fmt"
	"time"
	"math/rand"
)

const (
//declare the states - what the host can bring
	PapersTobacco = iota
	PapersLucifers = iota
	TobaccoLucifers = iota
	Empty = iota
)

func main() {
// set the values and start the schedular,
//who starts the game
	var waiter = make(chan int,1)
	var table = make(chan int,1)
	waiter <-1
	scheduler(waiter,table)
}

func scheduler(waiter chan int, table chan int){
// check always for waiting guests, which  need
// to be served
	for {
		<-waiter
		go givesomething(table)
		go getsomething(table,waiter)
	}

}

func givesomething(cs chan int){
//the host gives something to the table
//which is randomly choosen
	rand.Seed(time.Now().UnixNano())
	cs<-rand.Intn(Empty)
	fmt.Println("=== HOST IS SERVING ===")
	time.Sleep(2*1e9)
}


func getsomething(cs chan int, waits chan int){
//the smokers need to check who is able to smoke
//so they check whether there is something on the table
// and what it is  and decide because of that
	number := <-cs
	time.Sleep(2*1e9)
	if number==PapersTobacco {
		fmt.Println("Thats what PETER  was looking for - lets begin to smoke")
	}else if number==PapersLucifers {
		fmt.Println("Thats what JAN was looking for - lets begin to smoke")
	}else if number==TobaccoLucifers{
		fmt.Println("Thats what ADRIAN was looking for - lets begin to smoke")
	}else{
		// this state should never be reached
		fmt.Println("Waiting for something that fits")
	}
	fmt.Println("===SMOOKING===");
	time.Sleep(2*1e9)
	fmt.Println("===END OF SMOOKING===")
	time.Sleep(2*1e9)
	//having smoked, they call the host again and wait 
	waits <-1
}

