/* bahnhof.go */
package main

import (
	."sync"
	"fmt"
	"time"
)

type Imp struct {
	free []int
	val int
	cs, mutex Mutex
}

type Bahnhof struct {
	guard *Imp
	gleis []Mutex
}

func NewImp(n int) *Imp {
	x := new(Imp)
	x.val = n
	x.free = make([]int, n)
	for i := 0; i < n; i++ {
		x.free[i] = i
	}
	if n == 0 {
		x.cs.Lock()
	}
	return x
}

func (x *Imp) P() (i int) {
	x.cs.Lock()
	x.mutex.Lock()
	x.val--
	if x.val > 0 {
		x.cs.Unlock()
	}
	i = x.free[i]
	x.free = x.free[1:]
	x.mutex.Unlock()
	return i
}

func (x *Imp) V(i int) {
	x.mutex.Lock()
	x.val++
	if x.val == 1 {
		x.cs.Unlock()
	}
	x.free = append(x.free, i)
	x.mutex.Unlock()
}

func NewBahnhof(n int) *Bahnhof {
	b := new(Bahnhof)
	b.guard = NewImp(n)
	b.gleis = make([]Mutex, n)
	return b
}

func (b *Bahnhof) durchfahren(e chan bool) {
	i := b.guard.P()
	b.gleis[i].Lock()
	fmt.Printf("ich fahre hier auf Gleis %d!\n", i)
	fmt.Printf("Ich lasse mir jetzt 10 Sekunden Zeit..\n")
	x := time.Duration(10)
	time.Sleep(x * time.Second)
	b.gleis[i].Unlock()
	b.guard.V(i)
	e <- true
}

func main() {
	done1 := make(chan bool, 1)
	done2 := make(chan bool, 1)
	done3 := make(chan bool, 1)
	done4 := make(chan bool, 1)
	b := NewBahnhof(2)
	go b.durchfahren(done1)
	go b.durchfahren(done2)
	go b.durchfahren(done3)
	go b.durchfahren(done4)
	<- done1
	<- done2
	<- done3
	<- done4
}
