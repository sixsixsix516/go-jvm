package comparisons

import (
	"../base"
	"../../rtda"
)


// 把栈顶的两个int变量弹出,进行比较,满足条件则跳转

type IF_ICMPEQ struct {base.NoOperandsInstruction}
type IF_ICMPNE struct {base.NoOperandsInstruction}
type IF_ICMPLT struct {base.NoOperandsInstruction}
type IF_ICMPLE struct {base.NoOperandsInstruction}
type IF_ICMPGT struct {base.NoOperandsInstruction}
type IF_ICMPGE struct {base.NoOperandsInstruction}


func (self *IF_ICMPNE) Execute (frame rtda.Frame){
	stack := frame.OperandStack()
	val2 := stack.PopInt()
	val1 := stack.PopInt()
	if val1 != val2 {
		base.Branch(frame,self.Offset)
	}
}
