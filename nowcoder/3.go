package main

import (
	"fmt"
	"io"
	"sort"
)

func main() {
	var s []int
	for {
		var t int
		_, err := fmt.Scanf("%d", &t)
		if err != nil || err == io.EOF {
			break
		}
		s = append(s, t)
	}

	if s == nil {
		return
	}
	// 数组不为空的情况下。
	for i := 1; i < s[0]; i++ {
		for j := i + 1; j <= s[0]; j++ {
			if s[i] == s[j] {
				s[i] = -1
			}
		}
	}

	data := sort.IntSlice(s[1:])
	sort.Sort(data)
	for _, v := range data {
		if v != -1 {
			fmt.Println(v)
		}
	}
}
