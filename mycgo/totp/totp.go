package totp

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"fmt"
	"strings"
	"time"
)

func Totp(secret string) string {
	// secret := "AICRSHHFUHB2XGSHLO6QSNDMJYPIUKQC"
	key, _ := base32.StdEncoding.DecodeString(strings.ToUpper(secret))
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(time.Now().Unix()/30))
	hash := hmac.New(sha1.New, key)
	hash.Write(bs)
	h := hash.Sum(nil)
	o := (h[19] & 15)
	var header uint32
	r := bytes.NewReader(h[o : o+4])
	binary.Read(r, binary.BigEndian, &header)
	h12 := (int(header) & 0x7fffffff) % 1000000

	return fmt.Sprintf(fmt.Sprintf("%%0%dd", 6), h12)
}
