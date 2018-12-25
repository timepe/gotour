package classpath

import (
	"os"
	"path/filepath"
	"strings"
)

/**
遍历path目录，将所有的
*/
func newWildcardEntry(path string) CompositeEntry {
	baseDir := path[:len(path)-1]
	ce := []Entry{}
	walkFn := func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		/**
		  跳过目录
		*/
		if info.IsDir() && path != baseDir {
			return filepath.SkipDir
		}

		if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") {
			jarEntry := newZipEntry(path)
			ce = append(ce, jarEntry)
		}
		return nil
	}
	filepath.Walk(baseDir, walkFn)
	return ce
}
