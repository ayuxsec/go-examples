package main

import (
	"fmt"
)

// 1. We name the return variables 'result' and 'err' in the header.
func MagicDivider(a, b int) (result int, err error) {
	
	// 2. This runs LAST. It can see and modify 'err' before 
	// the caller ever gets it.
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("caught a panic: %v", r)
		}
	}()

	// 3. Instead of 'return a/b', we assign directly to the named variable.
	// If b is 0, this will panic, jumping straight to the defer above.
	result = a / b

	// 4. This is a "Naked Return". 
	// It automatically returns the current values of 'result' and 'err'.
	return 
}

func main() {
	// Success case
	val, err := MagicDivider(10, 2)
	fmt.Printf("Success: val=%d, err=%v\n", val, err)

	// Panic/Error case
	val, err = MagicDivider(10, 0)
	fmt.Printf("Failure: val=%d, err=%v\n", val, err)
}
