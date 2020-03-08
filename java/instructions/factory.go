package instructions

import "./base"

var (
	nop  = &NOP{}
	aconst_null = &ACONST_NULL{}
)

func NewInstruction(opcode byte)base.Instruction{
	switch opcode {
		case 0x00: return nop
	}
}

