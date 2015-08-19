package baseconv

import "testing"

func TestErrors(t *testing.T) {
	_, err := Convert("", DigitsHex, DigitsDec)
	if err == nil {
		t.Error("bad string should return error")
	}

	_, err = Convert("0", "", DigitsDec)
	if err == nil {
		t.Error("bad fromBase should return error")
	}

	_, err = Convert("0", DigitsDec, "")
	if err == nil {
		t.Error("bad toBase should return error")
	}

	_, err = Convert("bad", DigitsBin, DigitsDec)
	if err == nil {
		t.Error("bad string should return error")
	}

	_, err = Convert("BAD", DigitsHex, DigitsDec)
	if err == nil {
		t.Error("bad string should return error")
	}
}

func TestDec2Bin(t *testing.T) {
	tests := map[string]string{
		"0":     "0",
		"8":     "1000",
		"15":    "1111",
		"16":    "10000",
		"88":    "1011000",
		"10000": "10011100010000",
	}

	for n, exp := range tests {
		v0, err := Convert(n, DigitsDec, DigitsBin)
		if err != nil {
			t.Fatal(err)
		}
		if exp != v0 {
			t.Errorf("on TestDec2Bin(%s) expected %s, got: %s", n, exp, v0)
		}

		v1, err := Convert(exp, DigitsBin, DigitsDec)
		if err != nil {
			t.Fatal(err)
		}
		if n != v1 {
			t.Errorf("on TestBin2Dec(%s) expected %s, got: %s", exp, n, v1)
		}
	}
}

func TestDec2Hex(t *testing.T) {
	tests := map[string]string{
		"0":     "0",
		"8":     "8",
		"15":    "f",
		"16":    "10",
		"88":    "58",
		"10000": "2710",
	}

	for n, exp := range tests {
		v0, err := Convert(n, DigitsDec, DigitsHex)
		if err != nil {
			t.Fatal(err)
		}
		if exp != v0 {
			t.Errorf("on TestDec2Hex(%s) expected %s, got: %s", n, exp, v0)
		}

		v1, err := Convert(exp, DigitsHex, DigitsDec)
		if err != nil {
			t.Fatal(err)
		}
		if n != v1 {
			t.Errorf("on TestHex2Dec(%s) expected %s, got: %s", exp, n, v1)
		}
	}
}

func TestDec2SixtyTwo(t *testing.T) {
	tests := map[string]string{
		"16571982744576742462": "jKbR7u8J5PU",
		"46394851265279874948": "TheUtUU3miE",
		"21901407667833273510": "q5SG7U76tls",
		"8232087098322120342":  "9O72RLP5fF4",
		"6354358749246709610":  "7zp1TbLFPp8",
		"18089061068":          "jKbR7u",
		"50642057182":          "TheUtU",
		"23906366962":          "q5SG7U",
		"8985691605":           "9O72RL",
		"6936067049":           "7zp1Tb",
		"799310853702667":      "3EYjA0o7p",
	}

	for n, exp := range tests {
		v0, err := Convert(n, DigitsDec, Digits62)
		if err != nil {
			t.Fatal(err)
		}
		if exp != v0 {
			t.Errorf("on TestDec2SixtyTwo(%s) expected %s, got: %s", n, exp, v0)
		}

		v1, err := Convert(exp, Digits62, DigitsDec)
		if err != nil {
			t.Fatal(err)
		}
		if n != v1 {
			t.Errorf("on TestSixtyTwo2Dec(%s) expected %s, got: %s", exp, n, v1)
		}
	}
}

func TestArbitrary(t *testing.T) {
	tests := map[string]string{
		"355927353784509896715106760": "iihtspiphoeCrCeshhorsrrtrh",
	}

	arbDigits := "Christopher"

	for n, exp := range tests {
		v0, err := Convert(n, DigitsDec, arbDigits)
		if err != nil {
			t.Fatal(err)
		}
		if exp != v0 {
			t.Errorf("on TestBin2Hex(%s) expected %s, got: %s", n, exp, v0)
		}

		v1, err := Convert(exp, arbDigits, DigitsDec)
		if err != nil {
			t.Fatal(err)
		}
		if n != v1 {
			t.Errorf("on TestHex2Bin(%s) expected %s, got: %s", exp, n, v1)
		}
	}
}

