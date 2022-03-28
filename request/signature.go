package request

import (
	"crypto/md5"
	"encoding/base64"
	"net/url"
)

func sign(content, key string) string {
	signature := md5.Sum([]byte(content + key))

	return base64.StdEncoding.EncodeToString(signature[:])
}

func ValidateSignature(input, key string) bool {
	decodeInput, err := url.QueryUnescape(input)
	if err != nil {
		return false
	}

	values, err := url.ParseQuery(decodeInput)
	if err != nil {
		return false
	}

	signature := sign(values.Get("logistics_interface"), key)

	return signature == values.Get("data_digest")
}
