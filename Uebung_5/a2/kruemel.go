package kruemel
import . "sync"

type Imp struct {
	cookies uint
	mutex Mutex
	credit *Cond
	}

func New() *Imp {
	box:= new(Imp) 
	box.credit = NewCond(&box.mutex)
	return box
}

func (box *Imp) Bake(c uint) {
	box.mutex.Lock()
	box.cookies += c
	box.credit.Signal()
	box.mutex.Unlock()
}

func (box *Imp) Eat(c uint) {
	box.mutex.Lock()
	for box.cookies < c {
		box.credit.Wait()
	}
	box.cookies -= c
	box.credit.Signal()
	box.mutex.Unlock()
}

func (box *Imp) Balance() uint {
	return box.cookies
}
