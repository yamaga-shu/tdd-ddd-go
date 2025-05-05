package loan

type IRepository interface {
	Create(l *Loan) error
}
