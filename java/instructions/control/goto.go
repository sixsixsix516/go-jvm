package control

import (
	"../base"
	"../../rtda"
)

// 无条件跳转指令

type GOTO struct {base.NoOperandsInstruction}

func (self *GOTO) Execute(frame *rtda.Frame){
	base.Branch(frame,self.Offset)
}
