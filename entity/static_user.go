package entity

// static user debtor and creditor type constant
const (
	CreditorTypeStaticUser = 2
	DebtorTypeStaticUser   = 4
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
	return DebtorTypeStaticUser
}

// CreditorID for implement Creditor
func (sd StaticUser) CreditorID() int {
	return sd.ID
}

// CreditorType for implement Creditor
func (sd StaticUser) CreditorType() int {
	return CreditorTypeStaticUser
}
