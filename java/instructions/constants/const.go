package constants

import "../../rtda"
import "../base"

// 将null引用推入操作数栈顶
type ACONST_NULL struct{ base.NoOperandsInstruction }
type DCONST_0 struct{ base.NoOperandsInstruction }
type DCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_0 struct{ base.NoOperandsInstruction }
type FCONST_1 struct{ base.NoOperandsInstruction }
type FCONST_2 struct{ base.NoOperandsInstruction }
type ICONST_M1 struct{ base.NoOperandsInstruction }
type ICONST_0 struct{ base.NoOperandsInstruction }
type ICONST_1 struct{ base.NoOperandsInstruction }
type ICONST_2 struct{ base.NoOperandsInstruction }
type ICONST_3 struct{ base.NoOperandsInstruction }
type ICONST_4 struct{ base.NoOperandsInstruction }
type ICONST_5 struct{ base.NoOperandsInstruction }
type LCONST_0 struct{ base.NoOperandsInstruction }
type LCONST_1 struct{ base.NoOperandsInstruction }

func (self *ACONST_NULL) Execute(frame *rtda.Frame) {
	frame.OperabdStack().PushRef(nil)
}

func (self *DCONST_0) Execute(frame *rtda.Frame) {
	frame.OperabdStack().PushDouble(0.0)
}

func (self *ICONST_M1) Execute(frame *rtda.Frame) {
	frame.OperabdStack().PushInt(-1)
}
