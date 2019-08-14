
package classfile
import "encoding/binary"

// ClassReader仅仅是对byte数组的一个包装, 
// 进行包装的原因是我们想attach 一些方法到byte数组上. 那么最好的办法就是利用stuct, 形成一个类
type ClassReader struct {
	data []byte
}

// read one token
// u1
func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:]
	return val
}

// u2
func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:]
	return val
}

// u4
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.uint32(self.data)
	self.data = self.data[4:]
	return val
}

func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.uint64(self.data)
	self.data = self.data[8:]
	return val
}

// read array
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16()
	s := make([]uint16, n)  // 创建一个长度为n的uint16元素类型的slice
	for i:= range s {  // 遍历0~15
		s[i] = self.readUint16()
	}
	return s
}

// read bytes
func (self *ClassReader) readBytes(n uint32) []byute {
	bytes := self.data[:n]  // 左闭右开区间
	self.data = self.data[n:]
	return bytes
}



