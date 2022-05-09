package utils

import (
	"errors"
	"time"
)

type TimeRequiredValidator struct{}

func (o *TimeRequiredValidator) Validate(value interface{}) error {
	t, ok := value.(time.Time)
	if !ok {
		return errors.New("value is not a Time")
	}
	if t.IsZero() {
		return errors.New("cannot be blank")
	}
	return nil
}

func TimeRequired() *TimeRequiredValidator {
	return &TimeRequiredValidator{}
}
