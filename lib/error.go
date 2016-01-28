package lib

type MyError struct {
	msg string
}

func (err MyError) Error() string {
	return err.msg
}
