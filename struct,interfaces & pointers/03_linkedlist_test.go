package practice

import (
	"reflect"
	"testing"
)

// toSlice walks the list and collects its values, head to tail. Test helper only.
func toSlice(head *Node) []int {
	var out []int
	for n := head; n != nil; n = n.Next {
		out = append(out, n.Val)
	}
	return out
}

func TestInsertAtHead(t *testing.T) {
	var head *Node
	head = InsertAtHead(head, 3)
	head = InsertAtHead(head, 2)
	head = InsertAtHead(head, 1)

	want := []int{1, 2, 3}
	if got := toSlice(head); !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestInsertAtTail(t *testing.T) {
	var head *Node
	head = InsertAtTail(head, 1)
	head = InsertAtTail(head, 2)
	head = InsertAtTail(head, 3)

	want := []int{1, 2, 3}
	if got := toSlice(head); !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseList(t *testing.T) {
	head := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 4}}}}
	reversed := ReverseList(head)

	want := []int{4, 3, 2, 1}
	if got := toSlice(reversed); !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseSingleNode(t *testing.T) {
	head := &Node{Val: 42}
	reversed := ReverseList(head)
	want := []int{42}
	if got := toSlice(reversed); !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestReverseEmptyList(t *testing.T) {
	reversed := ReverseList(nil)
	if reversed != nil {
		t.Errorf("got %v, want nil", reversed)
	}
}
