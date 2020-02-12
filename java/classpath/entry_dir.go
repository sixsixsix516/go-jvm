package classpath

import "io/ioutil"
import "path/filepath"

type DirEntry struct {
	// 目录的绝对路径
	absDir string
}


/**
 * 构造函数
 */
func newDirEntry(path string) *DirEntry {
	// 将路径转换为绝对路径
	absDir, err := filepath.Abs(path)
	if err != nil {
		// 如果发生错误则终止程序
		panic(err)
	}
	// 创建实例并返回
	return &DirEntry{absDir}
}

func (self *DirEntry)readClass(className string)([]byte, Entry,error){
	fileName := filepath.Join(self.absDir,className)
	data,err := ioutil.ReadFile(fileName)
	return data,self,err
}

func (self *DirEntry) String() string {
	return self.absDir
}
