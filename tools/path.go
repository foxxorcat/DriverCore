package tools

import (
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"runtime"
)

func ToLinux(basePath string) string {
	return filepath.Clean(filepath.ToSlash(basePath))
}

func GetExecDir1() string {
	dir, _ := filepath.Abs(filepath.Dir(os.Args[0])) //返回绝对路径 filepath.Dir(os.Args[0])去除最后一个元素的路径
	return ToLinux(dir)
}

func GetExecDir2() string {
	file, _ := exec.LookPath(os.Args[0])
	path, _ := filepath.Abs(file)
	rst := filepath.Dir(path)
	return ToLinux(rst)
}

func GetExecDir3() string {
	_, filename, _, ok := runtime.Caller(0)
	var cwdPath string
	if ok {
		cwdPath = path.Join(path.Dir(filename), "") // the the main function file directory
	} else {
		cwdPath = "./"
	}
	return ToLinux(cwdPath)
}

func GetCurrentDir1() string {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return ToLinux(pwd)
}
