package entity

import (
	"math/rand"
	"testing"

	"github.com/feualpha/debby/constant"
)

func TestStatiUserType(t *testing.T) {
	id := rand.Int()
	su := StaticUser{ID: id}

	if su.CreditorType() != constant.CreditorTypeStaticUser {
		t.Errorf("Expected %d got %d", constant.CreditorTypeStaticUser, su.CreditorType())
	}

	if su.DebtorType() != constant.DebtorTypeStaticUser {
		t.Errorf("Expected %d got %d", constant.DebtorTypeStaticUser, su.DebtorType())
	}

	if su.CreditorID() != id {
		t.Errorf("Expected %d got %d", id, su.CreditorID())
	}

	if su.DebtorID() != id {
		t.Errorf("Expected %d got %d", id, su.DebtorID())
	}
}
