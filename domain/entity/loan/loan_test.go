package loan

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

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
	)
}

// TestReturnDate verifies the Return method
// It checks if the passed return date is valid or invalid
func TestReturnDate(t *testing.T) {
	// 2025/05/08 に返却（貸出日の後）
	loan := initLoan()
	returnDate := time.Date(2025, time.May, 8, 0, 0, 0, 0, time.UTC)
	if err := loan.Return(returnDate); err != nil {
		t.Errorf("Expected successful return, but got error: %v", err)
	}

	// 2025/04/10 に返却（貸出日の前）
	loan = initLoan()
	returnDate = time.Date(2025, time.April, 10, 0, 0, 0, 0, time.UTC)
	if err := loan.Return(returnDate); err == nil {
		t.Errorf("Expected error for return date before loan date, but got none")
	}
}

func TestReturnSideEffect(t *testing.T) {
	loan := initLoan()

	// first Return
	loan.Return(time.Now())

	// duplicated Return
	err := loan.Return(time.Now())
	if err == nil {
		t.Errorf("Expected error for duplicated Return call, but got none")
	}
}
