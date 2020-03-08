package rtda

// 帧
type Frame struct {
	lower *Frame
	// 局部变量表指针
	localVars LocalVars
	// 操作数栈指针
	operandStack *OperandStack

	// 以下两个字段是为了实现跳转指令
	thread *Thread
	nextPC int
}

func NewFrame(thread *Thread, maxLocals, maxStack uint) *Frame {
	return &Frame{
		thread: thread,
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxLocals),
	}
}

func ( frame *Frame) OperandStack() *OperandStack {
	return frame.operandStack
}
