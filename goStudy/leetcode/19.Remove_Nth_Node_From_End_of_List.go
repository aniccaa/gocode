package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ret := head
	d, isEnd := getDaddy(head, n)
	if isEnd {
		return head.Next
	}
	d.Next = d.Next.Next
	return ret
}

func getDaddy(head *ListNode, n int) (daddy *ListNode, isEnd bool) {
	daddy = head
	for head != nil {
		if n < 0 {
			daddy = daddy.Next
		}
		n--
		head = head.Next
	}
	isEnd = n == 0

	return
}

func main() {
	// removeNthFromEnd()
}
