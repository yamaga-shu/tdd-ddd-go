package loan

// LoanSatus represents current status of Loan
type LoanStatus int

const (
	onLoan   LoanStatus = iota // 貸出中
	returned                   // 返却済み
	overdue                    // 返却期限超過
)
