package classpath

import (
	"os"
	"strings"
)

// 系统分隔符
const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	// 寻找加载class文件
	// 参数: className class文件相对路径,路径之间用斜线分隔
	// 例: 读取java.lang.Object 则传入的参数是 java/lang/Object
	// 返回值: byte[] : 读取到的字节数据 , 最终定位到class文件的Entry , error: 错误信息
	readClass(className string) ([]byte, Entry, error)

	// 返回变量的字符串表示
	String() string
}

/**
 * 根据参数创建不同类型的Entry实例
 */
func newEntry(path string) Entry {
	// 如果包含分隔符
	if strings.Contains(path, pathListSeparator) {
		return newComposite(path)
	}
	// 如果有通配符后缀
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") ||
		strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}

	// 否则为目录情况
	return  newDirEntry(path)

}
