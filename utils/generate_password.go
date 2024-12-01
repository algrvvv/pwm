package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"
)

const (
	lowerChars = "abcdefghijklmnopqrstuvwxyz"
	upperChars = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	digits     = "0123456789"
	specials   = "!@#$%^&*()-_=+[]{}<>?/|"
	allChars   = lowerChars + upperChars + digits + specials
)

var (
	ErrInvalidPassLen = errors.New("invalid password len")
	ErrInvalidChars   = errors.New("no characters available for password generation")
)

func GeneratePassword(l int, useUpper, useDigits, useSpecials bool) (string, error) {
	if l <= 0 {
		return "", ErrInvalidPassLen
	}

	charset := lowerChars
	if useUpper {
		charset += upperChars
	}
	if useDigits {
		charset += digits
	}
	if useSpecials {
		charset += specials
	}

	if len(charset) == 0 {
		return "", ErrInvalidChars
	}

	password := make([]byte, l)
	for i := range password {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", fmt.Errorf("failed to generate random number: %w", err)
		}

		password[i] = charset[index.Int64()]
	}

	return string(password), nil
}
