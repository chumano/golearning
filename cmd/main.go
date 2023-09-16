package main

import "fmt"
import "github.com/google/go-cmp/cmp"

func main() {
    fmt.Println("Hello, world1.")
    fmt.Println(cmp.Diff("Hello World", "Hello Go"))
}