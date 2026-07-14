package practice

import (
	"reflect"
	"sort"
	"testing"
)

func TestByAgeSort(t *testing.T) {
	people := ByAge{
		{Name: "Carol", Age: 35},
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
	}

	sort.Sort(people)

	want := ByAge{
		{Name: "Bob", Age: 25},
		{Name: "Alice", Age: 30},
		{Name: "Carol", Age: 35},
	}

	if !reflect.DeepEqual(people, want) {
		t.Errorf("got %v, want %v", people, want)
	}
}

func TestByAgeIsStableForEqualAges(t *testing.T) {
	people := ByAge{
		{Name: "X", Age: 20},
		{Name: "Y", Age: 20},
	}
	sort.Sort(people)
	if people[0].Age != 20 || people[1].Age != 20 {
		t.Errorf("ages were changed by sorting: %v", people)
	}
}
