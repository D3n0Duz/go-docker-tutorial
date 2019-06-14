package models


// New returns an error that formats as the given text.
func New(text string, code int) Errors {
	return Errors{text, code}
}

// errorString is a trivial implementation of error.
type Errors struct {
	S string
	C int
}

func (e *Errors) Error() string {
	return e.S
}
func (e *Errors) Code() int {
	return e.C
}