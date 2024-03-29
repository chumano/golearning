package main

import "github.com/golearning/cmd"

func main() {
	app := cmd.App{}
	app.Setup()
	app.Run("localhost:8080")
}
