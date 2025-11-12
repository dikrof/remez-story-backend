package errors

func ContainByCode(err error, errorCode ErrorCode) bool {
	if err == nil {
		return false
	}

	switch err := err.(type) {
	case *Error:
		return EqualByCode(err, errorCode)
	case *Errors:
		return err.ContainsByCode(errorCode)
	default:
		return false
	}
}

func EqualByCode(err error, errorCode ErrorCode) bool {
	if err == nil {
		return false
	}

	errValue, ok := err.(*Error)
	if !ok {
		return false
	}

	return errValue.Code() == errorCode
}

func CastOrWrap(err error, orErrorCode ErrorCode) *Error {
	if err == nil {
		return nil
	}

	switch err := err.(type) {
	case *Error:
		return err
	default:
		return NewError(orErrorCode, err.Error())
	}
}
