package dummy_test

import (
	"github.com/llir/llvm/internal/dummy"
	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/value"
)

// Valutate that the relevant types satisfy the constant.Constant interface.
var (
	_ constant.Constant = &dummy.Global{}
)

// Valutate that the relevant types satisfy the ir.Instruction interface.
var (
	// Memory instructions
	_ ir.Instruction = &dummy.InstGetElementPtr{}
	// Other instructions
	_ ir.Instruction = &dummy.InstPhi{}
	_ ir.Instruction = &dummy.InstCall{}
)

// Valutate that the relevant types satisfy the ir.Terminator interface.
var (
	// Terminators
	_ ir.Terminator = &dummy.TermBr{}
	_ ir.Terminator = &dummy.TermCondBr{}
	_ ir.Terminator = &dummy.TermSwitch{}
)

// Valutate that the relevant types satisfy the value.Named interface.
var (
	_ value.Named = &dummy.Global{}
	_ value.Named = &dummy.Local{}
	// Memory instructions
	_ value.Named = &dummy.InstGetElementPtr{}
	// Other instructions
	_ value.Named = &dummy.InstPhi{}
	_ value.Named = &dummy.InstCall{}
)