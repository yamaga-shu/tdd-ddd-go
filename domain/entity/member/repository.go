package member

import "github.com/google/uuid"

type IRepository interface {
	GetMemberById(memberId uuid.UUID) (*Member, error)
}
