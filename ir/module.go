package ir

import (
	"bytes"
	"fmt"

	"github.com/llir/llvm/asm"
	"github.com/llir/llvm/types"
)

// A Module contains top-level type definitions, global variables, function
// definitions, external function declarations, and metadata.
//
// References:
//    http://llvm.org/docs/LangRef.html#module-structure
type Module struct {
	// Data layout of the module.
	//
	// Examples:
	//    target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
	//
	// References:
	//    http://llvm.org/docs/LangRef.html#data-layout
	Layout string
	// Target host of the module.
	//
	// Examples:
	//    x86_64-unknown-linux-gnu
	//
	// References:
	//    http://llvm.org/docs/LangRef.html#target-triple
	Target string
	// Type definitions.
	Types []*types.NamedStruct
	// Global variables.
	Globals []*GlobalDecl
	// Function definitions and external function declarations.
	Funcs []*Function
	// Metadata nodes.
	MetaNodes []*Metadata
}

// String returns a string representation of the module and its top-level
// declarations and definitions of types, global variables, functions and
// metadata notes.
func (m *Module) String() string {
	// Data layout; e.g.
	//    target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
	buf := new(bytes.Buffer)
	if len(m.Layout) > 0 {
		fmt.Fprintf(buf, "target datalayout = %q\n", m.Layout)
	}

	// Target triple; e.g.
	//    target triple = "x86_64-unknown-linux-gnu"
	if len(m.Target) > 0 {
		fmt.Fprintf(buf, "target triple = %q\n", m.Target)
	}

	// Type definitions; e.g.
	//    %foo = type {i32}
	for _, typ := range m.Types {
		fmt.Fprintln(buf, "%s = type %v\n", asm.EncLocal(typ.Name()), typ.Struct)
	}

	// Global variables; e.g.
	//    @x = global i32 42
	for _, global := range m.Globals {
		fmt.Fprintln(buf, global)
	}

	// TODO: Print functions.
	// TODO: Print named metadata.
	// TODO: Print metadata.
	panic("not yet implemented.")
}
