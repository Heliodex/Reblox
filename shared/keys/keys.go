package keys

import (
	"fmt"
	"strings"

	"github.com/ChainSafe/go-schnorrkel"
)

func SecretKey(sk []byte) (*schnorrkel.SecretKey, error) {
	sec := schnorrkel.NewSecretKey([32]byte(sk), [32]byte{})

	return sec, nil
}

func PublicKey(skBytes []byte) ([]byte, error) {
	sk := schnorrkel.NewSecretKey([32]byte(skBytes), [32]byte{})

	pk, err := sk.Public()
	if err != nil {
		return nil, err
	}

	pkBytes := pk.Encode()

	return pkBytes[:], nil
}

type KeyType uint8

const (
	Invalid KeyType = iota
	Public
	Secret
)

func FormatKey(keyType KeyType, key string) (string, error) {
	if len(key) != stringLength {
		return "", fmt.Errorf("invalid length: %d", len(key))
	}

	// 5 sections of 10 characters each
	var formatted string

	if keyType == Public {
		formatted = "repub:"
	} else if keyType == Secret {
		formatted = "resec:"
	} else {
		return "", fmt.Errorf("invalid key type: %d", keyType)
	}

	for i := range 5 {
		formatted += key[i*10 : (i+1)*10]
		if i < 4 {
			formatted += "-"
		}
	}

	return formatted, nil
}

func UnformatKey(key string) (KeyType, string, error) {
	expectedLength := stringLength + 10
	if len(key) != expectedLength {
		return 0, "", fmt.Errorf("invalid length: expected %d, got %d", expectedLength, len(key))
	}

	var keyType KeyType
	firstPart := key[:6]

	if firstPart == "repub:" {
		keyType = Public
	} else if firstPart == "resec:" {
		keyType = Secret
	} else {
		return Invalid, "", fmt.Errorf("invalid key type: %s", key[:6])
	}

	withDashes := key[6:]
	withoutDashes := strings.ReplaceAll(withDashes, "-", "")

	if len(withoutDashes) != stringLength {
		return Invalid, "", fmt.Errorf("invalid length of unformatted key: %d", len(withoutDashes))
	}

	return keyType, withoutDashes, nil
}

func EncodeFormatted(keyType KeyType, source []byte) (string, error) {
	encoded, err := Encode(source)
	if err != nil {
		return "", err
	}

	return FormatKey(keyType, encoded)
}

func DecodeFormatted(key string) ([]byte, error) {
	_, unformatted, err := UnformatKey(key)
	if err != nil {
		return nil, err
	}

	return Decode(unformatted)
}
