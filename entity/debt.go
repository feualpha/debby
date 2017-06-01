package entity

// Debtor is interface
type Debtor interface {
	DebtorType() int
	DebtorID() int
}

// Creditor interface
type Creditor interface {
	CreditorType() int
	CreditorID() int
}

// Debt type Entity
type Debt struct {
	ID            int
	ResourceOwner User
	Creditor      User
	Debtor        Debtor
	State         int
	Amount        int
}
