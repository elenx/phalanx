package main

import (
	"fmt"
	"phalanx/source"
)

func main() {
	a := "-- Start --"
	fmt.Printf("%v\n", a)
	source.Start()
}
