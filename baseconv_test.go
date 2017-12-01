package baseconv

import (
	"strings"
	"testing"
)

func TestErrors(t *testing.T) {
	tests := []struct {
		val, from, to string
	}{
		{"", DigitsHex, DigitsDec},
		{"0", "", DigitsDec},
		{"0", DigitsDec, ""},
		{"bad", DigitsBin, DigitsDec},
		{"BAD", DigitsHex, DigitsDec},
	}

	for i, test := range tests {
		_, err := Convert(test.val, test.from, test.to)
		if err == nil {
			t.Errorf("test %d Convert(%s, %s, %s) should produce error", i, test.val, test.from, test.to)
		}
	}
}

const (
	DigitsJapanese = `〇一二三四五六七八九`
	DigitsThai     = `๐๑๒๓๔๕๖๗๘๙`
	Digits96a      = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz~`!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?¿¡"
	Digits96b      = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ~`!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?¿¡"
)

func TestZeroes(t *testing.T) {
	z := strings.Repeat("0", 1000000)
	for i, base := range []string{DigitsBin, DigitsOct, DigitsDec, DigitsHex, Digits36, Digits62, Digits64, Digits96a, Digits96b} {
		v, err := Convert(z, base, DigitsDec)
		if err != nil {
			t.Fatalf("test %d base expected no error, got: %v", i, err)
		}
		if v != "0" {
			t.Errorf("test %d expected 0, got: %s", i, v)
		}
		v, err = Convert(z, DigitsDec, base)
		if err != nil {
			t.Fatalf("test %d base expected no error, got: %v", i, err)
		}
		if v != "0" {
			t.Errorf("test %d expected 0, got: %s", i, v)
		}
	}

	z += "1"
	for i, base := range []string{DigitsBin, DigitsOct, DigitsDec, DigitsHex, Digits36, Digits62, Digits64, Digits96a, Digits96b} {
		v, err := Convert(z, base, DigitsDec)
		if err != nil {
			t.Fatalf("test %d base expected no error, got: %v", i, err)
		}
		if v != "1" {
			t.Errorf("test %d expected 1, got: %s", i, v)
		}
		v, err = Convert(z, DigitsDec, base)
		if err != nil {
			t.Fatalf("test %d base expected no error, got: %v", i, err)
		}
		if v != "1" {
			t.Errorf("test %d expected 1, got: %s", i, v)
		}
	}
}

