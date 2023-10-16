package a

import "github.com/informalsystems/experiments/internalized/internal/b"

func MyFunc() b.InternalType {
	return b.InternalType{A: "Hello", B: "world!"}
}
