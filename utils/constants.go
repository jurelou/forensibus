package utils

const (
	Success = iota
	Failure
)

func IsErrored(status uint32) bool {
	return status > Success
}
