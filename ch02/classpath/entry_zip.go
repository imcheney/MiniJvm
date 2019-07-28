package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	absPath string
}

// instance create func
// return a ptr of ZipEntry
func newZipEntry(path string) *ZipEntry{
	var absPath, err = filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}  // get_addr运算符是最后的, 因此可以理解成 &(ZipEntry{absPath})
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error){
	var resource, err = zip.OneReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer resource.Close()
	for _, f := range resource.File {
		if f.Name == className {  // 文件名如果是类名, 那就把文件内容读出来返回
			content, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer content.Close()  // defer 类似 java finally. 会在return的时候必然调用一次
			data, err := ioutil.ReadAll(content)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}