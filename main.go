package main

import (
	"os"
)

func main() {
	playRound(os.Args[1:], os.Stdin, os.Stdout)
}
