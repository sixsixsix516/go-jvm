package conversions

import (
	"../base"
	"../../rtda"
)

// 类型转换
type D2F struct {base.NoOperandsInstruction}
type D2I struct {base.NoOperandsInstruction}
type D2L struct {base.NoOperandsInstruction}

func (self *D2I) Execute(frame *rtda.Frame){
	stack := frame.OperandStack()
	i :=  int32(stack.PopDouble())
	stack.PushInt(i)
}


