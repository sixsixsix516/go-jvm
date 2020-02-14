package rtda

type Stack struct {
	// 栈的容量
	maxSize uint
	// 栈的当前大小
	size uint
	// 栈顶指针
	_top *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{maxSize: maxSize}
}

func (self *Stack) push(frame *Frame) {
	if self.size >= self.maxSize {
		panic("java.lang.StackOverflowError")
	}

	if self._top != nil {
		frame.lower = self._top
	}

	self._top = frame
	self.size++
}

/**
 * 把栈顶元素弹出
 */
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top
}

func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}
