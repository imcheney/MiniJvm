package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry{

}

func (entry *DirEntry) readClass(className string) ([]byte, Entry, error) {

}

func (entry *DirEntry) String() string {

}