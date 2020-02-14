package constants

import "../base"
import "../../rtda"

// nop 什么也不做指令
type NOP struct{
	base.NoOperandsInstruction
}

func (self *NOP) Execute(frame *rtda.Frame){

}


