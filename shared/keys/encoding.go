package keys

// base(d) on eknkc/basex

import (
	"errors"
	"fmt"
	"slices"
)

var (
	alphabet = []rune("reblox0123456789acdfghijkmnpqstuvwyz")
	base     = len(alphabet)
	runeMap  = map[rune]int{
		'r': 0,
		'e': 1,
		'b': 2,
		'l': 3,
		'o': 4,
		'x': 5,
		'0': 6,
		'1': 7,
		'2': 8,
		'3': 9,
		'4': 10,
		'5': 11,
		'6': 12,
		'7': 13,
		'8': 14,
		'9': 15,
		'a': 16,
		'c': 17,
		'd': 18,
		'f': 19,
		'g': 20,
		'h': 21,
		'i': 22,
		'j': 23,
		'k': 24,
		'm': 25,
		'n': 26,
		'p': 27,
		'q': 28,
		's': 29,
		't': 30,
		'u': 31,
		'v': 32,
		'w': 33,
		'y': 34,
		'z': 35,
	}
)
var (
	stringLength = 50
	byteLength   = 32
)

// Encode takes a public/secret key and returns a string representation of itusing a custom alphabet
func Encode(source []byte) (string, error) {
	if len(source) != byteLength {
		// return "", fmt.Errorf("invalid length: %d", len(source))
		panic(fmt.Errorf("invalid length: %d", len(source)))
	}

	digits := make([]int, 1, stringLength)

	for i := range source {
		carry := int((source)[i])
		for j := range digits {
			carry += (digits)[j] << 8
			digits[j] = carry % base
			carry /= base
		}
		for carry > 0 {
			digits = append(digits, carry%base)
			carry /= base
		}
	}

	var res []rune

	for q := len(digits) - 1; q >= 0; q-- {
		res = append(res, alphabet[digits[q]])
	}

	// Pad with first characters to reach the ideal length
	for range stringLength - len(res) {
		res = append([]rune{alphabet[0]}, res...)
	}

	return string(res), nil
}

// Decode takes a string representation of a public/secret key and returns the byte representation of it
func Decode(source string) ([]byte, error) {
	if len(source) != stringLength {
		// return nil, fmt.Errorf("invalid length: %d", len(source))
		panic(fmt.Errorf("invalid length: %d", len(source)))
	}

	runes := []rune(source)

	bytes := []byte{0}
	for i := range runes {
		value, ok := runeMap[runes[i]]

		if !ok {
			return nil, errors.New("non-base character")
		}

		carry := int(value)

		for j := range bytes {
			carry += int(bytes[j]) * base
			bytes[j] = byte(carry & 0xff)
			carry >>= 8
		}

		for carry > 0 {
			bytes = append(bytes, byte(carry&0xff))
			carry >>= 8
		}
	}

	// Pad with zeros to reach the ideal length
	for range byteLength - len(bytes) {
		bytes = append(bytes, 0)
	}

	slices.Reverse(bytes)

	return bytes, nil
}
