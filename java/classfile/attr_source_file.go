package classfile

// 可选定长属性,只会出现在classFile结构中,用户指出源文件名
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}

func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}


