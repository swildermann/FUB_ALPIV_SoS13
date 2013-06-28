package async

type ImpAsync struct {cS, cR chan byte; cA chan bool; in int}

func New (n uint) *ImpAsync {
  p:= new (ImpAsync)
  p.cS, p.cR, p.cA = make (chan byte), make (chan byte), make (chan bool)
  go func() {
    buffer := make([]byte, n)
    var count, in, out uint   
    for {
      if count == 0 {
        buffer[in] = <-p.cS
        in = (in+1) % n; count = 1
      } else if count == n {
        p.cR <- buffer[out]
        out = (out+1) % n; count = n-1
      } else { 
        select {
        case  buffer[in] = <- p.cS:
          in = (in + 1) % n; count++
        case <- p.cA:
          p.cR <- buffer[out]
          out = (out+1) % n; count--
        }
      }
    }
  }()
  return p
}

func (p *ImpAsync) Send (b byte) { p.cS <- b }

func (p *ImpAsync) Recieve() (b byte) { 
  p.cA <- true // Absicht senden, etwas entnehmen zu wollen
  return <- p.cR 
}
