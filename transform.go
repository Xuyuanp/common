package common

import (
	"sync"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var bufpool = sync.Pool{
	New: func() interface{} {
		buf := make([]byte, 1024)
		return &buf
	},
}

// Transform GBK string to UTF-8 string.
func Transform(src string) (string, error) {
	buf := bufpool.Get().(*[]byte)
	defer bufpool.Put(buf)

	decoder := simplifiedchinese.GBK.NewDecoder()
	n, _, err := decoder.Transform(*buf, []byte(src), true)
	if err != nil {
		return "", err
	}
	return string((*buf)[:n]), nil
}
