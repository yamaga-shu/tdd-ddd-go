package member

import (
	"testing"

	"github.com/google/uuid"
)

func initMember() *Member {
	memberId, _ := uuid.NewV7()

	return New(memberId, "test member")
}

// TestSuspend verifies the Suspend method of the Member struct.
// It checks the initial state of the isSuspended field, ensures it is false initially,
// and verifies that it becomes true after calling the Suspend method.
func TestSuspend(t *testing.T) {
	mem := initMember()

	if mem.isSuspended {
		t.Errorf("Expected initial 'isSuspended' to be false, but got %t", mem.isSuspended)
	}

	mem.Suspend()

	if !mem.isSuspended {
		t.Errorf("Expected 'isSuspended' to be true after Suspend(), but got %t", mem.isSuspended)
	}
}

func TestCanBorrow(t *testing.T) {
	// initialize member
	mem := initMember()

	if !mem.CanBorrow() {
		t.Errorf("Expected 'CanBorrow' to be true, but got %t", mem.CanBorrow())
	}

	// increment until mem.LoanLimit.max(= 5)
	mem.LoanLimit.Increment() // mem.LoanLimit.current == 1
	mem.LoanLimit.Increment() // mem.LoanLimit.current == 2
	mem.LoanLimit.Increment() // mem.LoanLimit.current == 3
	mem.LoanLimit.Increment() // mem.LoanLimit.current == 4
	mem.LoanLimit.Increment() // mem.LoanLimit.current == 5

	if mem.CanBorrow() {
		t.Errorf("Expected 'CanBorrow' to be false when cannotLoanMore, but got %t", mem.CanBorrow())
	}

	// decrement
	mem.LoanLimit.Decrement()
	if !mem.CanBorrow() {
		t.Errorf("Expected 'CanBorrow' to be true, but got %t", mem.CanBorrow())
	}

	// suspend
	mem.Suspend()
	if mem.CanBorrow() {
		t.Errorf("Expected 'CanBorrow' to be false when member isSuspended, but got %t", mem.CanBorrow())
	}
}