func TestConvert(t *testing.T) {
	tests := []struct {
		from, to, val, exp string
	}{
		{DigitsDec, DigitsBin, "0", "0"},
		{DigitsDec, DigitsBin, "8", "1000"},
		{DigitsDec, DigitsBin, "15", "1111"},
		{DigitsDec, DigitsBin, "16", "10000"},
		{DigitsDec, DigitsBin, "88", "1011000"},
		{DigitsDec, DigitsBin, "10000", "10011100010000"},

		{DigitsDec, DigitsHex, "0", "0"},
		{DigitsDec, DigitsHex, "8", "8"},
		{DigitsDec, DigitsHex, "15", "f"},
		{DigitsDec, DigitsHex, "16", "10"},
		{DigitsDec, DigitsHex, "88", "58"},
		{DigitsDec, DigitsHex, "10000", "2710"},

		{DigitsDec, Digits62, "16571982744576742462", "jKbR7u8J5PU"},
		{DigitsDec, Digits62, "46394851265279874948", "TheUtUU3miE"},
		{DigitsDec, Digits62, "21901407667833273510", "q5SG7U76tls"},
		{DigitsDec, Digits62, "8232087098322120342", "9O72RLP5fF4"},
		{DigitsDec, Digits62, "6354358749246709610", "7zp1TbLFPp8"},
		{DigitsDec, Digits62, "18089061068", "jKbR7u"},
		{DigitsDec, Digits62, "50642057182", "TheUtU"},
		{DigitsDec, Digits62, "23906366962", "q5SG7U"},
		{DigitsDec, Digits62, "8985691605", "9O72RL"},
		{DigitsDec, Digits62, "6936067049", "7zp1Tb"},
		{DigitsDec, Digits62, "799310853702667", "3EYjA0o7p"},

		{DigitsDec, Digits64, "20100203105211888256765428281344829", "ZY4eMQ2qFcP-xIh3UcZ"},
		{DigitsDec, Digits64, "20110423215600563210173308035411215", "Z-5ew8KnbFn70adF2Qf"},

		{DigitsHex, DigitsBin, "70b1d707eac2edf4c6389f440c7294b51fff57bb", "111000010110001110101110000011111101010110000101110110111110100110001100011100010011111010001000000110001110010100101001011010100011111111111110101011110111011"},
		{DigitsHex, DigitsBin, "8fc60e7c3b3c48e9a6a7a5fe4f1fbc31", "10001111110001100000111001111100001110110011110001001000111010011010011010100111101001011111111001001111000111111011110000110001"},

		{DigitsHex, Digits36, "abcdef00001234567890", "3o47re02jzqisvio"},
		{DigitsHex, Digits36, "abcdef01234567890123456789abcdef", "a65xa07491kf5zyfpvbo76g33"},

		{Digits62, DigitsHex, "cBaidlJ84Ggc5JA7IYCgv", "6ad547ffe02477b9473f7977e4d5e17"},
		{Digits62, DigitsHex, "4nipILgJlXPutO1hsisIJr", "8fc60e7c3b3c48e9a6a7a5fe4f1fbc31"},
		{Digits62, DigitsHex, "4vqyd6OoARXqj9nRUNhtLQ", "941532a06be1443aa9d5d57bdf180a52"},
		{Digits62, DigitsHex, "5FY8KwTsQaUJ2KzHJGetfE", "ba86b8f06fdf494487a08a491a19490e"},
		{Digits62, DigitsHex, "7N42dgm5tFLK9N8MT7fHC7", "ffffffffffffffffffffffffffffffff"},

		{DigitsDec, "Christopher", "355927353784509896715106760", "iihtspiphoeCrCeshhorsrrtrh"},

		// unicode
		{DigitsDec, DigitsJapanese, "9876543210", `九八七六五四三二一〇`},
		{DigitsDec, DigitsJapanese, "98765432100123456789", `九八七六五四三二一〇〇一二三四五六七八九`},

		{DigitsDec, DigitsThai, "9876543210", `๙๘๗๖๕๔๓๒๑๐`},

		{DigitsHex, DigitsJapanese, "2710", "一〇〇〇〇"},
		{DigitsHex, DigitsThai, "2710", "๑๐๐๐๐"},
		{DigitsHex, "0一23456789", "2710", "一0000"},

		{DigitsHex, Digits96a, `45fb65999e10b359f34d2c73fb4b1ae01f7ecfce23f5a9908f138ef1770d7b75`, "E<jk|A(~&a$@,Zzy:K_m,`n8~l@%|RO8)UDNn~r"},
		{strings.ToUpper(DigitsHex), Digits96b, `45FB65999E10B359F34D2C73FB4B1AE01F7ECFCE23F5A9908F138EF1770D7B75`, "e<JK|a(~&A$@,zZY:k_M,`N8~L@%|ro8)udnN~R"},
	}

	for i, test := range tests {
		v0, err := Convert(test.val, test.from, test.to)
		if err != nil {
			t.Fatalf("test %d expected no error, got: %v", i, err)
		}
		if test.exp != v0 {
			t.Errorf("test %d (%d->%d) expected %s, got: %s", i, len(test.from), len(test.to), test.exp, v0)
		}

		v1, err := Convert(test.exp, test.to, test.from)
		if err != nil {
			t.Fatal(err)
		}
		if test.val != v1 {
			t.Errorf("test %d (%d->%d) expected %s, got: %s", i, len(test.to), len(test.from), test.val, v1)
		}
	}
}

func TestEncodeDecode(t *testing.T) {
	v0 := "1627734050041231452076"

	var tests = []struct {
		encode func(string) (string, error)
		decode func(string) (string, error)
		exp    string
	}{
		{EncodeBinFromDec, DecodeBinToDec, "10110000011110101011001000001100110100001101011100000100010011110101100"},
		{EncodeOctFromDec, DecodeOctToDec, "260365310146415340423654"},
		{EncodeHexFromDec, DecodeHexToDec, "583d5906686b8227ac"},
		{Encode36FromDec, Decode36ToDec, "9jird8fbzkui7g"},
		{Encode62FromDec, Decode62ToDec, "vhozdwL3WC8A"},
		{Encode64FromDec, Decode64ToDec, "m3Rp1CxHwyuI"},
	}

	for i, test := range tests {
		v1, err := test.encode(v0)
		if err != nil {
			t.Fatal(err)
		}
		if test.exp != v1 {
			t.Errorf("test %d values %s / %s should match", i, test.exp, v1)
		}

		v2, err := test.decode(v1)
		if err != nil {
			t.Fatal(err)
		}
		if v0 != v2 {
			t.Errorf("test %d values %s / %s should match", i, v0, v2)
		}

		v3, err := test.encode(v0)
		if err != nil {
			t.Fatal(err)
		}
		if test.exp != v3 {
			t.Errorf("test %d values %s / %s should match", i, test.exp, v3)
		}

		v4, err := test.decode(v1)
		if err != nil {
			t.Fatal(err)
		}
		if v0 != v4 {
			t.Errorf("test %d values %s / %s should match", i, v0, v4)
		}
	}
}