func TestHex2Bin(t *testing.T) {
	tests := map[string]string{
		"70b1d707eac2edf4c6389f440c7294b51fff57bb": "111000010110001110101110000011111101010110000101110110111110100110001100011100010011111010001000000110001110010100101001011010100011111111111110101011110111011",
		"8fc60e7c3b3c48e9a6a7a5fe4f1fbc31":         "10001111110001100000111001111100001110110011110001001000111010011010011010100111101001011111111001001111000111111011110000110001",
	}

	for n, exp := range tests {
		v0, err := Convert(n, DigitsHex, DigitsBin)
		if err != nil {
			t.Fatal(err)
		}
		if exp != v0 {
			t.Errorf("on TestBin2Hex(%s) expected %s, got: %s", n, exp, v0)
		}

		v1, err := Convert(exp, DigitsBin, DigitsHex)
		if err != nil {
			t.Fatal(err)
		}
		if n != v1 {
			t.Errorf("on TestHex2Bin(%s) expected %s, got: %s", exp, n, v1)
		}
	}
}

func TestSixtyTwo2Hex(t *testing.T) {
	tests := map[string]string{
		"cBaidlJ84Ggc5JA7IYCgv":  "6ad547ffe02477b9473f7977e4d5e17",
		"4nipILgJlXPutO1hsisIJr": "8fc60e7c3b3c48e9a6a7a5fe4f1fbc31",
		"4vqyd6OoARXqj9nRUNhtLQ": "941532a06be1443aa9d5d57bdf180a52",
		"5FY8KwTsQaUJ2KzHJGetfE": "ba86b8f06fdf494487a08a491a19490e",
		"7N42dgm5tFLK9N8MT7fHC7": "ffffffffffffffffffffffffffffffff",
	}

	for n, exp := range tests {
		v0, err := Convert(n, Digits62, DigitsHex)
		if err != nil {
			t.Fatal(err)
		}
		if exp != v0 {
			t.Errorf("on TestSixtyTwo2Hex(%s) expected %s, got: %s", n, exp, v0)
		}

		v1, err := Convert(exp, DigitsHex, Digits62)
		if err != nil {
			t.Fatal(err)
		}
		if n != v1 {
			t.Errorf("on TestHex2SixtyTwo(%s) expected %s, got: %s", exp, n, v1)
		}
	}
}

func TestEncodeDecode(t *testing.T) {
	v0 := "1627734050041231452076"

	efuncs := []func(string, ...string) (string, error){
		EncodeBin,
		EncodeOct,
		EncodeHex,
		Encode36,
		Encode62,
		Encode64,
	}

	dfuncs := []func(string, ...string) (string, error){
		DecodeBin,
		DecodeOct,
		DecodeHex,
		Decode36,
		Decode62,
		Decode64,
	}

	/*digits := []string{
		DigitsBin,
		DigitsOct,
		DigitsHex,
		Digits36,
		Digits62,
		Digits64,
	}*/

	evals := []string{
		"10110000011110101011001000001100110100001101011100000100010011110101100",
		"260365310146415340423654",
		"583d5906686b8227ac",
		"9jird8fbzkui7g",
		"vhozdwL3WC8A",
		"m3Rp1CxHwyuI",
	}

	for i, exp := range evals {
		v1, err := efuncs[i](v0)
		if err != nil {
			t.Fatal(err)
		}
		if exp != v1 {
			t.Errorf("values %s / %s should match", exp, v1)
		}

		v2, err := dfuncs[i](v1)
		if err != nil {
			t.Fatal(err)
		}
		if v0 != v2 {
			t.Errorf("values %s / %s should match", v0, v2)
		}

		v3, err := efuncs[i](v0, DigitsDec)
		if err != nil {
			t.Fatal(err)
		}
		if exp != v3 {
			t.Errorf("values %s / %s should match", exp, v3)
		}

		v4, err := dfuncs[i](v1, DigitsDec)
		if err != nil {
			t.Fatal(err)
		}
		if v0 != v4 {
			t.Errorf("values %s / %s should match", v0, v4)
		}
	}
}
