package validator

import (
	"errors"
	"fmt"
	"regexp"
	"strconv"
)

// ValidatePINFL checks if PINFL is valid (14-digit Uzbekistan ID)
func ValidatePINFL(pinfl string) error {
	// must be exactly 14 digits
	match, _ := regexp.MatchString(`^\d{14}$`, pinfl)
	if !match {
		return errors.New("PINFL must be exactly 14 digits")
	}

	// convert digits to int array
	digits := make([]int, 14)
	for i := 0; i < 14; i++ {
		d, _ := strconv.Atoi(string(pinfl[i]))
		digits[i] = d
	}

	// Weights for checksum
	weights := []int{7, 3, 1, 7, 3, 1, 7, 3, 1, 7, 3, 1, 7}
	sum := 0
	for i := 0; i < 13; i++ {
		sum += digits[i] * weights[i]
	}
	controlDigit := sum % 10

	// last digit must equal control digit
	if digits[13] != controlDigit {
		return fmt.Errorf("invalid PINFL: checksum failed")
	}

	return nil
}
