package exception

type NotFoundError struct {
	Error func() string
}

func NewNotFoundError(err error) NotFoundError {
	return NotFoundError{
		Error: func() string { return err.Error() },
	}
}
