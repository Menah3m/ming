package utils

import (
	"crypto/md5"
	"encoding/hex"
)

/*
   @Auth: menah3m
   @Desc: MD5 加解密
*/

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
