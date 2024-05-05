package main

import (
	"fmt"

	"github.com/franklange/go-scanln"
)

func main() {
	input := scanln.NewScanln()
	defer input.Stop()

	select {
	case s := <-input.C:
		fmt.Println(s)
	}
}
