package main
import "fmt"
func merge(a []int, m int) []int {
	n, b := len(a), make([]int, len(a))
	for j, i, k := 0, 0, m; j < n; j++ {
		if i < m && k < n {
			if a[i] < a[k] {
				b[j] = a[i]
				i++
			} else {
				b[j] = a[k]
				k++
			}
		} else if i < m {
			b[j] = a[i]
			i++
		} else if k < n {
			b[j] = a[k]
			k++
		}
	}
	copy(a, b)
	return a
}
func mergesort(a []int, e chan []int) {
	var ausgabe []int
	if len(a) > 1 {
		m := len(a) / 2
		c, d := make(chan []int), make(chan []int)
		go mergesort(a[:m], c)
		go mergesort(a[m:], d)
		<-c
		<-d
		ausgabe = merge(a, m)
	}
	e <- ausgabe
}
func main() {
	done := make(chan []int)
	numbers := []int{4, 3, 7, 4, 2, 5, 2, 7, 9, 1}
	go mergesort(numbers, done)
	ausgabe := <-done
	fmt.Print(ausgabe)
}
