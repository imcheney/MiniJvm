package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry{

}

func newCompositeEntry(pathList string) CompositeEntry {
	var compositeEntry = []Entry{}
	for _, path := range strings.Split(pathList, pathListSeparator) {
		entry := newEntry(path)
		compositeEntry = append(compositeEntry, entry) // append一个空数组会自动变成slice
	}
	return compositeEntry
}

// 首先自己是一个Entry数组, 因此直接遍历数组的每个元素, 获取到Entry
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found" + className)
}

func (self CompositeEntry) String() string {
	strs := make([]string, len(self))  // a slice, length = len(self)
	for i, entry := range self {
		strs[i] = entry.String()
	}

	return strings.Join(strs, pathListSeparator)
}