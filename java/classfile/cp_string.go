package classfile

type ConstantStringInfo struct {
	cp ConstantPool
	// 存放的是常量池中的索引
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}

// 从常量池中查找指定字符擦
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}

