package constant

import (
	"fmt"
	"strings"

	"github.com/llir/llvm/ir/enum"
	"github.com/llir/llvm/ir/types"
)

// --- [ Binary expressions ] --------------------------------------------------

// ~~~ [ add ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprAdd is an LLVM IR add expression.
type ExprAdd struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
	// (optional) Integer overflow flags.
	OverflowFlags []enum.OverflowFlag
}

func (e *ExprAdd) IsValue() {
}

// NewAdd returns a new add expression based on the given operands.
func NewAdd(x, y Constant) *ExprAdd {
	e := &ExprAdd{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprAdd) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprAdd) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprAdd) Ident() string {
	// 'add' OverflowFlags=OverflowFlag* '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("add")
	for _, flag := range e.OverflowFlags {
		fmt.Fprintf(buf, " %s", flag)
	}
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprAdd) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ fadd ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFAdd is an LLVM IR fadd expression.
type ExprFAdd struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
}

func (e *ExprFAdd) IsValue() {
}

// NewFAdd returns a new fadd expression based on the given operands.
func NewFAdd(x, y Constant) *ExprFAdd {
	e := &ExprFAdd{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFAdd) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFAdd) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFAdd) Ident() string {
	// 'fadd' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("fadd (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFAdd) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ sub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSub is an LLVM IR sub expression.
type ExprSub struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
	// (optional) Integer overflow flags.
	OverflowFlags []enum.OverflowFlag
}

func (e *ExprSub) IsValue() {
}

// NewSub returns a new sub expression based on the given operands.
func NewSub(x, y Constant) *ExprSub {
	e := &ExprSub{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSub) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSub) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSub) Ident() string {
	// 'sub' OverflowFlags=OverflowFlag* '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("sub")
	for _, flag := range e.OverflowFlags {
		fmt.Fprintf(buf, " %s", flag)
	}
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSub) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ fsub ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFSub is an LLVM IR fsub expression.
type ExprFSub struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
}

func (e *ExprFSub) IsValue() {
}

// NewFSub returns a new fsub expression based on the given operands.
func NewFSub(x, y Constant) *ExprFSub {
	e := &ExprFSub{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFSub) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFSub) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFSub) Ident() string {
	// 'fsub' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("fsub (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFSub) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ mul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprMul is an LLVM IR mul expression.
type ExprMul struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
	// (optional) Integer overflow flags.
	OverflowFlags []enum.OverflowFlag
}

func (e *ExprMul) IsValue() {
}

// NewMul returns a new mul expression based on the given operands.
func NewMul(x, y Constant) *ExprMul {
	e := &ExprMul{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprMul) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprMul) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprMul) Ident() string {
	// 'mul' OverflowFlags=OverflowFlag* '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("mul")
	for _, flag := range e.OverflowFlags {
		fmt.Fprintf(buf, " %s", flag)
	}
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprMul) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ fmul ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFMul is an LLVM IR fmul expression.
type ExprFMul struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
}

func (e *ExprFMul) IsValue() {
}

// NewFMul returns a new fmul expression based on the given operands.
func NewFMul(x, y Constant) *ExprFMul {
	e := &ExprFMul{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFMul) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFMul) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFMul) Ident() string {
	// 'fmul' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("fmul (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFMul) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ udiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprUDiv is an LLVM IR udiv expression.
type ExprUDiv struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
	// (optional) The result is a poison value if X is not a multiple of Y.
	Exact bool
}

func (e *ExprUDiv) IsValue() {
}

// NewUDiv returns a new udiv expression based on the given operands.
func NewUDiv(x, y Constant) *ExprUDiv {
	e := &ExprUDiv{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprUDiv) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprUDiv) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprUDiv) Ident() string {
	// 'udiv' Exactopt '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("udiv")
	if e.Exact {
		buf.WriteString(" exact")
	}
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprUDiv) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ sdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSDiv is an LLVM IR sdiv expression.
type ExprSDiv struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
	// (optional) The result is a poison value if the result would be rounded.
	Exact bool
}

func (e *ExprSDiv) IsValue() {
}

// NewSDiv returns a new sdiv expression based on the given operands.
func NewSDiv(x, y Constant) *ExprSDiv {
	e := &ExprSDiv{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSDiv) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSDiv) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSDiv) Ident() string {
	// 'sdiv' Exactopt '(' X=TypeConst ',' Y=TypeConst ')'
	buf := &strings.Builder{}
	buf.WriteString("sdiv")
	if e.Exact {
		buf.WriteString(" exact")
	}
	fmt.Fprintf(buf, " (%s, %s)", e.X, e.Y)
	return buf.String()
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSDiv) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ fdiv ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFDiv is an LLVM IR fdiv expression.
type ExprFDiv struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
}

func (e *ExprFDiv) IsValue() {
}

// NewFDiv returns a new fdiv expression based on the given operands.
func NewFDiv(x, y Constant) *ExprFDiv {
	e := &ExprFDiv{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFDiv) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFDiv) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFDiv) Ident() string {
	// 'fdiv' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("fdiv (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFDiv) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ urem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprURem is an LLVM IR urem expression.
type ExprURem struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
}

func (e *ExprURem) IsValue() {
}

// NewURem returns a new urem expression based on the given operands.
func NewURem(x, y Constant) *ExprURem {
	e := &ExprURem{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprURem) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprURem) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprURem) Ident() string {
	// 'urem' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("urem (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprURem) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ srem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprSRem is an LLVM IR srem expression.
type ExprSRem struct {
	// Operands.
	X, Y Constant // integer scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
}

func (e *ExprSRem) IsValue() {
}

// NewSRem returns a new srem expression based on the given operands.
func NewSRem(x, y Constant) *ExprSRem {
	e := &ExprSRem{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprSRem) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprSRem) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprSRem) Ident() string {
	// 'srem' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("srem (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprSRem) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}

// ~~~ [ frem ] ~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~~

// ExprFRem is an LLVM IR frem expression.
type ExprFRem struct {
	// Operands.
	X, Y Constant // floating-point scalar or vector constants

	// extra.

	// Type of result produced by the constant expression.
	Typ types.Type
}

func (e *ExprFRem) IsValue() {
}

// NewFRem returns a new frem expression based on the given operands.
func NewFRem(x, y Constant) *ExprFRem {
	e := &ExprFRem{X: x, Y: y}
	// Compute type.
	e.Type()
	return e
}

// String returns the LLVM syntax representation of the constant expression as a
// type-value pair.
func (e *ExprFRem) String() string {
	return fmt.Sprintf("%s %s", e.Type(), e.Ident())
}

// Type returns the type of the constant expression.
func (e *ExprFRem) Type() types.Type {
	// Cache type if not present.
	if e.Typ == nil {
		e.Typ = e.X.Type()
	}
	return e.Typ
}

// Ident returns the identifier associated with the constant expression.
func (e *ExprFRem) Ident() string {
	// 'frem' '(' X=TypeConst ',' Y=TypeConst ')'
	return fmt.Sprintf("frem (%s, %s)", e.X, e.Y)
}

// Simplify returns an equivalent (and potentially simplified) constant to the
// constant expression.
func (e *ExprFRem) Simplify() Constant {
	//panic("not yet implemented")
	// TODO: implement
	return e
}
