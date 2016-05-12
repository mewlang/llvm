// Package types declares the data types of LLVM IR.
package types

import "fmt"

// A Type represents an LLVM IR type.
//
// Type is one of the following concrete types:
//    *types.Void
//    *types.Int
//    *types.Float
//    *types.MMX
//    *types.Label
//    *types.Metadata
//    *types.Func
//    *types.Pointer
//    *types.Vector
//    *types.Array
//    *types.Struct
//
// References:
//    http://llvm.org/docs/LangRef.html#typesystem
type Type interface {
	fmt.Stringer
	// Equal reports whether t and u are of equal type.
	Equal(u Type) bool
}

// Make sure that each type implements the Type interface.
var (
	_ Type = &Void{}
	_ Type = &Int{}
	_ Type = &Float{}
	_ Type = &MMX{}
	_ Type = &Label{}
	_ Type = &Metadata{}
	_ Type = &Func{}
	_ Type = &Pointer{}
	_ Type = &Vector{}
	_ Type = &Array{}
	_ Type = &Struct{}
)

// Equal reports whether t and u are of equal type.
func Equal(t, u Type) bool {
	return t.Equal(u)
}

// IsInt returns true if t is an integer type, and false otherwise.
func IsInt(t Type) bool {
	_, ok := t.(*Int)
	return ok
}

// IsInts returns true if t is an integer type or a vector of integers type, and
// false otherwise.
func IsInts(t Type) bool {
	if t, ok := t.(*Vector); ok {
		return IsInt(t.Elem())
	}
	return IsInt(t)
}

// IsFloat returns true if t is a floating point type, and false otherwise.
func IsFloat(t Type) bool {
	_, ok := t.(*Float)
	return ok
}

// IsFloats returns true if t is a floating point type or a vector of floating
// points type, and false otherwise.
func IsFloats(t Type) bool {
	if t, ok := t.(*Vector); ok {
		return IsFloat(t.Elem())
	}
	return IsFloat(t)
}

// SameLength returns true if both types are vectors or arrays of the same
// length or if both types are distinct from vectors and arrays, and false
// otherwise.
func SameLength(a, b Type) bool {
	type Lener interface {
		Len() int
	}
	l1, ok1 := a.(Lener)
	l2, ok2 := b.(Lener)

	// Verify if both types are vectors or arrays of the same length.
	if ok1 && ok2 {
		return l1.Len() == l2.Len()
	}

	// Verify if both types are distinct from vectors and arrays.
	return !ok1 && !ok2
}

// Convenience types.
var (
	// I8 represents the i8 type.
	I8 *Int
)

func init() {
	I8, _ = NewInt(8)
}
