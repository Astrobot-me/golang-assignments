package practice

import (
	"reflect"
	"testing"
)

func TestDeepCopyValuesMatch(t *testing.T) {
	original := Profile{Name: "Aditya", Tags: []string{"go", "backend"}}
	dup := DeepCopy(original)

	if !reflect.DeepEqual(original, dup) {
		t.Errorf("copy = %v, want %v", dup, original)
	}
}

func TestDeepCopyIsIndependent(t *testing.T) {
	original := Profile{Name: "Aditya", Tags: []string{"go", "backend"}}
	dup := DeepCopy(original)

	dup.Tags[0] = "MUTATED"
	dup.Name = "Changed"

	if original.Tags[0] == "MUTATED" {
		t.Errorf("mutating copy.Tags affected original: %v", original.Tags)
	}
	if original.Name == "Changed" {
		t.Errorf("mutating copy.Name affected original: %v", original.Name)
	}
}

func TestDeepCopyAppendIsIndependent(t *testing.T) {
	original := Profile{Name: "Aditya", Tags: []string{"go"}}
	dup := DeepCopy(original)

	dup.Tags = append(dup.Tags, "new-tag")

	if len(original.Tags) != 1 {
		t.Errorf("appending to copy.Tags affected original: %v", original.Tags)
	}
}
