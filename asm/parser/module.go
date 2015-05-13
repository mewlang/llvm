package parser

// Characters
//    newline      = /* the Unicode code point U+000A */ .
//    unicode_char = /* an arbitrary Unicode code point except newline */ .
//
// Letters and digits
//    letter        = ( "-" | "a" … "z" | "A" … "Z" | "$" | "." | "_" ) .
//    decimal_digit = "0" … "9" .
//    hex_digit     = "0" … "9" | "A" … "F" | "a" … "f" .
//
// Identifiers
//    Global    = GlobalID | GlobalVar .
//    Local     = LocalID | LocalVar .
//    GlobalID  = "@" ID .
//    GlobalVar = "@" Var .
//    LocalVar  = "%" Var .
//    LocalID   = "%" ID .
//    ID        = int_lit .
//    Var       = letter { letter | decimal_digit } | string_lit .
//
// Integer literals
//    int_lit =  decimal_digit { decimal_digit } .
//
// String literals
//    unicode_value  = unicode_char | hex_byte_value .
//    hex_byte_value = `\` hex_digit hex_digit .
//    string_lit     =  `"` { unicode_value | newline } `"` .

import (
	"io"

	"github.com/llir/llvm/asm/token"
	"github.com/mewkiz/pkg/errutil"
)

// TODO: Complete TopLevelEntity EBNF definition.

// parseTopLevelEntity parses a top-level entity and stores it in the module.
//
// Syntax:
//    TopLevelEntity = TargetProperty | TypeDecl | FuncDecl | FuncDef .
func (p *parser) parseTopLevelEntity() error {
	tok := p.next()
	switch tok.Kind {
	case token.Error:
		return errutil.New(tok.Val)
	case token.EOF:
		// Terminate the parser at EOF.
		return io.EOF
	case token.KwTarget:
		// Target specifications (target data layout, target host).
		return p.parseTarget()
	case token.LocalID, token.LocalVar:
		// Type definitions.
		p.backup()
		return p.parseTypeDecl()
	case token.KwDeclare:
		f, err := p.parseDeclare()
		if err != nil {
			return err
		}
		_ = f
		panic("not yet implemented.")
	case token.KwDefine:
		f, err := p.parseDefine()
		if err != nil {
			return err
		}
		_ = f
		panic("not yet implemented.")
	default:
		return errutil.Newf("invalid token type %v; expected top-level entity", tok.Kind)
	}
}

// parseTarget parses a target specification and stores it in the module. A
// "target" token has already been consumed.
//
// Syntax:
//    TargetProperty = "target" ( DataLayout | TargetTriple ) .
//    DataLayout     = "datalayout" "=" string_lit .
//    TargetTriple   = "triple" "=" string_lit .
//
// Examples:
//    target datalayout = "e-m:e-i64:64-f80:128-n8:16:32:64-S128"
//    target triple = "x86_64-unknown-linux-gnu"
//
// References:
//    http://llvm.org/docs/LangRef.html#data-layout
//    http://llvm.org/docs/LangRef.html#target-triple
func (p *parser) parseTarget() error {
	property := p.next()
	switch property.Kind {
	case token.KwDatalayout, token.KwTriple:
		// valid.
	default:
		return errutil.Newf("unknown target property %q", property)
	}
	if !p.accept(token.Equal) {
		return errutil.Newf(`expected "=" after target %s, got %q token`, property, p.next())
	}
	s, err := p.expect(token.String)
	if err != nil {
		return errutil.Err(err)
	}
	switch property.Kind {
	case token.KwDatalayout:
		p.m.Layout = s
	case token.KwTriple:
		p.m.Target = s
	}
	return nil
}
