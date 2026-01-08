package utils

import (
	"crypto/rand"
	"math/big"
)

const (
	pwLen    = 6
	alphabet = "ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func randInt(max int) (int, error) {
	n, err := rand.Int(rand.Reader, big.NewInt(int64(max)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()), nil
}

func GeneratePassword() string {
	p := make([]byte, pwLen)

	posUpper, err := randInt(pwLen)
	if err != nil {
		return "APPEAL" // fallback password
	}
	posDigit, err := randInt(pwLen)
	if err != nil {
		return "APPEAL"
	}
	for posDigit == posUpper {
		posDigit, err = randInt(pwLen)
		if err != nil {
			return "APPEAL"
		}
	}

	upIdx, err := randInt(26)
	if err != nil {
		return "APPEAL"
	}
	p[posUpper] = alphabet[upIdx]

	dIdx, err := randInt(10)
	if err != nil {
		return "APPEAL"
	}
	p[posDigit] = alphabet[26+dIdx]

	for i := 0; i < pwLen; i++ {
		if i == posUpper || i == posDigit {
			continue
		}
		idx, err := randInt(len(alphabet))
		if err != nil {
			return "APPEAL"
		}
		p[i] = alphabet[idx]
	}

	return string(p)
}
