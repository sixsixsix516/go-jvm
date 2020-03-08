package math

import (
	"../base"
	"../../rtda"
)

// 给常量池中的int变量增加常量值
type IINC struct {
	// 局部变量表索引
	Index uint
	// 常量值
	Const int32
}

// 从字节码读取操作数
func (self *IINC) FetchOperands(reader *base.BytecodeReader){
	self.Index = uint(reader.ReadUint8())
	self.Const = uint(reader.ReadInt8())
}

// 从局部变量表中读取变量给他加上值
func (self *IINC) Execute(frame *rtda.Frame){
	localVars := frame.LocalVars()
	val := localVars.GetInt(self.Index)
	val += self.Const
	localVars.setInt(self.Index,val)
}



