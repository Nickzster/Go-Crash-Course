package main

import "fmt"

func main(){
// MAIN TYPES
// string
// bool
// int | int8 | int16 | int32 | int64
// uint | uint8 | uint16 | uint32 | uint64 | uintptr
// byte - alias for uint8
// rune - alias for int32
// float32 | float64
// complex64 | complex128

var name string = "Nick"
var age int = 25
hobby := "Coding" //Shorthand syntax must be used in function body. Is declared implicitly with "var". Types cannot change with reassignment.
const isCool = true





fmt.Println(name, age, isCool, hobby)
// %T is a "verb". Refer to google for more verbs for Go.
fmt.Printf("%T\n", name)

foo, bar := "Foo", "Bar"

fmt.Println(foo, bar)




}