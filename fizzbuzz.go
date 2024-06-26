package main

import "fmt"
import "os"
import "strconv"

func getInput() int {
	if len(os.Args) < 2{
		fmt.Println("Usage: go run fizzbuzz.go <number>")
		fmt.Println("Defaulting to max value of 100 \n")
		return 100
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
		if n % 3 == 0 && n % 5 == 0 {
			fmt.Println(n, "- FizzBuzz")
		} else if n % 3 == 0 {
		fmt.Println(n, "- Fizz")
		} else if n % 5 == 0 {
		fmt.Println(n, "- Buzz")
		} else{
			fmt.Println(n)
		}

	// Need to work to handle the 0 value result which defaults to the first if statement

	}
}