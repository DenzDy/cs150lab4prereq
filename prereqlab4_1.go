package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		log.Fatal("ERROR")
	}
	num1, _ := strconv.Atoi(args[1])
	num2, _ := strconv.Atoi(args[2])

	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)
}
