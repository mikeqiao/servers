package main

import (
	f "manager/func"
	"runtime"

	"github.com/mikeqiao/newworld"
)

func main() {

	runtime.GOMAXPROCS(4)

	newworld.Start(f.FF)
}
