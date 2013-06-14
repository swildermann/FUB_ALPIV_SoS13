package main
import "fmt"

/* Leider habe ich es nicht puenktlich
geschafft den Code erfolgreich zum compilieren zu bringen*/ 

func mult(matrix1 [][]int, matrix2 [][]int, e chan bool){

    newMa := [][]int{{0,0,0},{0,0,0},{0,0,0}}
    for i := 0 ; i < len(matrix1) ;i++{
	
        for j := 0 ; j < len(matrix2[0]) ;j++ {
	    w :=  make (chan bool)
            go rowCol(i,j,matrix1,matrix2,newMa,w)
	    <-w
	}
	 
    }
     fmt.Println(newMa)
    e <- true
}

func rowCol( r int , c int, matrix1 [][]int, matrix2 [][]int, newMa [][]int,  e chan bool){
    var sum int = 0
    for i := 0 ; i < len(matrix1[1]) ; i++ {
        sum += matrix1 [r][i] * matrix2 [i][c]
    }
     newMa[r][c] = sum
    e<-true

}

func main(){
    matrix1 := [][]int { {1,2,3} ,
                         {4,5,6} , 
                         {7,8,9} }
    matrix2 := [][]int {{0,0,1} , 
                         {0,1,0} , 
                         {1,0,0} }

  q := make (chan bool)
  mult(matrix1 , matrix2, q)
  <-q
}
