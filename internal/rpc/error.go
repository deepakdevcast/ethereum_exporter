package rpc

import (
	"errors"
)

var ErrJsonMarshal = errors.New("FAILED TO MARSHAL REQUEST")

var ErrStringConvert = errors.New("FAILED TO CONVERT BLOCK HEIGHT TO STRING")

var ErrNumberConvert = errors.New("FAILED TO CONVERT BLOCK HEIGHT TO NUMBER")
