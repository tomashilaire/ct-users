package validators

import "strings"

type Validators interface {
	ValidateSingUp()
	NormalizeEmail(email string) string
}

type validators struct{}

func NewValidators() Validators {
	return &validators{}
}

func (v *validators) ValidateSingUp() {

}

func (v *validators) NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
