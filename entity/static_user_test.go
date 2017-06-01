package entity

import (
	"math/rand"
	"testing"
)

func TestStatiUserType(t *testing.T) {
	id := rand.Int()
	su := StaticUser{ID: id}

	if su.CreditorType() != CreditorTypeStaticUser {
		t.Errorf("Expected %d got %d", CreditorTypeStaticUser, su.CreditorType())
	}

	if su.DebtorType() != DebtorTypeStaticUser {
		t.Errorf("Expected %d got %d", DebtorTypeStaticUser, su.DebtorType())
	}

	if su.CreditorID() != id {
		t.Errorf("Expected %d got %d", id, su.CreditorID())
	}

	if su.DebtorID() != id {
		t.Errorf("Expected %d got %d", id, su.DebtorID())
	}
}
