/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var carrier int
	var head, tail *ListNode
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		cur := &ListNode{}
		sum := (n1 + n2 + carrier)
		cur.Val, carrier = sum%10, sum/10
		if head == nil {
			head = cur
			tail = cur
		} else {
			tail.Next = cur
			tail = tail.Next
		}
	}
	if carrier != 0 {
		tail.Next = &ListNode{Val: carrier}
	}
	return head
}

// @lc code=end

