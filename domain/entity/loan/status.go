package loan

// LoanSatus represents current status of Loan
type LoanStatus int

const (
	onLoan   LoanStatus = iota // 貸出中
	returned                   // 返却済み
	overdue                    // 延滞
)

func (ls LoanStatus) String() string {
	switch ls {
	case onLoan:
		return "貸出中"
	case returned:
		return "返却済み"
	case overdue:
		return "延滞"
	default:
		return "予期しない貸出状態"
	}
}
