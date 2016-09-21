package errors

import "fmt"

type APIError struct {
	Status       int
	ResponseBody []byte
}

func (e *APIError) Error() string {
	if len(e.ResponseBody) > 0 {
		return fmt.Sprintf("Error returned from API: [%d] %s", e.Status, e.ResponseBody)
	} else {
		return fmt.Sprintf("Error returned from API: [%d]", e.Status)
	}
}
