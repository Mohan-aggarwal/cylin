package main

import (
	"fmt"

	"github.com/llir/llvm/ir"
	"github.com/llir/llvm/ir/constant"
	"github.com/llir/llvm/ir/types"
)

func main() {
	// Create convenience types and constants.
	i32 := types.I32
	zero := constant.NewInt(i32, 0)
	a := constant.NewInt(i32, 0x15A4E35) // multiplier of the PRNG.
	c := constant.NewInt(i32, 1)         // increment of the PRNG.

	// Create a new LLVM IR module.
	m := ir.NewModule()

	// Create an external function declaration and append it to the module.
	//
	//    int abs(int x);
	abs := m.NewFunc("abs", i32, ir.NewParam("x", i32))
	printi := m.NewFunc("printi", types.Void, ir.NewParam("x", i32))
	utime := m.NewFunc("utime", i32)
	// Create a global variable definition and append it to the module.
	//
	//    int seed = 0;
	seed := m.NewGlobalDef("seed", zero)

	// Create a function definition and append it to the module.
	//
	//    int urand(void) { ... }
	urand := m.NewFunc("urand", i32)

	// Create an unnamed entry basic block and append it to the `urand` function.
	entry := urand.NewBlock("")

	// Create instructions and append them to the entry basic block.
	tmp1 := entry.NewLoad(i32, seed)
	tmp2 := entry.NewMul(tmp1, a)
	tmp3 := entry.NewAdd(tmp2, c)
	entry.NewStore(tmp3, seed)
	tmp4 := entry.NewCall(abs, tmp3)
	entry.NewRet(tmp4)
	mainf := m.NewFunc("main", i32)
	me := mainf.NewBlock("entry")
	me.NewStore(me.NewCall(utime), seed)
	tmp5 := me.NewCall(urand)
	me.NewCall(printi, tmp5)
	me.NewRet(zero)
	// Print the LLVM IR assembly of the module.
	fmt.Println(m)
}
