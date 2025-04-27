package user

import "testing"

func TestEquals(t *testing.T) {
	nameA := NewFullName("LeBron", "James")
	nameB := NewFullName("Kevin", "Durant")
	nameC := NewFullName("LeBron", "James")

	if nameA.Equals(nameB) == true {
		t.Errorf("nameA.Equals(nameB) == true, wanted false")
	}

	if nameA.Equals(nameC) != true {
		t.Errorf("nameA.Equals(namec) == false, wanted true")
	}
}
