package classpath

import "os"
import "strings"

const pathListSeparator = string(os.PathListSeparator)

// type一般用来定义struct和interface的名字. 是Go中实现面向对象的方式. 既有C typedef又有Java class定义的功能
// Entry指的是存放class的目录
type Entry interface {
	readClass(className string) ([]byte, Entry, error) // Go allows function to return multiple values. 负责寻找和加载class文件.
	String() string                                    // return a string
}

func newEntry(path string) Entry {
	// ~/folder1;~/folder2;
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") { // e.g. ~/myfolder/*
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") ||
			strings.HasSuffix(path, ".JAR") ||
			strings.HasSuffix(path, ".zip") ||
			strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
