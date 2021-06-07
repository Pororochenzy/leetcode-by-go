package leetcode

import (
	"github.com/halfrost/LeetCode-Go/structures"
)

// ListNode define
type ListNode = structures.ListNode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 解法一
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	var fast, slow *ListNode
	fast = head
	slow = head
	step := 0
	for i := 0; i < n; i++ {
		// n maybe much larger than length of linklist
		if fast.Next == nil && step != 0 && step < n-1 {
			return head
		}
		fast = fast.Next
		step++
	}
	if fast == nil {
		head = head.Next
		return head
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return head
}

// 解法二
func removeNthFromEnd1(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}
	if n <= 0 {
		return head
	}
	current := head
	len := 0
	for current != nil {
		len++
		current = current.Next
	}
	if n > len {
		return head
	}
	if n == len {
		current := head
		head = head.Next
		current.Next = nil
		return head
	}
	current = head
	i := 0
	for current != nil {
		if i == len-n-1 {
			deleteNode := current.Next
			current.Next = current.Next.Next
			deleteNode.Next = nil
			break
		}
		i++
		current = current.Next
	}
	return head
}
