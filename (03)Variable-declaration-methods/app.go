package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("how we can declare variables in a concise way")

	const inflationRate = 2.5

	//var investmentAmount float32, years float64 = 1000, 10: this is not allowed in go
	var investmentAmount, years float64 = 1000, 10
	// var years = 10
	var returnRate = 5.5

	//if we want got to infer a type from the go and dont want to mention var then we should declare the variable like this 
	randomValue := 10.5 // this is the shorter syntax in the go for declaration

	futureValue := (investmentAmount) * math.Pow((1+returnRate/100), (years))

	fmt.Println(futureValue, randomValue);
}