package rtda

// 帧
type Frame struct {
	lower *Frame
	// 局部变量表指针
	localVars LocalVars
	// 操作数栈指针
	operandStack *OperandStack
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxLocals),
	}
}

func ( frame *Frame) OperandStack() *OperandStack {
	return frame.operandStack
}
