package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) Append(node *ListNode) *ListNode {
	if l == nil {
		l = node
		return l
	}

	l.Next = node
	return l
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	var outputList *ListNode
	for list1.Next != nil || list2.Next != nil {
		if list1.Val == list2.Val {
			outputList = outputList.Append(&ListNode{Val: list1.Val})
			outputList = outputList.Append(&ListNode{Val: list2.Val})
			list1, list2 = list1.Next, list2.Next
			continue
		}

		if list1.Val < list2.Val {
			outputList = outputList.Append(&ListNode{Val: list1.Val})
			list1 = list1.Next
			continue
		}

		outputList = outputList.Append(&ListNode{Val: list2.Val})
		list2 := list2.Next
	}

	return outputList
}

func main() {

}
