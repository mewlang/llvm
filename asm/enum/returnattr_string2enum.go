// Code generated by "string2enum -linecomment -type ReturnAttr ../../ir/enum"; DO NOT EDIT.

package enum

import "fmt"
import "github.com/llir/llvm/ir/enum"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the string2enum command to generate them again.
	var x [1]struct{}
	_ = x[enum.ReturnAttrInReg-0]
	_ = x[enum.ReturnAttrNoAlias-1]
	_ = x[enum.ReturnAttrNonNull-2]
	_ = x[enum.ReturnAttrSignExt-3]
	_ = x[enum.ReturnAttrZeroExt-4]
}

const _ReturnAttr_name = "inregnoaliasnonnullsignextzeroext"

var _ReturnAttr_index = [...]uint8{0, 5, 12, 19, 26, 33}

// ReturnAttrFromString returns the ReturnAttr enum corresponding to s.
func ReturnAttrFromString(s string) enum.ReturnAttr {
	if len(s) == 0 {
		return 0
	}
	for i := range _ReturnAttr_index[:len(_ReturnAttr_index)-1] {
		if s == _ReturnAttr_name[_ReturnAttr_index[i]:_ReturnAttr_index[i+1]] {
			return enum.ReturnAttr(i)
		}
	}
	panic(fmt.Errorf("unable to locate ReturnAttr enum corresponding to %q", s))
}
