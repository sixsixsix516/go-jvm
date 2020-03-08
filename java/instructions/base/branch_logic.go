package base

import "../../rtda"

// 跳转
func Branch(frame *rtda.Frame,offset int){
	pc := frame.Thread().PC()
	nextPC := pc + offset
	frame.SetNextPC(nextPC)
}
