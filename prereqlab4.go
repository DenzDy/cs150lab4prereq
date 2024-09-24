package main
import (
	"fmt"
	"os"
	"strconv"
)

func main(){
	args := os.Args
	num1, _ := strconv.Atoi(args[1])
	num2, _ := strconv.Atoi(args[2])
	fmt.Println(num1 + num2)
	fmt.Println(num1 - num2)
	fmt.Println(num1 * num2)

}