func BenchmarkConvert(b *testing.B) {
	tests := []struct {
		from, to, val, exp string
	}{
		{DigitsDec, DigitsBin, "0", "0"},
		{DigitsDec, DigitsBin, "8", "1000"},
		{DigitsDec, DigitsBin, "15", "1111"},
		{DigitsDec, DigitsBin, "16", "10000"},
		{DigitsDec, DigitsBin, "88", "1011000"},
		{DigitsDec, DigitsBin, "10000", "10011100010000"},

		{DigitsDec, DigitsHex, "0", "0"},
		{DigitsDec, DigitsHex, "8", "8"},
		{DigitsDec, DigitsHex, "15", "f"},
		{DigitsDec, DigitsHex, "16", "10"},
		{DigitsDec, DigitsHex, "88", "58"},
		{DigitsDec, DigitsHex, "10000", "2710"},

		{DigitsDec, Digits62, "16571982744576742462", "jKbR7u8J5PU"},
		{DigitsDec, Digits62, "46394851265279874948", "TheUtUU3miE"},
		{DigitsDec, Digits62, "21901407667833273510", "q5SG7U76tls"},
		{DigitsDec, Digits62, "8232087098322120342", "9O72RLP5fF4"},
		{DigitsDec, Digits62, "6354358749246709610", "7zp1TbLFPp8"},
		{DigitsDec, Digits62, "18089061068", "jKbR7u"},
		{DigitsDec, Digits62, "50642057182", "TheUtU"},
		{DigitsDec, Digits62, "23906366962", "q5SG7U"},
		{DigitsDec, Digits62, "8985691605", "9O72RL"},
		{DigitsDec, Digits62, "6936067049", "7zp1Tb"},
		{DigitsDec, Digits62, "799310853702667", "3EYjA0o7p"},

		{DigitsDec, Digits64, "20100203105211888256765428281344829", "ZY4eMQ2qFcP-xIh3UcZ"},
		{DigitsDec, Digits64, "20110423215600563210173308035411215", "Z-5ew8KnbFn70adF2Qf"},

		{DigitsHex, DigitsBin, "70b1d707eac2edf4c6389f440c7294b51fff57bb", "111000010110001110101110000011111101010110000101110110111110100110001100011100010011111010001000000110001110010100101001011010100011111111111110101011110111011"},
		{DigitsHex, DigitsBin, "8fc60e7c3b3c48e9a6a7a5fe4f1fbc31", "10001111110001100000111001111100001110110011110001001000111010011010011010100111101001011111111001001111000111111011110000110001"},

		{DigitsHex, Digits36, "abcdef00001234567890", "3o47re02jzqisvio"},
		{DigitsHex, Digits36, "abcdef01234567890123456789abcdef", "a65xa07491kf5zyfpvbo76g33"},

		{Digits62, DigitsHex, "cBaidlJ84Ggc5JA7IYCgv", "6ad547ffe02477b9473f7977e4d5e17"},
		{Digits62, DigitsHex, "4nipILgJlXPutO1hsisIJr", "8fc60e7c3b3c48e9a6a7a5fe4f1fbc31"},
		{Digits62, DigitsHex, "4vqyd6OoARXqj9nRUNhtLQ", "941532a06be1443aa9d5d57bdf180a52"},
		{Digits62, DigitsHex, "5FY8KwTsQaUJ2KzHJGetfE", "ba86b8f06fdf494487a08a491a19490e"},
		{Digits62, DigitsHex, "7N42dgm5tFLK9N8MT7fHC7", "ffffffffffffffffffffffffffffffff"},

		{DigitsDec, "Christopher", "355927353784509896715106760", "iihtspiphoeCrCeshhorsrrtrh"},
	}

	for n := 0; n < b.N; n++ {
		for i, test := range tests {
			v0, err := Convert(test.val, test.from, test.to)
			if err != nil {
				b.Fatal(err)
			}
			if test.exp != v0 {
				b.Errorf("test %d (%d->%d) expected %s, got: %s ", i, len(test.from), len(test.to), test.exp, v0)
			}
		}
	}
}
