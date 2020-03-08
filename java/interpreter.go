package main

import (
	"./classfile"
	"./rtda"
	"fmt"
	"./instructions"
	"./instructions/base"
)

func interpret(methodInfo *classfile.MemberInfo){
	codeArr := methodInfo.CodeAttribute()
	maxLocals := codeArr.MaxLocals()
	maxStack := codeArr.MaxStack()
	bytecode := codeArr.Code()

	thread := rtda.NewThread()
	frame := thread.NewFrame(maxLocals,maxStack)
	thread.PushFrame(frame)

	defer carchErr(frame)
	loop(thread,bytecode)
}


func carchErr(frame *rtda.Frame){
	if r := recover();r!= nil{
		fmt.Printf("LocalVars:%v \n",frame.LocalVars())
		fmt.Printf("OperandStack:%v \n",frame.OperandStack())
		panic(r)
	}
}

/*
	1. 计算pc
	2. 解码指令
	3. 执行指令
*/
func loop(thread *rtda.Thread,bytecode []byte){
		frame := thread.PopFrame()
		reader := &base.BytecodeReader{}
		for {
			pc := frame.NextPC()
			thread.SetPc(pc)

			// decode
			reader.Reset(bytecode,pc)
			opcode := reader.ReadUint8()
			inst := instructions.NewInstruction(opcode)
			inst.FetchOperands(reader)
			frame.SetNextPC(reader.PC())

			// execute
			fmt.Printf("pc : %2d inst: %T %v \n",pc,inst,inst)
			inst.Execute(frame)
		}
}

