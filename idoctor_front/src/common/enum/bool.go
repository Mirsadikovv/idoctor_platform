package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type Bool string

const (
	TRUE  Bool = "TRUE"
	FALSE Bool = "FALSE"
)

// String returns "TRUE" or "FALSE"
func (b Bool) String() string {
	return string(b)
}

// Bool converts to native Go bool
func (b Bool) Bool() bool {
	return b == TRUE
}

// IsValid checks for allowed values
func (b Bool) IsValid() bool {
	return b == TRUE || b == FALSE
}

// MarshalJSON outputs "TRUE" or "FALSE"
func (b Bool) MarshalJSON() ([]byte, error) {
	if !b.IsValid() {
		return nil, fmt.Errorf("invalid Bool value: %s", b)
	}
	return json.Marshal(string(b))
}

// UnmarshalJSON accepts only "TRUE" or "FALSE" (case-insensitive)
func (b *Bool) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	str = strings.ToUpper(strings.TrimSpace(str))
	switch str {
	case "TRUE":
		*b = TRUE
	case "FALSE":
		*b = FALSE
	default:
		return fmt.Errorf("invalid Bool string: %s", str)
	}
	return nil
}

// Value returns true or false for database (GORM)
func (b Bool) Value() (driver.Value, error) {
	if !b.IsValid() {
		return nil, errors.New("invalid Bool for database")
	}
	return b == TRUE, nil
}

// Scan reads true/false from DB and maps to "TRUE"/"FALSE"
func (b *Bool) Scan(value any) error {
	switch v := value.(type) {
	case bool:
		if v {
			*b = TRUE
		} else {
			*b = FALSE
		}
	default:
		return fmt.Errorf("unsupported type for Bool: %T", value)
	}
	return nil
}

// UnmarshalText handles form data binding - accepts "true"/"false" or "TRUE"/"FALSE"
func (b *Bool) UnmarshalText(text []byte) error {
	str := strings.ToUpper(strings.TrimSpace(string(text)))
	switch str {
	case "TRUE", "1":
		*b = TRUE
	case "FALSE", "0", "":
		*b = FALSE
	default:
		return fmt.Errorf("invalid Bool text: %s", string(text))
	}
	return nil
}
