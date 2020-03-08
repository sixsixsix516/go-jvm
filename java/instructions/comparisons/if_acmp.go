package comparisons

import (
	"../base"
	"../../rtda"
)

// 负责把栈顶的两个引用弹出,根据引用是否相同跳转
type IF_ACMPEQ struct {base.NoOperandsInstruction}
type IF_ACMPNE struct {base.NoOperandsInstruction}


func (self *IF_ACMPEQ) Execute (frame *rtda.Frame){
	stack := frame.OperandStack()
	ref2 := stack.PopRef()
	ref1 := stack.PopRef()
	if ref1 == ref2{
		base.Branch(frame,self.Offset)
	}
}
