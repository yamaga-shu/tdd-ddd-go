package member

import "github.com/google/uuid"

type IRepository interface {
	GetMemberById(memberId uuid.UUID) (*Member, error)
}

// MRepository is Mocked Respository of member
type MRepository struct{}

func (mr MRepository) GetMemberById(memberId uuid.UUID) (*Member, error) {
	mem := New(memberId, "Mocked Name")

	return mem, nil
}
