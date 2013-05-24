package rw
import . "sync"

type Imp struct {
	rc, rq, wq uint
	mutex Mutex
	OKtowrite *Cond
	OKtoread *Cond
	busy bool
	}

func New() *Imp {
	x:= new(Imp) 
	x.OKtowrite = NewCond(&x.mutex)
	x.OKtoread = NewCond(&x.mutex)
	x.busy = false
	x.rc, x.rq, x.wq = 0, 0 ,0
	return x
}

func (x *Imp) Reader_start() {
	x.mutex.Lock()
	if x.busy || x.wq != 0 { 
		x.OKtoread.Wait() 
		x.rq ++
	}
	x.rc ++
	x.OKtoread.Signal()
	x.mutex.Unlock()
}

func (x *Imp) Reader_end() {
	x.mutex.Lock()
	x.rc -- 
	if x.rc == 0 { x.OKtowrite.Signal() }
	x.mutex.Unlock()
}

func (x *Imp) Writer_start() {
	x.mutex.Lock()
	if x.rc != 0 || x.busy { 
		x.OKtowrite.Wait()
		x.wq ++
	}
	x.busy = true
	x.mutex.Unlock()
}

func (x *Imp) Writer_end() {
	x.mutex.Lock()
	x.busy = false
	if x.rq != 0 { 
		x.OKtoread.Signal() 
	} else {
		x.OKtowrite.Signal()
	}
	x.mutex.Unlock()
}


