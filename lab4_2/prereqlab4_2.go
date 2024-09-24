package main

import (
	"fmt"
	"log"
	"math/rand"
	"os/exec"
	"strconv"
	"strings"
)

func main() {
	program_output, err := exec.Command("./lab4", strconv.Itoa(rand.Intn(101)), strconv.Itoa(rand.Intn(101))).Output()
	if err != nil {
		log.Fatal(err)
	}
	var sum int = 0
	output_str := strings.Split(string(program_output), "\n")[:3]
	for _, elem := range output_str {
		x, err := strconv.Atoi(elem)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%d ", x)
		sum = sum + x
	}
	fmt.Println(sum)
}
