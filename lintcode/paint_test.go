package lintcode

import "testing"

type actual struct {
	costs [][]int
	fee   int
}

func TestMinCostII(t *testing.T) {
	cases := []struct {
		name     string
		descript string
		// input
		costs [][]int
		fee   int
	}{
		{
			name:     "",
			descript: "",
			costs:    [][]int{{14, 2, 11}, {11, 14, 5}, {14, 3, 10}},
			fee:      10,
		},
		{
			costs: [][]int{},
			fee:   0,
		},
		{
			costs: [][]int{{1, 2, 100}, {1001, 1002, 1}},
			fee:   2,
		},
	}

	for _, c := range cases {
		act := actual{}
		act.fee = minCostII(c.costs)
		if act.fee != c.fee {
			t.Errorf("Actual is %d\n", act.fee)
			t.Errorf("Expect is %d, anwser wrong!", c.fee)
		}
	}
}
