package member

type loanLimit struct {
	max     int
	current int
}

func newLoanLimit(current int) *loanLimit {
	return &loanLimit{
		max:     5,
		current: current,
	}
}

func (ll *loanLimit) canLoanMore() bool {
	return ll.current < ll.max
}

func (ll *loanLimit) Increment() {
	ll.current++
}

func (ll *loanLimit) Decrement() {
	if ll.current <= 0 {
		ll.current = 0
		return
	}

	ll.current--
}
