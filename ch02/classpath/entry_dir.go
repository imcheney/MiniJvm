package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// 类定义
// 目录形式的entry是最基本最常使用的
type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	var absDir, err = filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}

// 如果要求读这个entry下的某一个类, 就走的是这个接口
func (entry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	var fileName = filepath.Join(entry.absDir, className)
	var data, err = ioutil.ReadFile(fileName)
	return data, entry, err
}

// 似乎每个类都要实现String接口, 用来把打印本entry的信息
func (entry *DirEntry) String() string {
	return entry.absDir
}
