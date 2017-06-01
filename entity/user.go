package entity

import "time"

// user creditor debtor const
const (
	CreditorTypeUser = 1
	DebtorTypeUser   = 3
)

// User type, can be used as login and debtor
type User struct {
	ID        int
	Name      string
	Username  string
	Password  string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// DebtorID for implement Debtor
func (u User) DebtorID() int {
	return u.ID
}

// DebtorType for implement Debtor
func (u User) DebtorType() int {
	return DebtorTypeUser
}

// CreditorID for implement Debtor
func (u User) CreditorID() int {
	return u.ID
}

// CreditorType for implement Debtor
func (u User) CreditorType() int {
	return CreditorTypeUser
}
