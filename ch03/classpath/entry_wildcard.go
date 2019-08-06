package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

func newWildcardEntry(path string) CompositeEntry {
	var baseDir = path[:(len(path) - 1)]  // remove `*`
	var compositeEntry = []Entry{}  // slice of Entry
	var walkFn = func(path string, info os.FileInfo, err error) error {

		// 遍历每个目录下的文件, 如果是jar文件就解读, 并且最后合并到compositeEntry中
		if err != nil {
			return err
		}
		if (info.IsDir() && path != baseDir) { // 目录的话跳过
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
