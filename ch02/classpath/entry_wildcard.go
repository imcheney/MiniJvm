package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry{
	var baseDir = path[:(len(path) - 1)]  // remove `*`
	compositeEntry := []Entry{}  // slice of Entry
	walkFn := func(path string, info os.FileInfo, err error) error {

		// 遍历每个path
		if err != nil {
			return err
		}
		if (info.IsDir() && path != baseDir) {
			return filepath.SkipDir
		}
		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			compositeEntry = append(compositeEntry, jarEntry)  // append到slice中
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)
	return compositeEntry
}
