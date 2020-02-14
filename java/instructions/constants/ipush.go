package constants

import "../base"
import "../../rtda"

// Push Byte
type BIPUSH struct{ val int8 }

// Push short
type SIPUSH struct{ val int16 }

func (self *BIPUSH) FetchOperands(reader *base.BytecodeReader) {
	self.val = reader.ReadInt8()
}

func (self *BIPUSH) Execute(frame *rtda.Frame){
	i := int32(self.val)
	frame.OperandStack().PushInt(i)
}
