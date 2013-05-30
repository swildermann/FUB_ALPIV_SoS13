package channel

import "fmt"
import "sync"

type Kanal struct {
	botschaft byte 
	s,e,mutex Mutex
	}

func send(x Kanal, c byte){
	x.mutex.wait()
	x.botschaft = c
	x.s.signal()
	x.e.wait()
}

func init(x Kanal){
	x.s = 0
	x.e = 0
	x.mutex = 1
}

func recv(x Kanal, c byte){
	x.s.wait()
	c=x.botschaft
	x.e.signal()
	s.mutex.signal()
}
