package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("vaariables and operators")
	var investmentAmount = 1000
	var years = 10
	var returnRate = 5.5

	// we can give the data type to the variables
	var period float64 = 109

	//go is a statically typed language, so we need to specify the type explicitely when we are using different types.

	var futureValue = float64(investmentAmount) * math.Pow((1+returnRate/100), float64(years))

	fmt.Println(futureValue, period);
}
