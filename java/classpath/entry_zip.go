package classpath

import (
	"archive/zip"
	"errors"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	// 存放zip或jar包绝对路径
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) String() string {
	return self.absPath
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	// 打开zip包
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	// 关闭文件
	defer r.Close()

	// 便利zip包的内容
	for _, f := range r.File {
		// 从zip包内找到指定类
		if f.Name == className {
			// 找到后打开文件
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}

			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return nil, nil, err
			}
			return data, self, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
