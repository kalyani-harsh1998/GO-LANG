package main

import  "fmt"

func main() {
	fmt.Println("how to fetch input from the user")

	var num float64

	//scan is used to fetch the input from the user and we use the '&' to point to the variable for which we want to scan the value
	fmt.Print("enter number")
	fmt.Scan(&num);

	fmt.Print(num);
}