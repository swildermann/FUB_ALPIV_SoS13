package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Counter int
	done    chan bool
)

func v(p int64) { time.Sleep(time.Duration(rand.Int63n(p*1e5))) }

func inc(n *int,p  int64) {
	Accu := *n // "LDA n"
	v(p)
	Accu++ // "INA"
	v(p)
	*n = Accu //"STA n"
	fmt.Println(*n)
	v(p)
}

func count(p int64) {
	const N = 5
	for n := 0; n < N; n++ {
		inc(&Counter,p)
	}
	done <- true
}

func main() {
	Counter = 0
	done = make(chan bool)
	go count(800)
	go count(50)
	go count(1)
	go count(200)
	go count(30)
	<-done
	<-done
	<-done
	<-done
	<-done
	fmt.Printf("Zaehler = %d\n", Counter)
}
