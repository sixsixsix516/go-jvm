package classfile

type ConstantNameAndTypeInfo struct {
	// 字段名/方法 名索引
	nameIndex uint16
	// 描述符索引
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
