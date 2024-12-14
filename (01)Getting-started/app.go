package main
//Package main: every go code file needs such a package instruction inside of it
// when we write go code it is split across packages, but atleast we need one package across a application.
// We can name this package anything but we are naming it main because it specifies the entry point of the application.


import "fmt"
// Eg, here we are importing the built in fmt package here.(these packages are a part of go installation).


func main() {
	// we are calling a fucntion PRINT which is printing text to standard command line
	//we create strings with double quotes or back ticks
	// this is named main because go should know which code to run when the execution starts
	fmt.Print("hello world")
}

// we can onlu build a file if we have a module associated with our app
// command to setup module: go mod init example.com/<name>
// this will create a module file ane then we can build our project with go build
// we will get an executable file which we can share and run it without installing go