package ir

// An Instruction belongs to one of the following groups:
//
//    * terminator instructions
//    * binary instructions
//    * bitwise binary instructions
//    * memory instructions
//    * other instructions
//
// References:
//    http://llvm.org/docs/LangRef.html#instruction-reference
type Instruction interface {
	// isInst ensures that only instructions can be assigned to the Instruction
	// interface.
	isInst()
}

// =============================================================================
// Terminator Instructions
//
//    ref: http://llvm.org/docs/LangRef.html#terminators
// =============================================================================

// A ReturnInst returns control flow (and optionally a value) from a function
// back to the caller.
//
// Syntax:
//    ret <Type> <Val>
//    ret void
//
// Reference:
//    http://llvm.org/docs/LangRef.html#i-ret
type ReturnInst struct {
	// Return type.
	Type Type
	// Return value; or nil in case of a void return.
	Val Value
}

// A CondBranchInst transfers control flow to one of two basic blocks in the
// current function based on a boolean branching condition.
//
// Syntax:
//    br i1 <Cond>, label <TargetTrue>, label <TargetFalse>
//
// References:
//    http://llvm.org/docs/LangRef.html#i-br
type CondBranchInst struct {
	// Boolean branching condition.
	Cond Value
	// Target branch when the condition evaluates to true.
	True *BasicBlock
	// Target branch when the condition evaluates to false.
	False *BasicBlock
}

// A BranchInst transfers control flow to a basic block in the current function.
//
// Syntax:
//    br label <Target>
//
// References:
//    http://llvm.org/docs/LangRef.html#i-br
type BranchInst struct {
	// Target branch.
	Target *BasicBlock
}

// A SwitchInst transfers control flow to one of several basic blocks in the
// current function.
//
// Syntax:
//    switch <IntType> <Val>, label <TargetDefault> [ <IntType> <Const>, label <Target> ... ]
//
// References:
//    http://llvm.org/docs/LangRef.html#i-switch
type SwitchInst struct {
	// TODO(u): Restrict Type to IntType.
	// Comparasion type.
	Type Type
	// Comparasion value.
	Val Value
	// Default target.
	Default *BasicBlock
	// Switch cases.
	Cases []struct {
		// Case value.
		Val Constant
		// Case target.
		Target *BasicBlock
	}
}

// TODO(u): Add terminator instructions.

// =============================================================================
// Binary Operations
//
//    ref: http://llvm.org/docs/LangRef.html#binaryops
// =============================================================================

// TODO(u): Add binary operations.

// =============================================================================
// Bitwise Binary Operations
//
//    ref: http://llvm.org/docs/LangRef.html#bitwiseops
// =============================================================================

// TODO(u): Add bitwise binary operations.

// =============================================================================
// Memory Access and Addressing Operations
//
//    ref: http://llvm.org/docs/LangRef.html#memoryops
// =============================================================================

// TODO(u): Add memory access and addressing operations.

// =============================================================================
// Other Operations
//
//    ref: http://llvm.org/docs/LangRef.html#otherops
// =============================================================================

// TODO(u): Add other operations.

// isInst ensures that only instructions can be assigned to the Instruction
// interface.
func (ReturnInst) isInst()     {}
func (CondBranchInst) isInst() {}
func (BranchInst) isInst()     {}
func (SwitchInst) isInst()     {}