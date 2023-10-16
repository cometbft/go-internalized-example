package main

import (
	"fmt"

	"github.com/informalsystems/experiments/internalized/a"
)

func main() {
	t := a.MyFunc()
	fmt.Printf("%s %s\n", t.A, t.B)
}
