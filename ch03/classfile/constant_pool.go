package classfile

// pool 本质上是数组
type ConstantPool []ConstantInfo

func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := []ConstantInfo // 感觉这是最方便的slice创建方法

	// i得从1开始, 0在cp中无效
	for i := 1; i < cpCount; i++ {  
		cp[i] = readConstantInfo(reader, cp)
		switch cp[i].(type) {
			case *ConstantLongInfo, *ConstantDoubleINfo:
				i++
		}
	}

	return cp
}

func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)  //???
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)

	return name, _type
}
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstatUtf8Info)
	return utf8Info.str
}
