package rtda

type Thread struct {
	// 程序计数器
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	// 可以容纳1024帧
	return &Thread{stack: newStack(1024)}
}

func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}

func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}

/**
 * 返回当前帧
 */
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
