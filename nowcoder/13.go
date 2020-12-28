package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, err := reader.ReadString('\n')
	if err != nil {
		return
	}
	strArr := strings.Split(strings.TrimRight(str, "\n"), " ")
	t := revertWord(strArr)
	for _, v := range t {
		fmt.Printf("%s ", v)
	}

}

func revertWord(a []string) []string {
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
	return a
}
