package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	/**
	寻找和加载class文件
	*/
	readClass(className string) ([]byte, Entry, error)
	/**
	description
	*/
	String() string
}

/**
   查找目录分为4种情况：
	1. path字符串包含系统分隔符";", 使用CompositeEntry
	2. path字符串最后一个字符为*
	3. path字符串包含.jar或.zip结尾
	4. path字符串为一个目录
*/
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}

	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}

	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	return newDirEntry(path)
}
