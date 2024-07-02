package urlErrors

import (
	"encoding/json"
	"fmt"
)

type GenericError struct {
	StatusCode int   `json:"statusCode"`
	Err        error `json:"-"`
}

func (r *GenericError) Error() string {
	return fmt.Sprintf("status %d: err %v", r.StatusCode, r.Err)
}

func (r *GenericError) MarshalJSON() ([]byte, error) {
	type Alias GenericError
	return json.Marshal(&struct {
		*Alias
		ErrorMsg string `json:"error"`
	}{
		Alias:    (*Alias)(r),
		ErrorMsg: r.Err.Error(),
	})
}
