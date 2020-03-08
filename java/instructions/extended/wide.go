package extended

import (
	"../base"
	"../loads"
	"../math"
	"../../rtda"
)

// 通过额外的字节继承局部变量表索引
type WIDE struct {
	// 被改变的指令
	modifiedInstruction base.Instruction
}

func (self *WIDE) FetchOperands(reader *base.BytecodeReader){
	// 读取一字节的操作码
	opcode := reader.ReadUint8()
	// 创建子指令实例
	switch opcode {
		case 0x15:
			// 读取子指令操作数
			inst := &loads.ILOAD{}
			inst.Index = uint(reader.ReadInt16())
			self.modifiedInstruction = inst
		case 0x84:
			inst := &math.IINC{}
			inst.Index = uint(reader.ReadUint16())
			inst.Const = int32(reader.ReadInt16())
			self.modifiedInstruction = inst
	}
}

func (self *WIDE)Execute (frame *rtda.Frame){
	self.modifiedInstruction.Execute(frame)
}


