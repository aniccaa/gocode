package main

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	ret := &ListNode{}
	cur := ret
	var carry int

	for l1 != nil || l2 != nil || carry > 0 {
		sum := carry

		if l1 != nil {
			sum = sum + l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum = sum + l2.Val
			l2 = l2.Next
		}

		carry = sum / 10

		cur.Next = &ListNode{Val: sum % 10}
		cur = cur.Next
	}
	return ret.Next

}

func main() {
	a := ListNode{3, nil}
	b := ListNode{4, &a}
	c := ListNode{2, &b}

	d := ListNode{4, nil}
	e := ListNode{6, &d}
	f := ListNode{5, &e}

	l := addTwoNumbers(&c, &f)
	fmt.Print(l)
}
