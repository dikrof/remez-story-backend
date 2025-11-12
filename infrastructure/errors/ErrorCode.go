package errors

type ErrorCode string

func (c ErrorCode) String() string {
	return string(c)
}
