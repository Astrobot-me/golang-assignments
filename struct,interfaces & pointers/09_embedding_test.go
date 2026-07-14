package practice

import "testing"

func TestBaseDescribe(t *testing.T) {
	b := Base{ID: 7}
	if got := b.Describe(); got != "record #7" {
		t.Errorf("Describe() = %q, want %q", got, "record #7")
	}
}

func TestUserPromotedFieldsAndMethod(t *testing.T) {
	u := User{Base: Base{ID: 1}, Name: "Aditya"}

	// promoted field access — no need to write u.Base.ID
	if u.ID != 1 {
		t.Errorf("u.ID = %d, want 1", u.ID)
	}

	// promoted method call — no need to write u.Base.Describe()
	if got := u.Describe(); got != "record #1" {
		t.Errorf("Describe() = %q, want %q", got, "record #1")
	}
}

func TestProductPromotedFieldsAndMethod(t *testing.T) {
	p := Product{Base: Base{ID: 42}, Price: 9.99}

	if p.ID != 42 {
		t.Errorf("p.ID = %d, want 42", p.ID)
	}
	if got := p.Describe(); got != "record #42" {
		t.Errorf("Describe() = %q, want %q", got, "record #42")
	}
}
