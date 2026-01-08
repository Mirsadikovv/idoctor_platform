package enum

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

type ActivityStatus string

const (
	ACTIVE   ActivityStatus = "ACTIVE"
	INACTIVE ActivityStatus = "INACTIVE"
)

func (s ActivityStatus) String() string {
	return string(s)
}

func (s ActivityStatus) Bool() bool {
	return s == ACTIVE
}

func (s ActivityStatus) IsValid() bool {
	return s == ACTIVE || s == INACTIVE
}

func (s ActivityStatus) MarshalJSON() ([]byte, error) {
	if !s.IsValid() {
		return nil, fmt.Errorf("invalid ActivityStatus: %s", s)
	}
	return json.Marshal(string(s))
}

func (s *ActivityStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	str = strings.ToUpper(strings.TrimSpace(str))
	switch str {
	case "ACTIVE":
		*s = ACTIVE
	case "INACTIVE":
		*s = INACTIVE
	default:
		return fmt.Errorf("invalid ActivityStatus: %s", str)
	}
	return nil
}

func (s ActivityStatus) Value() (driver.Value, error) {
	if !s.IsValid() {
		return nil, errors.New("invalid ActivityStatus for DB")
	}
	return s == ACTIVE, nil
}

func (s *ActivityStatus) Scan(value any) error {
	switch v := value.(type) {
	case bool:
		if v {
			*s = ACTIVE
		} else {
			*s = INACTIVE
		}
	default:
		return fmt.Errorf("unsupported type for --- ActivityStatus: %T", value)
	}
	return nil
}
