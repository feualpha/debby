package contract

type Error interface {
	Error() string
	Code() string
}
