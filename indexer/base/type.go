package indexerbase

import (
	"encoding/json"
	"fmt"
	"time"
)

type Kind int

const (
	InvalidKind Kind = iota

	// StringKind is a string type and values of this type must be of the go type string
	// or implement fmt.Stringer().
	StringKind

	// BytesKind is a bytes type and values of this type must be of the go type []byte.
	BytesKind

	// Int8Kind is an int8 type and values of this type must be of the go type int8.
	Int8Kind

	// Uint8Kind is a uint8 type and values of this type must be of the go type uint8.
	Uint8Kind

	// Int16Kind is an int16 type and values of this type must be of the go type int16.
	Int16Kind

	// Uint16Kind is a uint16 type and values of this type must be of the go type uint16.
	Uint16Kind

	// Int32Kind is an int32 type and values of this type must be of the go type int32.
	Int32Kind

	// Uint32Kind is a uint32 type and values of this type must be of the go type uint32.
	Uint32Kind

	// Int64Kind is an int64 type and values of this type must be of the go type int64.
	Int64Kind

	// Uint64Kind is a uint64 type and values of this type must be of the go type uint64.
	Uint64Kind

	// DecimalKind represents an arbitrary precision decimal or integer number. Values of this type
	// must be of the go type string or a type that implements fmt.Stringer with the resulting string
	// formatted as decimal numbers with an optional fractional part. Exponential E-notation
	// is supported but NaN and Infinity are not.
	DecimalKind

	// BoolKind is a boolean type and values of this type must be of the go type bool.
	BoolKind

	// TimeKind is a time type and values of this type must be of the go type time.Time.
	TimeKind

	// DurationKind is a duration type and values of this type must be of the go type time.Duration.
	DurationKind

	// Float32Kind is a float32 type and values of this type must be of the go type float32.
	Float32Kind

	// Float64Kind is a float64 type and values of this type must be of the go type float64.
	Float64Kind

	// Bech32AddressKind is a bech32 address type and values of this type must be of the go type string or []byte.
	// Columns of this type are expected to set the AddressPrefix field in the column definition to the bech32
	// address prefix.
	Bech32AddressKind

	// EnumKind is an enum type and values of this type must be of the go type string or implement fmt.Stringer.
	// Columns of this type are expected to set the EnumDefinition field in the column definition to the enum definition.
	EnumKind

	// JSONKind is a JSON type and values of this type can either be of go type json.RawMessage
	// or any type that can be marshaled to JSON using json.Marshal.
	JSONKind
)

func (t Kind) Validate() error {
	if t <= InvalidKind {
		return fmt.Errorf("unknown type: %d", t)
	}
	if t > JSONKind {
		return fmt.Errorf("invalid type: %d", t)
	}
	return nil
}

func (t Kind) ValidateValue(value any) error {
	switch t {
	case StringKind:
		_, ok := value.(string)
		_, ok2 := value.(fmt.Stringer)
		if !ok && !ok2 {
			return fmt.Errorf("expected string or type that implements fmt.Stringer, got %T", value)
		}
	case BytesKind:
		_, ok := value.([]byte)
		if !ok {
			return fmt.Errorf("expected []byte, got %T", value)
		}
	case Int8Kind:
		_, ok := value.(int8)
		if !ok {
			return fmt.Errorf("expected int8, got %T", value)
		}
	case Uint8Kind:
		_, ok := value.(uint8)
		if !ok {
			return fmt.Errorf("expected uint8, got %T", value)
		}
	case Int16Kind:
		_, ok := value.(int16)
		if !ok {
			return fmt.Errorf("expected int16, got %T", value)
		}
	case Uint16Kind:
		_, ok := value.(uint16)
		if !ok {
			return fmt.Errorf("expected uint16, got %T", value)
		}
	case Int32Kind:
		_, ok := value.(int32)
		if !ok {
			return fmt.Errorf("expected int32, got %T", value)
		}
	case Uint32Kind:
		_, ok := value.(uint32)
		if !ok {
			return fmt.Errorf("expected uint32, got %T", value)
		}
	case Int64Kind:
		_, ok := value.(int64)
		if !ok {
			return fmt.Errorf("expected int64, got %T", value)
		}
	case Uint64Kind:
		_, ok := value.(uint64)
		if !ok {
			return fmt.Errorf("expected uint64, got %T", value)
		}
	case DecimalKind:
		_, ok := value.(string)
		_, ok2 := value.(fmt.Stringer)
		if !ok && !ok2 {
			return fmt.Errorf("expected string or type that implements fmt.Stringer, got %T", value)
		}
	case BoolKind:
		_, ok := value.(bool)
		if !ok {
			return fmt.Errorf("expected bool, got %T", value)
		}
	case TimeKind:
		_, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("expected time.Time, got %T", value)
		}
	case DurationKind:
		_, ok := value.(time.Duration)
		if !ok {
			return fmt.Errorf("expected time.Duration, got %T", value)
		}
	case Float32Kind:
		_, ok := value.(float32)
		if !ok {
			return fmt.Errorf("expected float32, got %T", value)
		}
	case Float64Kind:
		_, ok := value.(float64)
		if !ok {
			return fmt.Errorf("expected float64, got %T", value)
		}
	case Bech32AddressKind:
		_, ok := value.(string)
		_, ok2 := value.([]byte)
		if !ok && !ok2 {
			return fmt.Errorf("expected string or []byte, got %T", value)
		}
	case EnumKind:
		_, ok := value.(string)
		_, ok2 := value.(fmt.Stringer)
		if !ok && !ok2 {
			return fmt.Errorf("expected string or type that implements fmt.Stringer, got %T", value)
		}
	case JSONKind:
		return nil
	default:
		return fmt.Errorf("invalid type: %d", t)
	}
	return nil
}

func (t Kind) String() string {
	switch t {
	case StringKind:
		return "string"
	case BytesKind:
		return "bytes"
	case Int8Kind:
		return "int8"
	case Uint8Kind:
		return "uint8"
	case Int16Kind:
		return "int16"
	case Uint16Kind:
		return "uint16"
	case Int32Kind:
		return "int32"
	case Uint32Kind:
		return "uint32"
	case Int64Kind:
		return "int64"
	case Uint64Kind:
		return "uint64"
	case DecimalKind:
		return "decimal"
	case BoolKind:
		return "bool"
	case TimeKind:
		return "time"
	case DurationKind:
		return "duration"
	case Float32Kind:
		return "float32"
	case Float64Kind:
		return "float64"
	case Bech32AddressKind:
		return "bech32address"
	case EnumKind:
		return "enum"
	case JSONKind:
		return "json"
	default:
		return ""
	}
}

func TypeForGoValue(value any) Kind {
	switch value.(type) {
	case string, fmt.Stringer:
		return StringKind
	case []byte:
		return BytesKind
	case int8:
		return Int8Kind
	case uint8:
		return Uint8Kind
	case int16:
		return Int16Kind
	case uint16:
		return Uint16Kind
	case int32:
		return Int32Kind
	case uint32:
		return Uint32Kind
	case int64:
		return Int64Kind
	case uint64:
		return Uint64Kind
	case float32:
		return Float32Kind
	case float64:
		return Float64Kind
	case bool:
		return BoolKind
	case time.Time:
		return TimeKind
	case time.Duration:
		return DurationKind
	case json.RawMessage:
		return JSONKind
	default:
	}

	if _, ok := value.(fmt.Stringer); ok {
		return StringKind
	}

	return InvalidKind
}