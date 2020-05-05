package constant

import (
	"fmt"
)

// constant can also be shadowed
// the inner constant will shadow the outer constant
const sha int16 = 27

func Cons() {
	const a int = 14
	const b string = "foo"
	const c float32 = 3.14
	const d bool = true
	const sha int = 14
	fmt.Printf("%v %T\n", a, a)
	fmt.Printf("%v %T\n", b, b)
	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", d, d)
	fmt.Printf("%v %T\n", sha, sha)
}

//enumerated constant
func test_itoa() {

}
