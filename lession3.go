package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var result *ListNode
	var tmp *ListNode
	var next *ListNode
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	for {
		switch {
		case list1 != nil && list2 != nil:
			if list1.Val < list2.Val {
				//travese list 1
				next = list1
				//next list 1
				list1 = list1.Next
			} else {
				//travese list 2
				next = list2
				//next list 2
				list2 = list2.Next
			}
		case list1 == nil:
			next = list2
		case list2 == nil:
			next = list1
		default:
			next = nil
		}
		if next == nil {
			return result
		}
		//push to result
		if result == nil {
			result = next
			tmp = result
		} else {
			tmp.Next = next
			tmp = next
		}
		//has null in of list
		if list1 == nil && tmp != list2 {
			tmp.Next = list2
			return result
		}
		if list2 == nil && tmp != list1 {
			tmp.Next = list1
			return result
		}
	}
	return result
}

func maxProfit(prices []int) int {
	min := 0
	max := 0
	can := 0
	for i, p := range prices {
		if p < prices[can] {
			can = i
		}
		if p-prices[can] > prices[max]-prices[min] {
			min = can
			max = i
		}
	}
	if max > min {
		return prices[max] - prices[min]
	}
	return 0
}

func isPalindrome(head *ListNode) bool {
	if head == nil {
		return true
	}
	cur := head
	traveled := make([]int, 0)
	for {
		traveled = append(traveled, cur.Val)
		cur = cur.Next
		if cur == nil {
			break
		}
	}
	for idx := 0; idx < len(traveled)/2; idx++ {
		if traveled[idx] != traveled[len(traveled)-1-idx] {
			return false
		}
	}
	return true
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return 0
		}
		return -1
	}
	idx := len(nums) / 2
	if nums[idx] == target {
		return idx
	}
	if nums[idx] < target {
		return searchRight(nums[idx:], idx, target)
	}
	if nums[idx] > target {
		return searchLeft(nums[0:idx], 0, target)
	}
	return -1
}

func searchLeft(nums []int, from, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return from
		}
		return -1
	}
	idx := len(nums) / 2
	if nums[idx] == target {
		return from + idx
	}
	if nums[idx] < target {
		return searchRight(nums[idx:], from+idx, target)
	}
	if nums[idx] > target {
		return searchLeft(nums[0:idx], from, target)
	}
	return -1
}

func searchRight(nums []int, from int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	if len(nums) == 1 {
		if nums[0] == target {
			return from
		}
		return -1
	}
	idx := len(nums) / 2
	if nums[idx] == target {
		return from + idx
	}
	if nums[idx] < target {
		return searchRight(nums[idx:], from+idx, target)
	}
	if nums[idx] > target {
		return searchLeft(nums[0:idx], from, target)
	}
	return -1
}
