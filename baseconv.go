// Converts a string in an arbitrary base to any other arbitrary base.
package baseconv

import (
	"errors"
	"fmt"
	"strings"
)

// Convert num from specified base to a different base.
func Convert(num, fromBase, toBase string) (string, error) {
	if num == "" {
		return "", errors.New("invalid number")
	}

	if len(fromBase) < 2 {
		return "", errors.New("invalid fromBase")
	}

	if len(toBase) < 2 {
		return "", errors.New("invalid toBase")
	}

	fromLen := len(fromBase)
	toLen := len(toBase)
	numLen := len(num)
	result := make([]byte, 0)

	number := make([]int, numLen)
	for i := 0; i < numLen; i++ {
		number[i] = strings.IndexByte(fromBase, num[i])
		if number[i] < 0 {
			return "", errors.New(fmt.Sprintf("invalid character '%c' at %d", num[i], i))
		}
	}

	// loop until whole number is converted
	for {
		divide := 0
		newlen := 0

		// perform division manually (which is why this works with big numbers)
		for i := 0; i < numLen; i++ {
			divide = divide*fromLen + number[i]
			if divide >= toLen {
				number[newlen] = int(divide / toLen)
				divide = divide % toLen
				newlen++
			} else if newlen > 0 {
				number[newlen] = 0
				newlen++
			}
		}

		numLen = newlen
		result = append([]byte{toBase[divide]}, result...) // divide is basically num % toLen (i.e. the new character)

		if newlen == 0 {
			break
		}
	}

	return string(result), nil
}

const (
	DigitsBin = "01"
	DigitsOct = "01234567"
	DigitsDec = "0123456789"
	DigitsHex = "0123456789abcdef"
	Digits36  = "0123456789abcdefghijklmnopqrstuvwxyz"
	Digits62  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits64  = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ-_"
)

// Encode a string into DigitsBin with optional specified base (default: DigitsDec).
func EncodeBin(num string, base ...string) (string, error) {
	from := DigitsDec
	if len(base) > 0 {
		from = base[0]
	}

	return Convert(num, from, DigitsBin)
}

// Decode a string from DigitsBin with optional specified base (default: DigitsDec).
func DecodeBin(num string, base ...string) (string, error) {
	to := DigitsDec
	if len(base) > 0 {
		to = base[0]
	}

	return Convert(num, DigitsBin, to)
}

// Encode a string into DigitsOct with optional specified base (default: DigitsDec).
func EncodeOct(num string, base ...string) (string, error) {
	from := DigitsDec
	if len(base) > 0 {
		from = base[0]
	}

	return Convert(num, from, DigitsOct)
}

// Decode a string from DigitsOct with optional specified base (default: DigitsDec).
func DecodeOct(num string, base ...string) (string, error) {
	to := DigitsDec
	if len(base) > 0 {
		to = base[0]
	}

	return Convert(num, DigitsOct, to)
}

// Encode a string into DigitsHex with optional specified base (default: DigitsDec).
func EncodeHex(num string, base ...string) (string, error) {
	from := DigitsDec
	if len(base) > 0 {
		from = base[0]
	}

	return Convert(num, from, DigitsHex)
}

// Decode a string from DigitsHex with optional specified base (default: DigitsDec).
func DecodeHex(num string, base ...string) (string, error) {
	to := DigitsDec
	if len(base) > 0 {
		to = base[0]
	}

	return Convert(num, DigitsHex, to)
}

// Encode a string into Digits36 with optional specified base (default: DigitsDec).
func Encode36(num string, base ...string) (string, error) {
	from := DigitsDec
	if len(base) > 0 {
		from = base[0]
	}

	return Convert(num, from, Digits36)
}

// Decode a string from Digits36 with optional specified base (default: DigitsDec).
func Decode36(num string, base ...string) (string, error) {
	to := DigitsDec
	if len(base) > 0 {
		to = base[0]
	}

	return Convert(num, Digits36, to)
}

// Encode a string into Digits62 with optional specified base (default: DigitsDec).
func Encode62(num string, base ...string) (string, error) {
	from := DigitsDec
	if len(base) > 0 {
		from = base[0]
	}

	return Convert(num, from, Digits62)
}

// Decode a string from Digits62 with optional specified base (default: DigitsDec).
func Decode62(num string, base ...string) (string, error) {
	to := DigitsDec
	if len(base) > 0 {
		to = base[0]
	}

	return Convert(num, Digits62, to)
}

// Encode a string into Digits64 with optional specified base (default: DigitsDec).
func Encode64(num string, base ...string) (string, error) {
	from := DigitsDec
	if len(base) > 0 {
		from = base[0]
	}

	return Convert(num, from, Digits64)
}

// Decode a string from Digits64 with optional specified base (default: DigitsDec).
func Decode64(num string, base ...string) (string, error) {
	to := DigitsDec
	if len(base) > 0 {
		to = base[0]
	}

	return Convert(num, Digits64, to)
}
