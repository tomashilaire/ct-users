package validators

import "strings"

type Validators interface {
	ValidateSingUp(name string, email string, password string,
		confirmPassword string, userType string) error
	NormalizeEmail(email string) string
}

type validators struct{}

func NewValidators() Validators {
	return &validators{}
}

func (v *validators) ValidateSingUp(name string, email string, password string,
	confirmPassword string, userType string) error {
	return nil
}

func (v *validators) NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
