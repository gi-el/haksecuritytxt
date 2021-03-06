package securitytxt

import (
	"fmt"
)

const (
	separatorErrorMsg = "no ':' found to separate field-name and value, as per section 3.6.8 of RFC5322"
	fieldNameErrorMsg = "field-name should be printable US-ASCII except space and :, as per section 3.6.8 of RFC5322"
	valueErrorMsg = "value should be printable US-ASCII except space, as per 'unstructured' syntax in section 3.2.5 of RFC5322"
	unknownFieldErrorMsg = "unexpected field-name '%s', as per section 3.5 of draft-foudil-securitytxt-09"
	emptyValueErrorMsg = "value cannot be empty"
	multipleValueErrorMsg = "multiple values for field '%s', expecting one value as per section 3.5 of draft-foudil-securitytxt-09"
	invalidTimeErrorMsg = "invalid time in field '%s' according to section 3.3 of RFC5322: %s"
	acknowledgmentsErrorMsg = "invalid field name 'acknowledgements', should be 'acknowledgments' as per section 3.5.1 of draft-foudil-securitytxt-09"
	contentTypeErrorMsg = "invalid Content-Type of '%s', expecting 'text/plain; charset=utf-8' as per section 3 of draft-foudil-securitytxt-09"
)

type SyntaxError struct {
	lineNo int
	line string
	msg string
}

func (e SyntaxError) Error() string {
	return fmt.Sprintf("Error in line %d: %s", e.lineNo, e.msg)
}

type HTTPError struct {
	err error
}

func (e HTTPError) Error() string {
	return e.err.Error()
}

func (e HTTPError) Unwrap() error {
	return e.err
}
