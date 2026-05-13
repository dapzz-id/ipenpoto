package helpers

import (
	"strconv"
)

// ConversionHelpers untuk konversi tipe data
type ConversionHelpers struct{}

// StringToInt mengkonversi string ke integer
func (c ConversionHelpers) StringToInt(str string) (int, error) {
	return strconv.Atoi(str)
}

// StringToInt64 mengkonversi string ke int64
func (c ConversionHelpers) StringToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

// IntToString mengkonversi integer ke string
func (c ConversionHelpers) IntToString(num int) string {
	return strconv.Itoa(num)
}

// Int64ToString mengkonversi int64 ke string
func (c ConversionHelpers) Int64ToString(num int64) string {
	return strconv.FormatInt(num, 10)
}

// StringToBool mengkonversi string ke boolean
func (c ConversionHelpers) StringToBool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

// BoolToString mengkonversi boolean ke string
func (c ConversionHelpers) BoolToString(b bool) string {
	return strconv.FormatBool(b)
}

// NewConversionHelpers membuat instance ConversionHelpers
func NewConversionHelpers() ConversionHelpers {
	return ConversionHelpers{}
}
