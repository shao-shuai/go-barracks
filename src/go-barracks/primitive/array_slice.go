package primitive

import "fmt"

// The first number is inclusive and the second number is exclusive
func Slice() {
	// All a to e points to the same underlying data
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	b := a[:]
	c := a[3:]
	d := a[:6]
	e := a[3:6]
	a[5] = 999
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
}

func MakeSlice() {
	a := make([]int, 3, 100)
	fmt.Println(a)
	fmt.Printf("Length: %v\n", len(a))
	fmt.Printf("Capacity: %v\n", cap(a))
}
