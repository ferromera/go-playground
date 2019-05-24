package errors

type ApiError struct {
	Message  string `json:"message"`
	ErrorStr string `json:"error"`
	Status   int    `json:"status"`
	Cause    string `json:"cause"`
}

func (e ApiError) Error() string {
	return e.Message
}
