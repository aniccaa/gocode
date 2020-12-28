package main

import (
	"strings"
)

func compareVersion(version1, version2 string) int {
	var cv1, cv2 []string
	cv1 = strings.Split(version1, ".")
	cv2 = strings.Split(version2, ".")

	if len(cv1) > len(cv2) {
		// 1. 处理数组边界 2. 优化if分支
		for i, v := range cv1 {
			if v > cv2[i] {
				return 1
			} else if v == cv2[i] {
				return 0
			} else {
				return -1
			}

		}
	}
}
