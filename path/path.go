package path

import (
	"path"
	"path/filepath"
	"runtime"
)

func GetRootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	dir := filepath.Dir(d)

	return dir
}

func GetRootPath() string {
	dir := GetRootDir()

	return filepath.Base(dir)
}
