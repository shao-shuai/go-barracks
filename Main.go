package main

import (
	"fmt"
	"go-barracks/constant"
	"go-barracks/primitive"
	"strconv"
)

// Declare variables at package level
var (
	actorName    string = "Elisabeth Sladen"
	companion    string = "Sarah Jane Smith"
	doctorNumber int    = 3
	season       int    = 11
)

// lower case varibale means this variable is scoped to this package
var scope int = 7

// uppper case variable means this variable is globally visible
var SCOPE int = 7

const (
	_  = iota // ignore first value by assigning to blank identifier
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	errorSpecialist = iota
	catSpecialist
	dogSpecialist
	snakeSpecialist
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)

func main() {
	// 3 ways to declare variables
	// first - declare a variable but not ready to use it
	var i int
	i = 31
	fmt.Println(i)

	// second - flexible to change variable type, more control
	var j int = 30
	var jj float32 = 30
	fmt.Println(j)
	fmt.Printf("%v %T\n", j, j)
	fmt.Printf("%v %T", jj, jj)

	// third
	k := 29
	kk := 29.
	fmt.Println(k)
	fmt.Printf("%v, %T\n", k, k)
	fmt.Printf("%v, %T\n", kk, kk)

	// shadowing
	fmt.Printf("Season now is %v\n", season)
	var season int = 99
	fmt.Printf("Season now is %v\n", season)

	// Declared variables must be used!
	// var use int = 999 this will cause compile time error

	// Variable conversion
	var inte int = 313
	fmt.Printf("%v, %T\n", inte, inte)

	var strin string = strconv.Itoa(inte) //Itoa is used to convert a integer to string
	fmt.Printf("%v, %T\n", strin, strin)

	primitive.Boolean(true)
	primitive.Integer()
	primitive.Bitoper()
	primitive.Bitshift()

	//Constants
	fmt.Printf("\n")
	const myConst int = 45
	fmt.Printf("%v %T\n", myConst, myConst)

	constant.Cons()

	fmt.Printf("%v\n", canSeeSouthAmerica)

	var roles byte = isAdmin | canSeeFinancials | canSeeEurope | canSeeSouthAmerica
	fmt.Printf("Below is the permission: \n")
	fmt.Printf("%b\n", roles)

	//Arrays
	grades := [3]int{97, 85, 93}
	fmt.Printf("Grades: %v\n", grades[0])

	primitive.MakeSlice()

}
