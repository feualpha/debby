package entity

import (
	"time"

	"github.com/feualpha/debby/contract"
)

// Debt State
const (
	PendingDebtState = 1
	PaidDebtState    = 2
	RejectDebtState  = 3
)

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
	Creditor      Creditor
	Debtor        Debtor
	State         int
	Amount        int
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

// ChangeStateTo is Debt State Machine
func (d *Debt) ChangeStateTo(targetState int) contract.Error {
	validBefore := debtBeforeStateRules(targetState)
	validFlag := false
	for i := range validBefore {
		if validBefore[i] == d.State {
			validFlag = true
			break
		}
	}

	if validFlag {
		d.State = targetState
		return nil
	}

	return InvalidDebtStateTransition
}

func debtBeforeStateRules(state int) []int {
	switch state {
	case PendingDebtState:
		return []int{}
	case PaidDebtState:
		return []int{PendingDebtState}
	case RejectDebtState:
		return []int{PendingDebtState}
	default:
		panic(InvalidDebtState)
	}
}
