package username

import "testing"

func TestInit(t *testing.T) {
	_, err := NewFirstName("aa")
	if err == nil {
		t.Error("expected error for firstname shorter than 3 characters, got nil")
	}
}
