# About baseconv [![Build Status](https://travis-ci.org/knq/baseconv.svg)](https://travis-ci.org/knq/baseconv) [![Coverage Status](https://coveralls.io/repos/knq/baseconv/badge.svg?branch=master&service=github)](https://coveralls.io/github/knq/baseconv?branch=master) #

A simple [Go](http://www.golang.org/project/) package for converting between
strings in arbitrary bases.

## Installation ##

Install the package via the following:

    go get -u github.com/knq/baseconv

## Usage ##

The baseconv package can be used similarly to the following:

```go
package main

import (
    "fmt"

    "github.com/knq/baseconv"
)

func main() {
    valHex := "70b1d707eac2edf4c6389f440c7294b51fff57bb"
    valDec := baseconv.DecodeHex(valHex)
    val62 := baseconv.Convert(valHex, baseconv.DigitsHex, baseconv.Digits62)
    val36 := baseconv.Convert(val62, baseconv.Digits62, baseconv.Digits36)

    fmt.Println("dec string: " + valDec)
    fmt.Println("62 string:  " + val62)
    fmt.Println("36 string:  " + val36)

    fmt.Printf("dec and 36 values same: %t\n", valDec == baseconv.Decode36(val36, baseconv.DigitsDec))
}
```
