package ast

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/llir/llvm/internal/enc"
)

// A Function represents an LLVM IR function definition or external function
// declaration. The body of a function definition consists of a set of basic
// blocks, interconnected by control flow instructions.
type Function struct {
	// Function name.
	Name string
	// Function signature.
	Sig *FuncType
	// Basic blocks of the function; or nil if defined externally.
	Blocks []*BasicBlock
	// Metadata attached to the function.
	Metadata []*AttachedMD
}

// GetName returns the name of the value.
func (f *Function) GetName() string {
	return f.Name
}

// SetName sets the name of the value.
func (f *Function) SetName(name string) {
	f.Name = name
}

// isValue ensures that only values can be assigned to the ast.Value interface.
func (*Function) isValue() {}

// isConstant ensures that only constants can be assigned to the ast.Constant
// interface.
func (*Function) isConstant() {}

// AssignIDs assigns unique local IDs to unnamed basic blocks and local
// variables of the function.
func (f *Function) AssignIDs() {
	id := 0
	setName := func(n NamedValue) {
		name := n.GetName()
		switch {
		case isUnnamed(name):
			n.SetName(strconv.Itoa(id))
			id++
		case isID(name):
			want := strconv.Itoa(id)
			if name != want {
				panic(fmt.Errorf("invalid local ID in function %s; expected %s, got %s", enc.Global(f.Name), enc.Local(want), enc.Local(name)))
			}
			id++
		}
	}
	for _, param := range f.Sig.Params {
		// Assign local IDs to unnamed parameters of function definitions.
		if len(f.Blocks) > 0 {
			setName(param)
		}
	}
	for _, block := range f.Blocks {
		// Assign local IDs to unnamed basic blocks.
		setName(block)
		for _, inst := range block.Insts {
			n, ok := inst.(NamedValue)
			if !ok {
				continue
			}
			if inst, ok := inst.(*InstCall); ok {
				if _, ok := inst.Type.(*VoidType); ok {
					continue
				}
				if sig, ok := inst.Type.(*FuncType); ok {
					if _, ok := sig.Ret.(*VoidType); ok {
						continue
					}
				}
			}
			// Assign local IDs to unnamed local variables.
			setName(n)
		}
	}
}

// isUnnamed reports whether the given identifier is unnamed.
func isUnnamed(name string) bool {
	return len(name) == 0
}

// isID reports whether the given identifier is an ID (e.g. "%42").
func isID(name string) bool {
	for _, r := range name {
		if strings.IndexRune("0123456789", r) == -1 {
			return false
		}
	}
	return len(name) > 0
}
