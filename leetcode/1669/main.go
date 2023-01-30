package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */


type ListNode struct {
	     Val int
	     Next *ListNode
}
// 0->1->2->3->4->5->6
func mergeInBetween(list1 *ListNode, a int, b int, list2 *ListNode) *ListNode {
	index := 0
	aPoint := list1
	bPoint := list2
	for aPoint != nil {
		if index == a-1 {
			temp := aPoint.Next
			aPoint.Next = list2
			aPoint = temp
			index++
			continue
		}
		if index == b {
			temp := aPoint.Next
			for bPoint.Next != nil {
				bPoint = bPoint.Next
			}
			bPoint.Next = temp
			aPoint.Next = nil
			break
		}
		index++
		aPoint = aPoint.Next

	}
	return list1
}

func initList(n int) *ListNode {
	i := 0
	var head *ListNode = new(ListNode)
	point := head
	for ; i < n; i++ {
		node := new(ListNode)
		node.Val = i
		point.Next = node
		point = node
	}
	return head.Next
}

func printList(list *ListNode) {
	for point := list; point != nil; point = point.Next {
		fmt.Println(point.Val)
	}
}

func main() {
	list1 := initList(7)
	list2 := initList(5)
	mergeList := mergeInBetween(list1, 2, 5, list2)
	printList(mergeList)
}