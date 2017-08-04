# About baseconv [![Build Status](https://travis-ci.org/knq/baseconv.svg)](https://travis-ci.org/knq/baseconv) [![Coverage Status](https://coveralls.io/repos/knq/baseconv/badge.svg?branch=master&service=github)](https://coveralls.io/github/knq/baseconv?branch=master) #

A simple [Go](http://www.golang.org/project/) package for converting between
strings in arbitrary bases.

This package is useful when working with extremely large numbers (larger than
int64), and need to convert them to different base (ie, decimal, hex, octal,
etc) representations, and thus cannot use the standard go libraries.

This was written for a specific use case where there was a need to
encode/decode large numbers stored as strings in a database.

This is similar in concept to PHP's [```base_convert```](http://php.net/manual/en/function.base-convert.php)
function.

## Installation ##

Install the package via the following:

    go get -u github.com/knq/baseconv

## Usage ##

Please see [the GoDoc API page](http://godoc.org/github.com/knq/baseconv) for a
full API listing.

The baseconv package can be used similarly to the following:
```go
// example/example.go
package main

import (
	"fmt"

	"github.com/knq/baseconv"
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
```

Example output:
```sh
$ go run example.go
hex string: 70b1d707eac2edf4c6389f440c7294b51fff57bb
dec string: 643372930067913326838082478477533553256088688571
62 string:  g4WuOGCMWgcPa70d91BezVvvvaX
36 string:  d5wjfaew7fypqn2ka6xpofdlwns9ha3
dec and 36 values same: true
```
