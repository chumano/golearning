package utils

import (
	"fmt"

	"github.com/google/go-cmp/cmp"
)

func SayHello() {
	fmt.Println("Hello, world!")
}

func Compare() {
	fmt.Println(cmp.Diff("Hello", "Hello"))
	fmt.Println("=======================")
}
