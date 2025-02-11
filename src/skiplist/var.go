package skiplist

import "bytes"

const PROPABILITY = 0x3FFF

var (
	DefaultMaxLevel int = 32
	defaultSource       = defaultRandSource{}

	Byte GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(byte) > rhs.(byte)
	}
	ByteAscending               = Byte
	ByteAsc                     = Byte
	ByteDescending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(byte) < rhs.(byte)
	}
	ByteDesc LessThanFunc = ByteDescending

	Float32 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(float32) > rhs.(float32)
	}
	Float32Ascending               = Float32
	Float32Asc                     = Float32
	Float32Descending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(float32) < rhs.(float32)
	}
	Float32Desc LessThanFunc = Float32Descending

	Float64 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(float64) > rhs.(float64)
	}
	Float64Ascending               = Float64
	Float64Asc                     = Float64
	Float64Descending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(float64) < rhs.(float64)
	}
	Float64Desc LessThanFunc = Float64Descending

	Int GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int) > rhs.(int)
	}
	IntAscending               = Int
	IntAsc                     = Int
	IntDescending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int) < rhs.(int)
	}
	IntDesc LessThanFunc = IntDescending

	Int16 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int16) > rhs.(int16)
	}
	Int16Ascending               = Int16
	Int16Asc                     = Int16
	Int16Descending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int16) < rhs.(int16)
	}
	Int16Desc LessThanFunc = Int16Descending

	Int32 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int32) > rhs.(int32)
	}
	Int32Ascending               = Int32
	Int32Asc                     = Int32
	Int32Descending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int32) < rhs.(int32)
	}
	Int32Desc LessThanFunc = Int32Descending

	Int64 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int64) > rhs.(int64)
	}
	Int64Ascending               = Int64
	Int64Asc                     = Int64
	Int64Descending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int64) < rhs.(int64)
	}
	Int64Desc LessThanFunc = Int64Descending

	Int8 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int8) > rhs.(int8)
	}
	Int8Ascending               = Int8
	Int8Asc                     = Int8
	Int8Descending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(int8) < rhs.(int8)
	}
	Int8Desc LessThanFunc = Int8Descending

	Rune GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(rune) > rhs.(rune)
	}
	RuneAscending               = Rune
	RuneAsc                     = Rune
	RuneDescending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(rune) < rhs.(rune)
	}
	RuneDesc LessThanFunc = RuneDescending

	String GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(string) > rhs.(string)
	}
	StringAscending               = String
	StringAsc                     = String
	StringDescending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(string) < rhs.(string)
	}
	StringDesc LessThanFunc = StringDescending

	Uint GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint) > rhs.(uint)
	}
	UintAscending               = Uint
	UintAsc                     = Uint
	UintDescending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint) < rhs.(uint)
	}
	UintDesc LessThanFunc = UintDescending

	Uint16 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint16) > rhs.(uint16)
	}
	Uint16Ascending               = Uint16
	Uint16Asc                     = Uint16
	Uint16Descending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint16) < rhs.(uint16)
	}
	Uint16Desc LessThanFunc = Uint16Descending

	Uint32 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint32) > rhs.(uint32)
	}
	Uint32Ascending               = Uint32
	Uint32Asc                     = Uint32
	Uint32Descending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint32) < rhs.(uint32)
	}
	Uint32Desc LessThanFunc = Uint32Descending

	Uint64 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint64) > rhs.(uint64)
	}
	Uint64Ascending               = Uint64
	Uint64Asc                     = Uint64
	Uint64Descending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint64) < rhs.(uint64)
	}
	Uint64Desc LessThanFunc = Uint64Descending

	Uint8 GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint8) > rhs.(uint8)
	}
	Uint8Ascending               = Uint8
	Uint8Asc                     = Uint8
	Uint8Descending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uint8) < rhs.(uint8)
	}
	Uint8Desc LessThanFunc = Uint8Descending

	Uintptr GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uintptr) > rhs.(uintptr)
	}
	UintptrAscending               = Uintptr
	UintptrAsc                     = Uintptr
	UintptrDescending LessThanFunc = func(lhs, rhs interface{}) bool {
		return lhs.(uintptr) < rhs.(uintptr)
	}
	UintptrDesc LessThanFunc = UintptrDescending

	// the type []byte.
	Bytes GreaterThanFunc = func(lhs, rhs interface{}) bool {
		return bytes.Compare(lhs.([]byte), rhs.([]byte)) > 0
	}
	BytesAscending = Bytes
	BytesAsc       = Bytes
	// the type []byte. reversed order.
	BytesDescending LessThanFunc = func(lhs, rhs interface{}) bool {
		return bytes.Compare(lhs.([]byte), rhs.([]byte)) < 0
	}
	BytesDesc LessThanFunc = BytesDescending
)
