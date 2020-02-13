package classfile

// 变长属性,只存在于method_info中,用于存放字节码等方法的相关信息
type CodeAttribute struct {
	cp ConstantPool
	// 操作数栈的最大深度
	maxStack uint16
	// 局部变量表大小
	maxLocals uint16
	// 字节码
	code []byte
	// 异常处理表
	exceptionTable []*ExceptionTableEntity
	// 属性表
	attributes []AttributeInfo
}

type ExceptionTableEntity struct {
	startPc uint16
	endPc uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader)[]*ExceptionTableEntity{
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntity,exceptionTableLength)
	for i := range exceptionTable{
		exceptionTable[i] = &ExceptionTableEntity{
			startPc:reader.readUint16(),
			endPc:reader.readUint16(),
			handlerPc:reader.readUint16(),
			catchType:reader.readUint16(),
		}
	}
	return exceptionTable
}
