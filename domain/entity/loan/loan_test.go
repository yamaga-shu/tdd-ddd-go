package loan

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/yamaga-shu/tdd-ddd-go/domain/entity/member"
)

// tMemberRepository is Mocked MemberRepository for this test.
type tMemberRepository struct{}

func (tmr tMemberRepository) GetMemberById(memberId uuid.UUID) (*member.Member, error) {
	mem := member.New(memberId, "Mocked Name")

	return mem, nil
}

// tLoanRepository is Mocked LoanRepository for this test.
type tLoanRepository struct{}

func (tlr tLoanRepository) Create(l *Loan) error {
	return nil
}

func initLoan() *Loan {
	loanId, _ := uuid.NewV7()
	memberId, _ := uuid.NewV7()
	bookInventoryId, _ := uuid.NewV7()

	// 貸出日は 2025/05/01 に設定
	loanDate := time.Date(2025, time.May, 1, 0, 0, 0, 0, time.UTC)

	return New(
		loanId,
		memberId,
		bookInventoryId,
		loanDate,
		tLoanRepository{},
		tMemberRepository{},
	)
}

// TestReturnDate verifies the Return method
// It checks if the passed return date is valid or invalid
func TestReturnDate(t *testing.T) {
	// 2025/05/08 に返却（貸出日の後）
	loan := initLoan()
	validDate := time.Date(2025, time.May, 8, 0, 0, 0, 0, time.UTC)
	if err := loan.Return(validDate); err != nil {
		t.Errorf("Expected successful return, but got error: %v", err)
	}

	// 2025/04/10 に返却（貸出日の前）
	loan = initLoan()
	invalidDate := time.Date(2025, time.April, 10, 0, 0, 0, 0, time.UTC)
	if err := loan.Return(invalidDate); err == nil {
		t.Errorf("Expected error for return date before loan date, but got none")
	}
}

// TestReturnSideEffectDate verifies the side effects of the Return method on the returnDate field.
// It checks if the returnDate is correctly updated after a valid return and ensures that an error
// is returned when attempting to return the loan a second time.
func TestReturnSideEffectDate(t *testing.T) {
	loan := initLoan()
	validDate := time.Date(2025, time.May, 8, 0, 0, 0, 0, time.UTC)

	// Test the first Return call
	loan.Return(validDate)

	if loan.returnDate != validDate {
		t.Errorf("Expected returnDate was updated")
	}

	// Test duplicated Return call
	err := loan.Return(validDate)
	if err == nil {
		t.Errorf("Expected error for duplicated Return call, but got none")
	}
}

// TestReturnSideEffectStatus verifies the side effects of the Return method on the status field.
// It checks the initial status of the loan, ensures the status is updated to 'returned' after a valid
// return, and verifies that the status is not incorrectly set.
func TestReturnSideEffectStatus(t *testing.T) {
	loan := initLoan()
	validDate := time.Date(2025, time.May, 8, 0, 0, 0, 0, time.UTC)

	// Check initial status
	if loan.status != onLoan {
		t.Errorf("Expected status is onLoan")
	}

	// Test status update after Return
	loan.Return(validDate)

	if loan.status != returned {
		t.Errorf("Expected status is returned")
	}
}
