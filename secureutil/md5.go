package secureutil

import (
	"crypto/md5"
	"encoding/hex"
	"io/ioutil"
)

func GetMd5Str(inputStr string) string {
	h := md5.New()
	h.Write([]byte(inputStr))
	cipherStr := h.Sum(nil)
	return hex.EncodeToString(cipherStr)
}

func GetFileMd5Str(filePath string) (string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}

	return GetMd5Str(string(content)), nil
}
