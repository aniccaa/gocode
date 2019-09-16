package main

import (
	"awesomeProject/mycli/cmd"
	"fmt"
)

func main() {
	// start here
	fmt.Println("hello")

	cmd.Execute()

	// end here
	fmt.Println("world.")
}
