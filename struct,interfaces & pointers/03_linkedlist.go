package practice

// Node is a singly linked list node.
type Node struct {
	Val  int
	Next *Node
}

// InsertAtHead inserts val at the front of the list and returns the new
// head.
func InsertAtHead(head *Node, val int) *Node {
	// TODO: implement

	ll := Node{
		Val:  val,
		Next: head,
	}

	head = &ll
	return head
}

// InsertAtTail inserts val at the end of the list and returns the head.
// If head is nil, return a new single-node list.
func InsertAtTail(head *Node, val int) *Node {

	if head == nil {
		ll := Node{
			Val:  val,
			Next: nil,
		}
		return &ll
	}

	temp := head
	for temp.Next != nil {
		temp = temp.Next
	}

	ll := Node{
		Val:  val,
		Next: nil,
	}

	temp.Next = &ll
	return head
}

// ReverseList reverses the linked list in place (re-pointing Next fields,
// not copying nodes) and returns the new head.
func ReverseList(head *Node) *Node {
	// TODO: implement

	// var prev Node

	prev := head
	curr := head.Next

	for curr != nil {
		next := curr.Next

		curr.Next = prev
		prev = curr
		curr = next

	}

	return prev
}
