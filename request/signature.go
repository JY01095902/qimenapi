package request

import (
	"crypto/md5"
	"encoding/base64"
)

func sign(content, key string) string {
	signature := md5.Sum([]byte(content + key))

	return base64.StdEncoding.EncodeToString(signature[:])
}
