package utils

const (
	Success = iota
	Failure
	GrpcFailure
	Timeout
)

func IsErrored(status uint32) bool {
	return status > Success
}
