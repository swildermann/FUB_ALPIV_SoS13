const (p_i = i ; nK = 5)
var (nP[5] bool; x *Imp)

func c(k uint) bool{
	return nP[k+1 mod 5] == false && nP[k-1 mod 5]
}

func in(x Any, k uint){
	nP[k] = true	
}

func out(x Any, k uint){
	nP[k] = false
}
func (x *ImpCS) PhilosophEat() { x.Enter(p_i, nil) }

func main() { x = New(nK, c, in, out) }
