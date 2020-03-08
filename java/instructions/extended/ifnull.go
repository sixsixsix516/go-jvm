package extended

import (
	"../base"
	"../../rtda"
)

// 根据引用是否是null进行跳转
type IFNULL struct {base.BranchInstruction}
type IFNOTNULL struct {base.BranchInstruction}

func (self *IFNULL)Execute(frame *rtda.Frame) {
	ref := frame.OperandStack().PopRef()
	if ref == nil{
		base.Branch(frame,self.Offset)
	}
}
