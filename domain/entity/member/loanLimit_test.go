package member

import "testing"

func TestLoanLimitCurrent(t *testing.T) {
	// initalize loanLimit
	loanLim := newLoanLimit(0)
	if loanLim.current != 0 {
		t.Errorf("Expected current is 0, but got %d", loanLim.current)
	}

	// increment loanLimit.current
	loanLim.Increment()
	if loanLim.current != 1 {
		t.Errorf("Expected current is 1, but got %d", loanLim.current)
	}

	// decrement loanLimit.current
	loanLim.Decrement()
	if loanLim.current != 0 {
		t.Errorf("Expected current is 0, but got %d", loanLim.current)
	}

	// decrement again
	loanLim.Decrement()
	if loanLim.current != 0 {
		t.Errorf("Expected current is positive, but got %d", loanLim.current)
	}
}

func TestCanLoanMore(t *testing.T) {
	// initalize loanLimit
	loanLim := newLoanLimit(4)
	if !loanLim.canLoanMore() {
		t.Errorf("Expected 'canLoanMore' is true, but got %t", loanLim.canLoanMore())
	}

	// increment -> loanLim.current must be 5 (== loanLim.max)
	loanLim.Increment()
	if loanLim.canLoanMore() {
		t.Errorf("Expected 'canLoanMore' is false, but got %t", loanLim.canLoanMore())
	}
}
