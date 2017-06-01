package entity

import "testing"

func TestStatePendingPaidTransition(t *testing.T) {
	deb := Debt{State: PendingDebtState}
	err := deb.ChangeStateTo(PaidDebtState)
	if err != nil {
		t.Errorf("expect no error, got %s", err.Error())
	}

	if deb.State != PaidDebtState {
		t.Errorf("expect state %d, but got %d", PaidDebtState, deb.State)
	}
}

func TestStatePendingRejectTranstion(t *testing.T) {
	deb := Debt{State: PendingDebtState}
	err := deb.ChangeStateTo(RejectDebtState)
	if err != nil {
		t.Errorf("expect no error, got %s", err.Error())
	}

	if deb.State != RejectDebtState {
		t.Errorf("expect state %d, but got %d", RejectDebtState, deb.State)
	}
}

func TestStatePendingPendingTransition(t *testing.T) {
	deb := Debt{State: PendingDebtState}
	err := deb.ChangeStateTo(PendingDebtState)
	if err != InvalidDebtStateTransition {
		t.Errorf("expecting error \"%s\"", InvalidDebtStateTransition.Error())
	}

	if deb.State != PendingDebtState {
		t.Errorf("expect state %d, but got %d", PaidDebtState, deb.State)
	}
}

func TestStatePaidPendingTransition(t *testing.T) {
	deb := Debt{State: PaidDebtState}
	err := deb.ChangeStateTo(PendingDebtState)
	if err != InvalidDebtStateTransition {
		t.Errorf("expecting error \"%s\"", InvalidDebtStateTransition.Error())
	}

	if deb.State != PaidDebtState {
		t.Errorf("expect state %d, but got %d", PaidDebtState, deb.State)
	}
}

func TestStatePaidRejectTransition(t *testing.T) {
	deb := Debt{State: PaidDebtState}
	err := deb.ChangeStateTo(RejectDebtState)
	if err != InvalidDebtStateTransition {
		t.Errorf("expecting error \"%s\"", InvalidDebtStateTransition.Error())
	}

	if deb.State != PaidDebtState {
		t.Errorf("expect state %d, but got %d", PaidDebtState, deb.State)
	}
}

func TestStatePaidPaidTransition(t *testing.T) {
	deb := Debt{State: PaidDebtState}
	err := deb.ChangeStateTo(PaidDebtState)
	if err != InvalidDebtStateTransition {
		t.Errorf("expecting error \"%s\"", InvalidDebtStateTransition.Error())
	}

	if deb.State != PaidDebtState {
		t.Errorf("expect state %d, but got %d", PaidDebtState, deb.State)
	}
}

func TestStateRejectPendingTransition(t *testing.T) {
	deb := Debt{State: RejectDebtState}
	err := deb.ChangeStateTo(PendingDebtState)
	if err != InvalidDebtStateTransition {
		t.Errorf("expecting error \"%s\"", InvalidDebtStateTransition.Error())
	}

	if deb.State != RejectDebtState {
		t.Errorf("expect state %d, but got %d", RejectDebtState, deb.State)
	}
}

func TestStateRejectPaidTransition(t *testing.T) {
	deb := Debt{State: RejectDebtState}
	err := deb.ChangeStateTo(PaidDebtState)
	if err != InvalidDebtStateTransition {
		t.Errorf("expecting error \"%s\"", InvalidDebtStateTransition.Error())
	}

	if deb.State != RejectDebtState {
		t.Errorf("expect state %d, but got %d", RejectDebtState, deb.State)
	}
}

func TestStateRejectRejectTransition(t *testing.T) {
	deb := Debt{State: RejectDebtState}
	err := deb.ChangeStateTo(RejectDebtState)
	if err != InvalidDebtStateTransition {
		t.Errorf("expecting error \"%s\"", InvalidDebtStateTransition.Error())
	}

	if deb.State != RejectDebtState {
		t.Errorf("expect state %d, but got %d", RejectDebtState, deb.State)
	}
}

func TestInvalidState(t *testing.T) {
	defer func() {
		r := recover()
		if r == nil {
			t.Errorf("The code did not panic")
		}

		if r != InvalidDebtState {
			t.Errorf("expecting \"%s\" got \"%s\"", InvalidDebtState.Error(), r)
		}
	}()

	debtBeforeStateRules(4)
}
