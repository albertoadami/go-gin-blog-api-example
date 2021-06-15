package errors

type ApplicationErrorStatus int

const (
	EmailAlreadyInUse = iota + 1
)

type EmailAlreadyInUseError struct {
	Status  ApplicationErrorStatus
	Message string
}

func (ee EmailAlreadyInUseError) Error() string {
	return ee.Message
}
