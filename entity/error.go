package entity

type entityError struct {
	code    int
	message string
}

func (er entityError) Code() string {
	return string(er.code)
}

func (er entityError) Error() string {
	return er.message
}

// TODO i18n
var (
	InvalidDebtState = entityError{
		code:    10001,
		message: "invalid debt's state",
	}
	InvalidDebtStateTransition = entityError{
		code:    10002,
		message: "invalid debt's state transition",
	}
)
