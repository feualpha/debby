package entity

import (
	"math/rand"
	"testing"
)

func TestEntityErrorAttr(t *testing.T) {
	code := rand.Int()
	message := "error message"
	ee := entityError{code: code, message: message}

	if ee.Code() != string(code) {
		t.Errorf("expected error code %s, but get %s", string(code), ee.Code())
	}

	if ee.Error() != message {
		t.Errorf("expected error message %s, but get %s", message, ee.Error())
	}
}
