package entity

import (
	"github.com/feualpha/debby/constant"
)

// StaticUser for non user debtor
type StaticUser struct {
	ID   int
	Name string
	Info string
}

// DebtorID for implement Debtor
func (sd StaticUser) DebtorID() int {
	return sd.ID
}

// DebtorType for implement Debtor
func (sd StaticUser) DebtorType() int {
	return constant.DebtorTypeStaticUser
}

// CreditorID for implement Creditor
func (sd StaticUser) CreditorID() int {
	return sd.ID
}

// CreditorType for implement Creditor
func (sd StaticUser) CreditorType() int {
	return constant.CreditorTypeStaticUser
}
