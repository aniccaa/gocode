package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var res []byte
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	t := []byte(strings.TrimRight(str, "\n"))

	for _, v := range t {
		if !isExistChar(res, v) {
			res = append(res, v)
		}
	}
	fmt.Print(len(res))
}

func isExistChar(a []byte, b byte) bool {
	for _, v := range a {
		if v == b {
			return true
		}
	}
	return false
}
