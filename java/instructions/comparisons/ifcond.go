package comparisons

import (
	"../base"
	"../../rtda"
)

// 把操作数栈顶的int变量弹出,然后与0进行比较,满足条件则跳转

// if x == 0
type IFEQ struct{ base.NoOperandsInstruction }

// if x != 0
type IFNE struct{ base.NoOperandsInstruction }

// if x < 0
type IFLT struct{ base.NoOperandsInstruction }

// if x <= 0
type IFLE struct{ base.NoOperandsInstruction }

// if x > 0
type IFGT struct{ base.NoOperandsInstruction }

// if x >= 0
type IFGE struct{ base.NoOperandsInstruction }

func (self *IFEQ) Execute(frame rtda.Frame) {
	val := frame.OperandStack().PopInt()
	if val == 0 {
		base.Branch(frame,self.Offset)
	}
}


