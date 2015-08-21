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

func TestConvert(t *testing.T) {
	tests := [][]string{
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

	for idx, test := range tests {
		from := test[0]
		to := test[1]
		exp0 := test[2]
		exp1 := test[3]

		v0, err := Convert(exp0, from, to)
		if err != nil {
			t.Fatal(err)
		}
		if exp1 != v0 {
			t.Errorf("on test %d (%d->%d) expected %s, got: %s ", idx, len(from), len(to), exp1, v0)
		}

		v1, err := Convert(exp1, to, from)
		if err != nil {
			t.Fatal(err)
		}
		if exp0 != v1 {
			t.Errorf("on test %d (%d->%d) expected %s, got: %s ", idx, len(to), len(from), exp0, v1)
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
