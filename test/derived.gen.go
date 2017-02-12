//DO NOT EDIT generated by goderive

package test

import (
	"bytes"
)

var _ = bytes.MinRead

func deriveEqualPtrToBuiltInTypes(this, that *BuiltInTypes) bool {
	return (this == nil && that == nil) ||
		(this != nil && that != nil) &&
			this.Bool == that.Bool &&
			this.Byte == that.Byte &&
			this.Complex128 == that.Complex128 &&
			this.Complex64 == that.Complex64 &&
			this.Float64 == that.Float64 &&
			this.Float32 == that.Float32 &&
			this.Int == that.Int &&
			this.Int16 == that.Int16 &&
			this.Int32 == that.Int32 &&
			this.Int64 == that.Int64 &&
			this.Int8 == that.Int8 &&
			this.Rune == that.Rune &&
			this.String == that.String &&
			this.Uint == that.Uint &&
			this.Uint16 == that.Uint16 &&
			this.Uint32 == that.Uint32 &&
			this.Uint64 == that.Uint64 &&
			this.Uint8 == that.Uint8 &&
			this.UintPtr == that.UintPtr
}
func deriveEqualPtrToPtrToBuiltInTypes(this, that *PtrToBuiltInTypes) bool {
	return (this == nil && that == nil) ||
		(this != nil && that != nil) &&
			((this.Bool == nil && that.Bool == nil) || (this.Bool != nil && that.Bool != nil && *this.Bool == *that.Bool)) &&
			((this.Byte == nil && that.Byte == nil) || (this.Byte != nil && that.Byte != nil && *this.Byte == *that.Byte)) &&
			((this.Complex128 == nil && that.Complex128 == nil) || (this.Complex128 != nil && that.Complex128 != nil && *this.Complex128 == *that.Complex128)) &&
			((this.Complex64 == nil && that.Complex64 == nil) || (this.Complex64 != nil && that.Complex64 != nil && *this.Complex64 == *that.Complex64)) &&
			((this.Float64 == nil && that.Float64 == nil) || (this.Float64 != nil && that.Float64 != nil && *this.Float64 == *that.Float64)) &&
			((this.Float32 == nil && that.Float32 == nil) || (this.Float32 != nil && that.Float32 != nil && *this.Float32 == *that.Float32)) &&
			((this.Int == nil && that.Int == nil) || (this.Int != nil && that.Int != nil && *this.Int == *that.Int)) &&
			((this.Int16 == nil && that.Int16 == nil) || (this.Int16 != nil && that.Int16 != nil && *this.Int16 == *that.Int16)) &&
			((this.Int32 == nil && that.Int32 == nil) || (this.Int32 != nil && that.Int32 != nil && *this.Int32 == *that.Int32)) &&
			((this.Int64 == nil && that.Int64 == nil) || (this.Int64 != nil && that.Int64 != nil && *this.Int64 == *that.Int64)) &&
			((this.Int8 == nil && that.Int8 == nil) || (this.Int8 != nil && that.Int8 != nil && *this.Int8 == *that.Int8)) &&
			((this.Rune == nil && that.Rune == nil) || (this.Rune != nil && that.Rune != nil && *this.Rune == *that.Rune)) &&
			((this.String == nil && that.String == nil) || (this.String != nil && that.String != nil && *this.String == *that.String)) &&
			((this.Uint == nil && that.Uint == nil) || (this.Uint != nil && that.Uint != nil && *this.Uint == *that.Uint)) &&
			((this.Uint16 == nil && that.Uint16 == nil) || (this.Uint16 != nil && that.Uint16 != nil && *this.Uint16 == *that.Uint16)) &&
			((this.Uint32 == nil && that.Uint32 == nil) || (this.Uint32 != nil && that.Uint32 != nil && *this.Uint32 == *that.Uint32)) &&
			((this.Uint64 == nil && that.Uint64 == nil) || (this.Uint64 != nil && that.Uint64 != nil && *this.Uint64 == *that.Uint64)) &&
			((this.Uint8 == nil && that.Uint8 == nil) || (this.Uint8 != nil && that.Uint8 != nil && *this.Uint8 == *that.Uint8)) &&
			((this.UintPtr == nil && that.UintPtr == nil) || (this.UintPtr != nil && that.UintPtr != nil && *this.UintPtr == *that.UintPtr))
}
func deriveEqualPtrToSliceOfBuiltInTypes(this, that *SliceOfBuiltInTypes) bool {
	return (this == nil && that == nil) ||
		(this != nil && that != nil) &&
			deriveEqualSliceOfbool(this.Bool, that.Bool) &&
			bytes.Equal(this.Byte, that.Byte) &&
			deriveEqualSliceOfcomplex128(this.Complex128, that.Complex128) &&
			deriveEqualSliceOfcomplex64(this.Complex64, that.Complex64) &&
			deriveEqualSliceOffloat64(this.Float64, that.Float64) &&
			deriveEqualSliceOffloat32(this.Float32, that.Float32) &&
			deriveEqualSliceOfint(this.Int, that.Int) &&
			deriveEqualSliceOfint16(this.Int16, that.Int16) &&
			deriveEqualSliceOfint32(this.Int32, that.Int32) &&
			deriveEqualSliceOfint64(this.Int64, that.Int64) &&
			deriveEqualSliceOfint8(this.Int8, that.Int8) &&
			deriveEqualSliceOfrune(this.Rune, that.Rune) &&
			deriveEqualSliceOfstring(this.String, that.String) &&
			deriveEqualSliceOfuint(this.Uint, that.Uint) &&
			deriveEqualSliceOfuint16(this.Uint16, that.Uint16) &&
			deriveEqualSliceOfuint32(this.Uint32, that.Uint32) &&
			deriveEqualSliceOfuint64(this.Uint64, that.Uint64) &&
			bytes.Equal(this.Uint8, that.Uint8) &&
			deriveEqualSliceOfuintptr(this.UintPtr, that.UintPtr)
}
func deriveEqualPtrToSomeComplexTypes(this, that *SomeComplexTypes) bool {
	return (this == nil && that == nil) ||
		(this != nil && that != nil) &&
			deriveEqualSliceOfPtrToRecursiveType(this.J, that.J) &&
			deriveEqualSliceOfRecursiveType(this.K, that.K) &&
			this.L.Equal(that.L) &&
			this.M.Equal(&that.M) &&
			deriveEqualMapOfintToRecursiveType(this.N, that.N) &&
			deriveEqualMapOfstringToPtrToRecursiveType(this.O, that.O) &&
			deriveEqualMapOfint64Tostring(this.P, that.P)
}
func deriveEqualPtrToRecursiveType(this, that *RecursiveType) bool {
	return (this == nil && that == nil) ||
		(this != nil && that != nil) &&
			bytes.Equal(this.Bytes, that.Bytes) &&
			deriveEqualMapOfintToRecursiveType(this.N, that.N)
}
func deriveEqualSliceOfrune(this, that []rune) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfstring(this, that []string) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfuint32(this, that []uint32) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfuint64(this, that []uint64) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualMapOfintToRecursiveType(this, that map[int]RecursiveType) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(v.Equal(&thatv)) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOffloat32(this, that []float32) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfuint16(this, that []uint16) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfint64(this, that []int64) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfint8(this, that []int8) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualMapOfint64Tostring(this, that map[int64]string) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(v == thatv) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOffloat64(this, that []float64) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfint16(this, that []int16) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfuint(this, that []uint) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfcomplex128(this, that []complex128) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfcomplex64(this, that []complex64) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfPtrToRecursiveType(this, that []*RecursiveType) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i].Equal(that[i])) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfbool(this, that []bool) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfint(this, that []int) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfint32(this, that []int32) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
func deriveEqualMapOfstringToPtrToRecursiveType(this, that map[string]*RecursiveType) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for k, v := range this {
		thatv, ok := that[k]
		if !ok {
			return false
		}
		if !(v.Equal(thatv)) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfRecursiveType(this, that []RecursiveType) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i].Equal(&that[i])) {
			return false
		}
	}
	return true

}
func deriveEqualSliceOfuintptr(this, that []uintptr) bool {
	if this == nil || that == nil {
		return this == nil && that == nil
	}
	if len(this) != len(that) {
		return false
	}
	for i := 0; i < len(this); i++ {
		if !(this[i] == that[i]) {
			return false
		}
	}
	return true

}
