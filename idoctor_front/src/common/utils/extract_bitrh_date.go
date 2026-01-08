package utils

import (
	"fmt"
)

func ExtractBirthDateFromPINFL[T ~string | ~int64](pinfl T) (string, error) {
	var pinflStr = fmt.Sprintf("%v", pinfl)
	if len(pinflStr) != 14 {
		return "", fmt.Errorf("invalid PINFL length")
	}

	day := pinflStr[1:3]
	month := pinflStr[3:5]
	year := pinflStr[5:7]

	century := pinflStr[0]
	switch century {
	case '1', '2':
		year = "18" + year
	case '3', '4':
		year = "19" + year
	case '5', '6':
		year = "20" + year
	case '7', '8':
		year = "21" + year
	default:
		return "", fmt.Errorf("invalid century digit in PINFL")
	}

	date := fmt.Sprintf("%s-%s-%s", year, month, day)

	return date, nil
}
