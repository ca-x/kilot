package ioplus

import (
	"os"
	"path"
	"path/filepath"
)

func IsDirExist(dir string) bool {
	if len(dir) == 0 {
		return false
	}
	if _, err := os.Stat(dir); os.IsExist(err) {
		return true
	}
	return false
}

func GetFileDir(filePath string) string {
	return filepath.Dir(filePath)
}

func EnsureDir(dir string) error {
	if !IsDirExist(dir) {
		return CreateDir(dir)
	}
	return nil
}

func EnsureFileDir(filePath string) error {
	dir := filepath.Dir(filePath)
	if !IsDirExist(dir) {
		return CreateDir(dir)
	}
	return nil
}

func CreateDir(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

func DirWithSub(parentDir string, subDir string) string {
	sPath := path.Join(parentDir, subDir)
	if path.IsAbs(sPath) {
		return sPath
	}
	absPath, _ := filepath.Abs(sPath)
	return absPath
}

func FilePathFrom(parentDir string, subDir string, fileName string) string {
	sPath := path.Join(parentDir, subDir, fileName)
	if path.IsAbs(sPath) {
		return sPath
	}
	absPath, _ := filepath.Abs(sPath)
	return absPath
}

func GetWorkingDir() string {
	if dir, err := os.Getwd(); err != nil {
		return "."
	} else {
		return dir
	}
}
