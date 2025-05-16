package http

import (
	"fmt"
	"strings"
)

type EmailRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Terms bool   `json:"terms"`
}
type Validatable interface {
	Validate() error
}

func (request *EmailRequest) Validate() error {
	if request.Name == "" || request.Email == "" || !request.Terms {
		return fmt.Errorf("required field")
	}
	if !strings.Contains(request.Email, "@") {
		return fmt.Errorf("wrong email reference")
	}
	return nil
}
