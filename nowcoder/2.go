package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var a string
	var b byte
	var num int
	reader := bufio.NewReader(os.Stdin)

	a, err := reader.ReadString('\n')
	if err != nil {
		return

	}
	b, err = reader.ReadByte()
	if err != nil {
		return
	}

	// char insensitive
	a = strings.ToLower(a)
	sliceChar := []byte(a)
	for _, v := range sliceChar {
		if b == v {
			num++
		}
	}
	fmt.Println(num)
}
