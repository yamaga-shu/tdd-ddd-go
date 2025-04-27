package username

import "testing"

func TestLength(t *testing.T) {
	_, err := NewFirstName("aa")
	if err == nil {
		t.Error("expected error for firstname shorter than 3 characters, got nil")
	}

	_, err = NewFirstName("shu")
	if err != nil {
		t.Error("expected no error for Longer than 3 characters")
	}
}

func TestValidChar(t *testing.T) {
	_, err := NewFirstName("John@Doe")
	if err == nil {
		t.Error("expected error for firstname containing invalid characters, got nil")
	}

	_, err = NewFirstName("John_Doe")
	if err == nil {
		t.Error("expected error for firstname containing underscore, got nil")
	}

	_, err = NewFirstName("John-Doe")
	if err != nil {
		t.Error("expected no error for valid firstname with hyphen, got error")
	}

	_, err = NewFirstName("John123")
	if err != nil {
		t.Error("expected no error for valid firstname with numbers, got error")
	}
}
