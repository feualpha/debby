package entity

import (
	"math/rand"
	"testing"

	"github.com/feualpha/debby/constant"
)

func TestUserType(t *testing.T) {
	id := rand.Int()
	u := User{ID: id}

	if u.CreditorType() != constant.CreditorTypeUser {
		t.Errorf("Expected %d got %d", constant.CreditorTypeUser, u.CreditorType())
	}

	if u.DebtorType() != constant.DebtorTypeUser {
		t.Errorf("Expected %d got %d", constant.DebtorTypeUser, u.DebtorType())
	}

	if u.CreditorID() != id {
		t.Errorf("Expected %d got %d", id, u.CreditorID())
	}

	if u.DebtorID() != id {
		t.Errorf("Expected %d got %d", id, u.DebtorID())
	}
}
