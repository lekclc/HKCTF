package logic

import (
	"crypto/md5"
	"crypto/sha256"
	cfg "ctf/config"
	"encoding/hex"
)

func Passwd_hash(passwd string) string {
	hash := sha256.New()
	key_ := Rmd5(cfg.Get("key.hash").(string))
	hash.Write([]byte(passwd + key_))
	return hex.EncodeToString(hash.Sum(nil))
}
func Rmd5(str string) string {
	bytes := []byte(str)
	hash := md5.New()
	hash.Write(bytes)
	return hex.EncodeToString(hash.Sum(nil))
}
func Md5(str string) string {
	str = Passwd_hash(str)
	bytes := []byte(str)
	hash := md5.New()
	hash.Write(bytes)
	return hex.EncodeToString(hash.Sum(nil))
}
