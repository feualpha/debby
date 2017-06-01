package entity

import (
	"math/rand"
	"testing"
)

func TestUserType(t *testing.T) {
	id := rand.Int()
	u := User{ID: id}

	if u.CreditorType() != CreditorTypeUser {
		t.Errorf("Expected %d got %d", CreditorTypeUser, u.CreditorType())
	}

	if u.DebtorType() != DebtorTypeUser {
		t.Errorf("Expected %d got %d", DebtorTypeUser, u.DebtorType())
	}

	if u.CreditorID() != id {
		t.Errorf("Expected %d got %d", id, u.CreditorID())
	}

	if u.DebtorID() != id {
		t.Errorf("Expected %d got %d", id, u.DebtorID())
	}
}
