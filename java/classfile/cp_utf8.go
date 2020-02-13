package classfile

type ConstantUtf8Info struct {
	// Java虚拟机中 字符串采用 MUTF8编码
	str string
}
func (self *ConstantUtf8Info) readInfo(reader *ClassReader){
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = string(bytes)
}
