package enum

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
)

// enum/acception_status.go
var IntToStatus = map[int64]AcceptionStatus{
	0: UNDER_REVIEW,
	1: IN_PROGRESS,
	2: COMPLETED,
	3: REJECTED,
	//4: RETURNED,
}

type AcceptionStatus string

const (
	UNDER_REVIEW AcceptionStatus = "UNDER_REVIEW"
	IN_PROGRESS  AcceptionStatus = "IN_PROGRESS"
	COMPLETED    AcceptionStatus = "COMPLETED"
	REJECTED     AcceptionStatus = "REJECTED"
	//RETURNED     AcceptionStatus = "RETURNED"
)

var statusToInt = map[AcceptionStatus]int64{
	UNDER_REVIEW: 0,
	IN_PROGRESS:  1,
	COMPLETED:    2,
	REJECTED:     3,
	//RETURNED:     4,
}

var intToStatus = map[int64]AcceptionStatus{
	0: UNDER_REVIEW,
	1: IN_PROGRESS,
	2: COMPLETED,
	3: REJECTED,
	//4: RETURNED,
}

var strToAcceptionStatus = map[string]AcceptionStatus{
	"UNDER_REVIEW": UNDER_REVIEW,
	"IN_PROGRESS":  IN_PROGRESS,
	"COMPLETED":    COMPLETED,
	"REJECTED":     REJECTED,
	//"RETURNED":     RETURNED,
}

func (s AcceptionStatus) Value() (driver.Value, error) {
	v, ok := statusToInt[s]
	if !ok {
		return nil, fmt.Errorf("invalid AcceptionStatus for DB: %s", s)
	}
	return v, nil
}
func (s AcceptionStatus) IsValid() bool {
	_, ok := statusToInt[s]
	return ok
}

func (s AcceptionStatus) IntValue() int64 {
	if !s.IsValid() {
		return -1 // or some other sentinel value
	}
	return statusToInt[s]
}

func (s *AcceptionStatus) Scan(value interface{}) error {
	switch v := value.(type) {
	case int64:
		if status, ok := intToStatus[v]; ok {
			*s = status
			return nil
		}
		return fmt.Errorf("invalid AcceptionStatus int64: %d", v)
	case []byte:
		return s.fromString(string(v))
	case string:
		return s.fromString(v)
	default:
		return fmt.Errorf("unsupported type for AcceptionStatus: %T", value)
	}
}

func (s *AcceptionStatus) fromString(strVal string) error {
	strVal = strings.ToUpper(strings.TrimSpace(strVal))

	// Handle number-as-string like "0"
	if intVal, err := strconv.Atoi(strVal); err == nil {
		if status, ok := intToStatus[int64(intVal)]; ok {
			*s = status
			return nil
		}
	}

	if status, ok := strToAcceptionStatus[strVal]; ok {
		*s = status
		return nil
	}

	return fmt.Errorf("invalid AcceptionStatus string: %s", strVal)
}

func (s AcceptionStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(string(s))
}

func (s *AcceptionStatus) UnmarshalJSON(data []byte) error {
	var str string
	if err := json.Unmarshal(data, &str); err != nil {
		return err
	}
	return s.fromString(str)
}
