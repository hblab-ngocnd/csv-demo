package main

import (
	"fmt"
	"testing"
)

func Test_mergeTwoLists(t *testing.T) {
	list1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}
	list2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	for cur := mergeTwoLists(list1, list2); cur != nil; {
		fmt.Println(cur)
		cur = cur.Next
	}
}
func Test_mergeTwoLists2(t *testing.T) {
	var list1 *ListNode
	list2 := &ListNode{Val: 0}
	for cur := mergeTwoLists(list1, list2); cur != nil; {
		fmt.Println(cur)
		cur = cur.Next
	}
}

func Test_maxProfit(t *testing.T) {
	fmt.Println(maxProfit([]int{7, 1, 5, 3, 6, 4}))
	fmt.Println(maxProfit([]int{2, 4, 1}))
	fmt.Println(maxProfit([]int{2, 4, 1, 4}))
}

func Test_isPalindrome(t *testing.T) {
	input := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1}}}}}
	fmt.Println(isPalindrome(input))
}

func Test_search(t *testing.T) {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 3))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2))
	fmt.Println(search([]int{5}, -5))
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 12))
}
