package main

import "fmt"

func main() {
	var a float64
	_, err := fmt.Scanf("%f", &a)
	if err != nil {
		return
	}
	t := int(a)
	if (a - float64(t)) >= 0.5 {
		fmt.Print(int(a + float64(0.5)))
	} else {
		fmt.Print(int(a))
	}
}
