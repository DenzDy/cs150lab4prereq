package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	m := map[string]struct{}{}
	args := os.Args
	if len(args) != 6 {
		log.Fatal("ERROR")
	}
	for i := 1; i < 6; i++ {
		m[args[i]] = struct{}{}
	}
	for k := range m {
		fmt.Printf("%s ", k)
	}
	fmt.Println()
}
