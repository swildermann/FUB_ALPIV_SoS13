package main

import (
	"fmt"
	"time"
)

func main() {
	x := New()
	done1 := make(chan bool, 1)
	done2 := make(chan bool, 1)
	done3 := make(chan bool, 1)
	t := time.Duration(3)
	go x.essen(3, done1)
	time.Sleep(t * time.Second)
	go x.essen(2, done2)
	time.Sleep(t * time.Second)
	go x.backen(5, done3)
	<- done1
	<- done2
	<- done3
}

type Dose struct {
	kekse chan int
	klingel chan string
}

func New() *Dose {
	x := new(Dose)
	x.kekse = make(chan int, 1)
	x.klingel = make(chan string, 1)
	x.kekse <- 0
	return x
}

func (x *Dose) Recv() (n int) {
	n = <- x.kekse
	return n
}

func (x *Dose) Send(n int) {
	x.kekse <- n
}

func (x *Dose) backen(b int, e chan bool) {
	fmt.Printf("> %d Keks(e) werden gebacken!\n", b)
	n := x.Recv()
	fmt.Printf("> Anzahl der Kekse in der Dose: %d\n", n)
	x.Send(n + b)
	fmt.Printf("> Anzahl der Kekse nach dem Backen: %d\n", n + b)
	// wecke wartende Monster auf
	fmt.Printf("> Wecke wartende Monster: \"Ding!\"\n")
	x.klingel <- "Ding"
	e <- true
}

func (x *Dose) essen(b int, e chan bool) {
	fmt.Printf("%d Keks(e) sollen gegessen werden!\n", b)
	n := x.Recv()
	fmt.Printf("Anzahl der Kekse in der Dose: %d\n", n)
	for n < b {
		fmt.Printf("Lege die Kekse wieder zurueck :(\n")
		x.Send(n)
		fmt.Printf("Warte auf die Klingel...\n")
		// warte bis jemand wieder Kekse backt
		<- x.klingel
		fmt.Printf("Es hat geklingelt!\n")
		n = x.Recv()
		fmt.Printf("Anzahl der Kekse in der Dose nach dem Klingeln: %d\n", n)
	}
	// jetzt darf ich essen (n >= b)
	fmt.Printf("Ich darf %d Keks(e) essen! :) \n", b)
	x.Send(n - b)
	fmt.Printf("Anzahl der Kekse nach dem Essen: %d\n", n - b)
	// falls noch andere Monster gewartet hatten
	x.klingel <- "Dong"
	e <- true
}
