package practice

import "testing"

func TestHasCycleNoCycle(t *testing.T) {
	head := &Node{Val: 1, Next: &Node{Val: 2, Next: &Node{Val: 3}}}
	if HasCycle(head) {
		t.Errorf("HasCycle() = true, want false")
	}
}

func TestHasCycleEmptyList(t *testing.T) {
	if HasCycle(nil) {
		t.Errorf("HasCycle(nil) = true, want false")
	}
}

func TestHasCycleWithCycle(t *testing.T) {
	third := &Node{Val: 3}
	second := &Node{Val: 2, Next: third}
	first := &Node{Val: 1, Next: second}
	third.Next = first // creates a cycle back to the head

	if !HasCycle(first) {
		t.Errorf("HasCycle() = false, want true")
	}
}

func TestHasCycleSelfLoop(t *testing.T) {
	node := &Node{Val: 1}
	node.Next = node // points to itself

	if !HasCycle(node) {
		t.Errorf("HasCycle() = false, want true")
	}
}
