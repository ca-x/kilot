package ioplus

import (
	"os"
)

const regularPerm = 0o666

func SaveFile(fileName string, data []byte) error {
	return os.WriteFile(fileName, data, regularPerm)
}
