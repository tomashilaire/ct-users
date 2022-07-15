package validators

import (
	"regexp"
	"root/pkg/apperrors"
	"root/pkg/errors"
	"strings"
)

type Validators interface {
	ValidateSingUp(name string, lastName string, email string, password string,
		confirmPassword string, userType string) error
	NormalizeEmail(email string) string
}

type validators struct{}

func NewValidators() Validators {
	return &validators{}
}

func (v *validators) ValidateSingUp(name string, lastName string, email string, password string,
	confirmPassword string, userType string) error {
	validatedName, err := regexp.Match(`\b[\w][^\d]+[a-zA-Z]\b`, []byte(name))
	if err != nil {
		return err
	}
	validateLastname, err := regexp.Match(`\b[\w][^\d]+[a-zA-Z]\b`, []byte(lastName))
	if err != nil {
		return err
	}
	validateUserType, err := regexp.Match(`^([s|S]ponsor|[p|P]artner|[a|A]dmin)$`, []byte(userType))
	if err != nil {
		return err
	}
	validateEmail, err := regexp.Match(`[\w.%+-]+@[\w.-]+\.[a-zA-Z]{2,4}`, []byte(email))
	if err != nil {
		return err
	}
	validatePassword, err := regexp.Match(`^[a-zA-Z\d_-]{8,16}$`, []byte(password))
	if err != nil {
		return err
	}
	if password != confirmPassword {
		return errors.LogError(errors.New(apperrors.InvalidInput, err, "The password is incorrect", ""))
	}
	if !(validatedName && validateLastname && validateUserType && validateEmail && validatePassword) {
		return errors.LogError(errors.New(apperrors.InvalidInput, err, "One of the data entered is wrong", ""))
	}

	return nil
}

func (v *validators) NormalizeEmail(email string) string {
	return strings.TrimSpace(strings.ToLower(email))
}
