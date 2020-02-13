package classfile

import "encoding/binary"

type ClassReader struct {
	data []byte
}

/**
 * uint8(8比特无符号整数) -> byte 来读取u1(1字节无符号整数)
 */
func (self *ClassReader) readerUint8() uint8 {
	val := self.data[0]
	// 剩下的还给他
	self.data = self.data[1:]
	return val
}

/**
 * uint16(16比特无符号整数) 读取u2(2字节无符号整数)
 */
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val;
}

/**
 * 读取u4类型数据
 */
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:]
	return val
}

/**
 * 读取 uint64数据
 */
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:]
	return val
}

/**
 * 读取16表
 */
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}

/**
 * 读取指定数量的字节
 */
func (self *ClassReader) readBytes(n uint32) [] byte {
	bytes := self.data[:n]
	self.data = self.data[n:]
	return bytes
}
