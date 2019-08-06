package classpath

import (
	"os"
	"path/filepath"
)

type Classpath struct {
	bootClasspath Entry
	extClasspath Entry
	userClasspath Entry
}

// 解析classpath
func Parse(jreOption, cpOption string) *Classpath {
	var cp = &Classpath{}  // 初始化一个struct就是用大括号
	cp.parseBootAndExtClasspath(jreOption)  // 无论如何boot, ext class都会被加载
	cp.parseUserClasspath(cpOption)  // user class只会接受输入值或者是当前目录
	return cp
}  // 生成Classpath对象

func (self *Classpath) parseBootAndExtClasspath(jreOption string) {
	jreDir := getJreDir(jreOption)

	jreLibPath := filepath.Join(jreDir, "lib", "*")
	self.bootClasspath = newWildcardEntry(jreLibPath)

	jreExtPath := filepath.Join(jreDir, "lib", "ext", "*")
	self.extClasspath = newWildcardEntry(jreExtPath)
}

func exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func getJreDir(jreOption string) string {
	if jreOption != "" && exists(jreOption) {
		return jreOption
	}
	if exists("./jre") {
		return "./jre"
	}
	if jh := os.Getenv("JAVA_HOME"); jh != "" {
		return filepath.Join(jh, "jre")
	}
	panic("Can not find jre folder!")
}

func (self *Classpath) parseUserClasspath(cpOption string) {
	if cpOption == "" {
		cpOption = "."
	}
	self.userClasspath = newEntry(cpOption)
}

// 类名不应该包含.class, 因为这里会帮忙加上
func (self *Classpath) ReadClass(className string) ([]byte, Entry, error) {
	className = className + ".class"
	if data, entry, err := self.bootClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	if data, entry, err := self.extClasspath.readClass(className); err == nil {
		return data, entry, nil
	}
	return self.userClasspath.readClass(className)  // entry.readClass, 记得接口已经规定了这个方法!
}

func (self *Classpath) String() string {
	return self.userClasspath.String()
}

