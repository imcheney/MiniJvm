package classpath

import (
	"io/ioutil"
	"path/filepath"
)

// 类定义
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

func (entry *DirEntry) readClass(className string) ([]byte, Entry, error) {
	var fileName = filepath.Join(entry.absDir, className)
	var data, err = ioutil.ReadFile(fileName)
	return data, entry, err
}

// 似乎每个类都要实现String接口, 用来把打印object
func (entry *DirEntry) String() string {
	return entry.absDir
}
