package urlErrors

import (
	"encoding/json"
	"errors"
	"fmt"
)

type RedirectError struct {
	GenericError
	UrlExternalId string `json:"urlExternalId"`
}

func (r *RedirectError) Error() string {
	return fmt.Sprintf("status %d: err %v, urlExternalId:%s", r.StatusCode, r.Err, r.UrlExternalId)
}

func (r *RedirectError) MarshalJSON() ([]byte, error) {
	type Alias RedirectError
	return json.Marshal(&struct {
		*Alias
		ErrorMsg string `json:"error"`
	}{
		Alias:    (*Alias)(r),
		ErrorMsg: r.Err.Error(),
	})
}

func FromExternalID(externalId string) *RedirectError {
	return &RedirectError{GenericError: GenericError{StatusCode: 400, Err: errors.New("something went wrong")}, UrlExternalId: externalId}
}

func FromExternalIDWithCustomMessageAndCode(externalId string, message string, statusCode int) *RedirectError {
	return &RedirectError{GenericError: GenericError{StatusCode: statusCode, Err: errors.New(message)}, UrlExternalId: externalId}
}
