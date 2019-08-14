package classfile
import "fmt"

type ClassFile struct {
	// magic		uint32
	minorVersion	uint16
	majorVersion	uint16
	constantPool	ConstantPool
	accessFlags		uint16
	thisClass		uint16
	superClass		uint16
	interfaces		[]uint16
	fields			[]*MemberInfo
	methods			[]*MemberInfo
	attributes		[]AttributesInfo   // 为什么这里要field之外再来一个attributes?
}

// 所有首字母大写的类型, 结构体, 字段, 变量, 函数, 方法都是公开的.
// 首字母小写的是私有的
// parse classfile to classdata bytes
// 输入byte数组, 由一个空的cr对象接收byte数组, 然后解析后返回一个cf
// 所以cf是最终结果
func Parse(classData []byte) (cf *ClassFile, err error) {
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)  // ????
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()  // 执行一个匿名函数

	cr := &ClassReader{classData}  // define a pointer
	cf = &ClassFile{}
	cf.read(cr)
	return
}

// classfile.read(specific_reader)
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
		case 45:
			return
		case 46, 47, 48, 49, 50, 51, 52:
			if self.minorVersion == 0 {
				return
			}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// getter
func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
}

func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
}

func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
}

func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
}


func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	return self.constantPool.getClassName(self.superClass)
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))  // slice
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
