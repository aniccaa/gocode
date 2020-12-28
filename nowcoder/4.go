package main

import (
	"fmt"
)

func main() {

	var s []string

	for i := 0; i < 2; i++ {
		var t string
		_, err := fmt.Scanf("%s", &t)
		if err != nil {
			return
		}
		s = append(s, t)
	}
	for _, v := range s {
		wrap(v)
	}
}

func addZero(s string) string {
	var ret [8]byte
	if len(s) < 8 {
		b := []byte(s)
		for i, v := range b {
			ret[i] = v
		}
		for i, v := range ret {
			if v == 0 {
				ret[i] = 48
			}
		}
		// fmt.Println(ret)
	}
	return string(ret[:])
}

func wrap(s string) {
	length := len(s)
	if length < 8 {
		fmt.Println(addZero(s))
		return
	}
	if length == 8 {
		fmt.Println(s)
		return
	}
	if length > 8 {
		char := []byte(s)
		for i := 0; i < length; i = i + 8 {
			if (i + 8) > length {
				fmt.Println(addZero(string(char[i:length])))
				return
			}
			fmt.Println(string(char[i : i+8]))
		}
	}

}
