package gcutmon

import (
	"math"

	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Decimal decimal.Decimal

func NewDecimalFromFloat(x float64) Decimal {
	return Decimal(decimal.NewFromFloat(x))
}

func NewDecimalFromFloatWithExp(x float64, exp int) Decimal {
	intValue := math.Floor(x / math.Pow10(exp))
	return Decimal(decimal.New(int64(intValue), int32(exp)))
}

func NewDecimalFromString(value string) (Decimal, error) {
	d, err := decimal.NewFromString(value)
	return Decimal(d), err
}

// Copy not yet implemented

// Abs not yet implemented

// Add returns d + d2
func (d Decimal) Add(d2 Decimal) Decimal {
	return Decimal(d.AsDecimal().Add(d2.AsDecimal()))
}

// Sub returns d - d2
func (d Decimal) Sub(d2 Decimal) Decimal {
	return Decimal(d.AsDecimal().Sub(d2.AsDecimal()))
}

// Neg returns -d
func (d Decimal) Neg() Decimal {
	return Decimal(d.AsDecimal().Neg())
}

// Mul returns d * d2.
func (d Decimal) Mul(d2 Decimal) Decimal {
	return Decimal(d.AsDecimal().Mul(d2.AsDecimal()))
}

// Shift not yet implemented

// Div returns d / d2. If it doesn't divide exactly, the result will have
// decimal.DivisionPrecision digits after the decimal point.
func (d Decimal) Div(d2 Decimal) Decimal {
	return Decimal(d.AsDecimal().Div(d2.AsDecimal()))
}

// QuoRem not yet implemented

// DivRound not yet implemented

// Mod not yet implemented

// Pow not yet implemented

// ExpHullAbrham not yet implemented

// ExpTaylor not yet implemented

// NumDigits not yet implemented

// IsInteger not yet implemented

// Cmp compares the numbers represented by d and d2 and returns:
//
//	-1 if d <  d2
//	 0 if d == d2
//	+1 if d >  d2
func (d Decimal) Cmp(d2 Decimal) int {
	return d.AsDecimal().Cmp(d2.AsDecimal())
}

// Equal returns whether the numbers represented by d and d2 are equal.
func (d Decimal) Equal(other Decimal) bool {
	return d.AsDecimal().Equal(other.AsDecimal())
}

// GreaterThan (GT) returns true when d is greater than d2.
func (d Decimal) GreaterThan(d2 Decimal) bool {
	return d.AsDecimal().GreaterThan(d2.AsDecimal())
}

// GreaterThanOrEqual (GTE) returns true when d is greater than or equal to d2.
func (d Decimal) GreaterThanOrEqual(d2 Decimal) bool {
	return d.AsDecimal().GreaterThanOrEqual(d2.AsDecimal())
}

// LessThan (LT) returns true when d is less than d2.
func (d Decimal) LessThan(d2 Decimal) bool {
	return d.AsDecimal().LessThan(d2.AsDecimal())
}

// LessThanOrEqual (LTE) returns true when d is less than or equal to d2.
func (d Decimal) LessThanOrEqual(d2 Decimal) bool {
	return d.AsDecimal().LessThanOrEqual(d2.AsDecimal())
}

// Sign not yet implemented

// IsPositive return
//
//	true if d > 0
//	false if d == 0
//	false if d < 0
func (d Decimal) IsPositive() bool {
	return d.AsDecimal().IsPositive()
}

// IsNegative return
//
//	true if d < 0
//	false if d == 0
//	false if d > 0
func (d Decimal) IsNegative() bool {
	return d.AsDecimal().IsNegative()
}

// IsZero return
//
//	true if d == 0
//	false if d > 0
//	false if d < 0
func (d Decimal) IsZero() bool {
	return d.AsDecimal().IsZero()
}

// Exponent not yet implemented

// Coefficient not yet implemented

// CoefficientInt64 not yet implemented

// IntPart not yet implemented

// BigInt not yet implemented

// BigFloat not yet implemented

// Rat not yet implemented

// Float64 not yet implemented

// InexactFloat64 not yet implemented

// String returns the string representation of the decimal
// with the fixed point.
//
// Example:
//
//	d := New(-12345, -3)
//	println(d.String())
//
// Output:
//
//	-12.345
func (d Decimal) String() string {
	return d.AsDecimal().String()
}

// StringFixed not yet implemented

// StringFixedBank not yet implemented

// StringFixedCash not yet implemented

// Round not yet implemented

// RoundCeil not yet implemented

// RoundFloor not yet implemented

// RoundUp not yet implemented

// RoundDown not yet implemented

// RoundBank not yet implemented

// RoundCash not yet implemented

// Floor not yet implemented

// Ceil not yet implemented

// Truncate truncates off digits from the number, without rounding.
//
// NOTE: precision is the last digit that will not be truncated (must be >= 0).
//
// Example:
//
//	decimal.NewFromString("123.456").Truncate(2).String() // "123.45"
func (d Decimal) Truncate(precision int32) Decimal {
	return Decimal(d.AsDecimal().Truncate(precision))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Decimal) UnmarshalJSON(b []byte) error {
	decimalValue := d.AsDecimal()
	if err := decimalValue.UnmarshalJSON(b); err != nil {
		return err
	}
	*d = Decimal(decimalValue)
	return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (d Decimal) MarshalJSON() ([]byte, error) {
	return d.AsDecimal().MarshalJSON()
}

// UnmarshalBinary not yet implemented

// MarshalBinary not yet implemented

// Scan not yet implemented

// Value not yet implemented

// UnmarshalText not yet implemented

// MarshalText not yet implemented

// GobEncode not yet implemented

// GobDecode not yet implemented

// StringScaled not yet implemented

// UnmarshalBSONValue implements the bson.ValueUnmarshaler interface.
func (d *Decimal) UnmarshalBSONValue(t bsontype.Type, data []byte) error {
	rawValue := bson.RawValue{Type: t, Value: data}

	var storedValue primitive.Decimal128
	if err := rawValue.Unmarshal(&storedValue); err != nil {
		return err
	}

	bi, exp, err := storedValue.BigInt()
	if err != nil {
		return err
	}

	*d = Decimal(decimal.NewFromBigInt(bi, int32(exp)))
	return nil
}

// MarshalBSONValue implements the bson.ValueMarshaler interface.
func (d Decimal) MarshalBSONValue() (bsontype.Type, []byte, error) {
	bsonDecimal, _ := primitive.ParseDecimal128FromBigInt(
		d.AsDecimal().Coefficient(),
		int(d.AsDecimal().Exponent()),
	)
	return bson.MarshalValue(bsonDecimal)
}

func (d Decimal) AsDecimal() decimal.Decimal {
	return decimal.Decimal(d)
}
