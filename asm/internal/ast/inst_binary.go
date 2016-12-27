// generated by gen.go using 'go generate'; DO NOT EDIT.

// === [ Binary instructions ] =================================================
//
// References:
//    http://llvm.org/docs/LangRef.html#binary-operations

package ast

// --- [ add ] -----------------------------------------------------------------

// InstAdd represents an addition instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#add-instruction
type InstAdd struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Operands.
	X, Y Value
}

// GetName returns the name of the value.
func (inst *InstAdd) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstAdd) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstAdd) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstAdd) isInst() {}

// --- [ fadd ] ----------------------------------------------------------------

// InstFAdd represents a floating-point addition instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fadd-instruction
type InstFAdd struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Operands.
	X, Y Value
}

// GetName returns the name of the value.
func (inst *InstFAdd) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstFAdd) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstFAdd) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstFAdd) isInst() {}

// --- [ sub ] -----------------------------------------------------------------

// InstSub represents a subtraction instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#sub-instruction
type InstSub struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Operands.
	X, Y Value
}

// GetName returns the name of the value.
func (inst *InstSub) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstSub) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstSub) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstSub) isInst() {}

// --- [ fsub ] ----------------------------------------------------------------

// InstFSub represents a floating-point subtraction instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fsub-instruction
type InstFSub struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Operands.
	X, Y Value
}

// GetName returns the name of the value.
func (inst *InstFSub) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstFSub) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstFSub) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstFSub) isInst() {}

// --- [ mul ] -----------------------------------------------------------------

// InstMul represents a multiplication instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#mul-instruction
type InstMul struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Operands.
	X, Y Value
}

// GetName returns the name of the value.
func (inst *InstMul) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstMul) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstMul) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstMul) isInst() {}

// --- [ fmul ] ----------------------------------------------------------------

// InstFMul represents a floating-point multiplication instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fmul-instruction
type InstFMul struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Operands.
	X, Y Value
}

// GetName returns the name of the value.
func (inst *InstFMul) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstFMul) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstFMul) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstFMul) isInst() {}

// --- [ udiv ] ----------------------------------------------------------------

// InstUDiv represents an unsigned division instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#udiv-instruction
type InstUDiv struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Operands.
	X, Y Value
}

// GetName returns the name of the value.
func (inst *InstUDiv) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstUDiv) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstUDiv) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstUDiv) isInst() {}

// --- [ sdiv ] ----------------------------------------------------------------

// InstSDiv represents a signed division instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#sdiv-instruction
type InstSDiv struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Operands.
	X, Y Value
}

// GetName returns the name of the value.
func (inst *InstSDiv) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstSDiv) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstSDiv) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstSDiv) isInst() {}

// --- [ fdiv ] ----------------------------------------------------------------

// InstFDiv represents a floating-point division instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#fdiv-instruction
type InstFDiv struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Operands.
	X, Y Value
}

// GetName returns the name of the value.
func (inst *InstFDiv) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstFDiv) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstFDiv) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstFDiv) isInst() {}

// --- [ urem ] ----------------------------------------------------------------

// InstURem represents an unsigned remainder instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#urem-instruction
type InstURem struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Operands.
	X, Y Value
}

// GetName returns the name of the value.
func (inst *InstURem) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstURem) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstURem) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstURem) isInst() {}

// --- [ srem ] ----------------------------------------------------------------

// InstSRem represents a signed remainder instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#srem-instruction
type InstSRem struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Operands.
	X, Y Value
}

// GetName returns the name of the value.
func (inst *InstSRem) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstSRem) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstSRem) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstSRem) isInst() {}

// --- [ frem ] ----------------------------------------------------------------

// InstFRem represents a floating-point remainder instruction.
//
// References:
//    http://llvm.org/docs/LangRef.html#frem-instruction
type InstFRem struct {
	// Name of the local variable associated with the instruction.
	Name string
	// Operands.
	X, Y Value
}

// GetName returns the name of the value.
func (inst *InstFRem) GetName() string {
	return inst.Name
}

// SetName sets the name of the value.
func (inst *InstFRem) SetName(name string) {
	inst.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*InstFRem) isValue() {}

// isInst ensures that only instructions can be assigned to the ast.Instruction
// interface.
func (*InstFRem) isInst() {}