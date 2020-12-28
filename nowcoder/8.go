package main

import (
	"fmt"
	"sort"
)

func main() {
	var a int
	var k, v int
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		return
	}
	b := make(map[int]int, 0)
	for i := 0; i < a; i++ {
		_, err := fmt.Scanf("%d %d", &k, &v)
		if err != nil {
			return
		}
		// 判断map中key是否存在。
		if _, ok := b[k]; ok {
			b[k] = b[k] + v
		} else {
			b[k] = v
		}
	}

	//对key进行排序。
	var t []int
	for key := range b {
		t = append(t, key)
	}
	sort.Ints(t)

	for _, val := range t {
		fmt.Printf("%d %d\n", val, b[val])
	}
}
