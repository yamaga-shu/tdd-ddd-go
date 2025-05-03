package loan

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Loan represents Domain model of Loan
type Loan struct {
	loanId          uuid.UUID  // 貸出のID
	memberId        uuid.UUID  // 借りた会員のID
	bookInventoryId uuid.UUID  // 借りられた本の蔵書ID
	loanDate        time.Time  // 貸出日
	dueDate         time.Time  // 返却予定日(貸出日 + 14日)
	returnDate      time.Time  // 返却された日
	status          LoanStatus // 現在の貸出状態
	extended        bool       // 延長は1度のみ
}

func New(loanId, memberId, bookInventoryId uuid.UUID, now time.Time) *Loan {
	return &Loan{
		loanId:          loanId,
		memberId:        memberId,
		bookInventoryId: bookInventoryId,
		loanDate:        now,
		dueDate:         now.AddDate(0, 0, 14),
		status:          onLoan,
		extended:        false,
	}
}

// Return updates the loan's return date and status to indicate the book has been returned
func (l *Loan) Return(date time.Time) {
	// update returnDate
	l.returnDate = date

	// update loanStatus
	l.status = returned
}

// IsOverdue checks if the loan is overdue by comparing the current time with the due date.
// It returns true if the loan is still on loan and the due date has passed.
func (l *Loan) IsOverdue(now time.Time) bool {
	return l.status == onLoan && l.dueDate.Before(now)
}

// ExtendDueDate extends the due date of the loan by 7 days if it hasn't been extended before.
func (l *Loan) ExtendDueDate() error {
	if l.extended {
		return errors.New("dueDate has been already extended")
	}

	// extend
	l.dueDate = l.dueDate.AddDate(0, 0, 7)
	l.extended = true

	return nil
}
