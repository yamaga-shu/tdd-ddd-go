package member

import "github.com/google/uuid"

// Member: 会員
type Member struct {
	memberId    uuid.UUID  // 会員のID
	name        string     // 氏名
	isSuspended bool       // 貸出停止中かどうか
	LoanLimit   *loanLimit // 貸出上限管理
}

func New(memberId uuid.UUID, name string) *Member {
	return &Member{
		memberId:    memberId,
		name:        name,
		isSuspended: false,
		LoanLimit:   newLoanLimit(0),
	}
}

func (m *Member) CanBorrow() bool {
	return !m.isSuspended && m.LoanLimit.canLoanMore()
}

func (m *Member) Suspend() {
	m.isSuspended = true
}
