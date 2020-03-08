package control

import (
	"../base"
	"../../rtda"
	"github.com/astaxie/beego/utils"
)

type TABLE_SWITCH struct {
	// 默认情况下执行跳转所需要的字节码漂移量
	defaultOffset int32
	// 记录case的取值范围
	low int32
	// 记录case的取值范围
	high int32
	// 索引表
	jumpOffsets []int32
}

func (self *TABLE_SWITCH) FetchOperands(reader *base.BytecodeReader){
	reader.SkipPadding()
	self.defaultOffset = reader.ReadInt32()
	self.low = reader.ReadInt32()
	self.high = reader.ReadInt32()

	jumpOffsetCount := self.high - self.low +1
	self.jumpOffsets = reader.ReadInt32s(jumpOffsetCount)
}


func (self *TABLE_SWITCH) Execute(frame *rtda.Frame){
	index := frame.OperandStack().PopInt()
	// 定义偏移量
	var offset int
	if index >= self.low && index <= self.high{
		offset = int(self.jumpOffsets[index - self.low])
	}else{
		offset = int(self.defaultOffset)
	}

	base.Branch(frame,offset)
}


