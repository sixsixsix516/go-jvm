package stack

import (
	"../base"
	"../../rtda"
)

// 交换栈顶的两个变量
type SWAP struct {base.NoOperandsInstruction}

func (self *SWAP) Execute (frame *rtda.Frame){
	stack := frame.OperandStack()
	slot1 := stack.PopSlot()
	slot2 := stack.PopSlot()
	stack.PushSlot(slot1)
	stack.PushSlot(slot2)
}

