package main

import "fmt"
import "os"
import "strconv"

func main(){
	
	if len(os.Args) < 2{
		fmt.Println("Usage: go run fizzbuzz.go <number>")
		return
	}

	str := os.Args[1]
	max_num, err := strconv.Atoi(str)
	
	if err != nil {
		fmt.Println("The program failed with exit code:", err)
		return
	}

	for i := 0; i < max_num; i++ {
		fmt.Print("hello, ", i)
	} 
}