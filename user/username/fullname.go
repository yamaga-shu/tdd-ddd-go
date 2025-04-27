package username

type FullName struct {
	firstName string
	lastName  string
}

func NewFullName(firstName, lastName string) *FullName {
	return &FullName{
		firstName: firstName,
		lastName:  lastName,
	}
}

// Equals compares FullName fields
func (fn FullName) Equals(other *FullName) bool {
	return fn.firstName == other.firstName && fn.lastName == other.lastName
}
