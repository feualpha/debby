package entity

import (
	"github.com/feualpha/debby/constant"
)

// User type, can be used as login and debtor
type User struct {
	ID       int
	Name     string
	Username string
	Password string
	Email    string
}

// DebtorID for implement Debtor
func (u User) DebtorID() int {
	return u.ID
}

// DebtorType for implement Debtor
func (u User) DebtorType() int {
	return constant.DebtorTypeUser
}

// CreditorID for implement Debtor
func (u User) CreditorID() int {
	return u.ID
}

// CreditorType for implement Debtor
func (u User) CreditorType() int {
	return constant.CreditorTypeUser
}
