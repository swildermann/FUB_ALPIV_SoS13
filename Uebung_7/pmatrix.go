package main
import "fmt"

const N = 3
type vector [N]int
type matrix [N]vector

func scalarproduct(v, w vector, p *int, d chan bool) {
	for j := 0; j < N; j++ {
		*p += v[j] * w[j]
	}
}
func column(a matrix, k int) (s vector) {
	for j := 0; j < N; j++ {
		s[j] = a[j][k]
	}
	return s
}
func product(a, b matrix) (p matrix) {
	done := make(chan bool)
	for i := 0; i < N; i++ {
		for k := 0; k < N; k++ {
			go scalarproduct(a[i], column(b, k), &p[i][k], done)
		}
	}
	for j := 0; j < N*N; j++ {
		<-done
	}
	return p
}
func toString(a matrix) (s string) {
	i := a[1][1]
	return "Test"
}
func main() {
	a := matrix{vector{1, 2, 9}, vector{4, 5, 6}, vector{7, 8, 9}}
	b := matrix{vector{9, 8, 7}, vector{6, 5, 4}, vector{3, 2, 1}}
	c := product(a, b)
	fmt.Printf("Hello\n" + c)
}
