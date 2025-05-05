package loan

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/yamaga-shu/tdd-ddd-go/domain/entity/member"
)

// Loan represents Domain model of Loan
type Loan struct {
	loanId          uuid.UUID  // 貸出のID
	memberId        uuid.UUID  // 借りた会員のID
	bookInventoryId uuid.UUID  // 借りられた蔵書のID
	loanDate        time.Time  // 貸出日
	dueDate         time.Time  // 返却予定日(貸出日 + 14日)
	returnDate      time.Time  // 実際に返却された日
	extended        bool       // 延長の有無(延長は1度のみ)
	status          LoanStatus // 現在の貸出状態

	// 依存性の注入
	repositories struct {
		loanRepository   IRepository        // 貸出リポジトリ
		memberRepository member.IRepository // 会員リポジトリ
	}
}

func New(loanId, memberId, bookInventoryId uuid.UUID, now time.Time, loanRepository IRepository, memberRepository member.IRepository) *Loan {
	return &Loan{
		loanId:          loanId,
		memberId:        memberId,
		bookInventoryId: bookInventoryId,
		loanDate:        now,
		dueDate:         now.AddDate(0, 0, 14),
		extended:        false,
		status:          onLoan,

		// 依存性の注入
		repositories: struct {
			loanRepository   IRepository
			memberRepository member.IRepository
		}{
			loanRepository:   loanRepository,
			memberRepository: memberRepository,
		},
	}
}

// Register attempts to create a new loan record in the loanRepository.
// It returns an error if the creation process fails.
func (l *Loan) Register() error {
	err := l.repositories.loanRepository.Create(l)
	if err != nil {
		return err
	}

	return nil
}

// Return updates the loan's return date and status to indicate the book has been returned
func (l *Loan) Return(now time.Time) error {
	if l.loanDate.After(now) {
		return errors.New("returnDate must be after loanDate")
	}
	if l.status == returned {
		return errors.New("this loan has been already returned")
	}
	mem, err := l.repositories.memberRepository.GetMemberById(l.memberId)
	if err != nil {
		return err
	}

	// decrement member's loanLimit
	mem.LoanLimit.Decrement()

	// update returnDate
	l.returnDate = now

	// update loanStatus
	l.status = returned

	return nil
}

// Extend extends the due date of the loan by 7 days if it hasn't been extended before.
func (l *Loan) Extend() error {
	if l.extended {
		return errors.New("dueDate has been already extended")
	}

	// extend dueDate
	l.dueDate = l.dueDate.AddDate(0, 0, 7)
	l.extended = true

	return nil
}
