package main

import (
	"fmt"
	"os"
)

const VERSION = "0.1.0"

func goMain(args []string) int {
	fmt.Println("Hello World")
	return 0
}

func main() {
	status := goMain(os.Args)
	os.Exit(status)
}
