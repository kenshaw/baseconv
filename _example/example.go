// _example/example.go
package main

import (
	"fmt"

	"github.com/kenshaw/baseconv"
)

func main() {
	valHex := "70b1d707eac2edf4c6389f440c7294b51fff57bb"
	fmt.Println("hex string: " + valHex)
	valDec, _ := baseconv.DecodeHexToDec(valHex)
	val62, _ := baseconv.Convert(valHex, baseconv.DigitsHex, baseconv.Digits62)
	val36, _ := baseconv.Convert(val62, baseconv.Digits62, baseconv.Digits36)

	fmt.Println("dec string: " + valDec)
	fmt.Println("62 string:  " + val62)
	fmt.Println("36 string:  " + val36)

	conVal36, _ := baseconv.Decode36ToDec(val36)
	fmt.Printf("dec and 36 values same: %t\n", valDec == conVal36)
}
