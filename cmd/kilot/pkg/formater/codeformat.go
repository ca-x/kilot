package formater

import (
	"go/format"
)

func FormatCode(code []byte) ([]byte, error) {
	return format.Source(code)
}
