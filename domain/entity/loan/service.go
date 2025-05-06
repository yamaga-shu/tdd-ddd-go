package loan

type IRepository interface {
	Store(l *Loan) error
}

type Service struct {
	loanRepository IRepository
}

func (s Service) Register(loan *Loan) error {
	err := s.loanRepository.Store(loan)
	if err != nil {
		return err
	}

	return nil
}
