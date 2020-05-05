package primitive

import (
	"fmt"
)

func Boolean(x bool) {

	fmt.Printf("%v's data type is boolean.\n", x)
}

func Integer() {
	var num int = 3
	var num_8 int8 = 10
	var num_16 int16 = 222
	var num_32 int32 = 8938923
	var num_64 int64 = 322323
	fmt.Printf("%v, %T\n", num, num)
	fmt.Printf("%v, %T\n", num_8, num_8)
	fmt.Printf("%v, %T\n", num_16, num_16)
	fmt.Printf("%v, %T\n", num_32, num_32)
	fmt.Printf("%v, %T\n", num_64, num_64)
}

func Bitoper() {
	a := 9  //1001
	b := 10 //1010
	fmt.Println(a & b)
	fmt.Println(a | b)
	fmt.Println(a ^ b)
	fmt.Println(a &^ b)
}

func Bitshift() {
	a := 8
	fmt.Println(a << 4)
	fmt.Println(a >> 3)
}
