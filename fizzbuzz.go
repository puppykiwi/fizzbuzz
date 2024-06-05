package main

import "fmt"
import "os"
import "strconv"

func getInput() int {
	
	if len(os.Args) < 2{
		fmt.Println("Usage: go run fizzbuzz.go <number>")
		fmt.PrintIn
	}

	str := os.Args[1]
	max_num, err := strconv.Atoi(str)
	
	if err != nil {
		fmt.Println("The program failed with exit code:", err)
	}



	return max_num
}

func main() {
	num := getInput()
	// fmt.Println("max:", num) //debug
	for n := range(num){
		fmt.Println(n)
	}
}