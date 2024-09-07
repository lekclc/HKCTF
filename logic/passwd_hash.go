package logic

import (
	"crypto/sha256"
	cfg "ctf/config"
	"encoding/hex"
)

func Passwd_hash(passwd string) string {
	hash := sha256.New()
	hash.Write([]byte(passwd + cfg.Get("key.hash").(string)))
	return hex.EncodeToString(hash.Sum(nil))
}
