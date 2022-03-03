package main 

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main(){

	fmt.Println("Clouseres")
	//clouseres are functions that have access to variables outside their scope  //anonymous function
	//(scope is the area of code where the variable is accessible) 
	sum := adder()
	
	for i := 0; i < 10; i++ {
		fmt.Println(sum(i))
	}
}