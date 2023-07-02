package force

import (
	"crypto/sha256"
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Decrypt(x string) string {
	n := time.Duration(rand.Intn(100)+1) * time.Millisecond
	defer time.Sleep(n)

	for i := 0; i < 10; i++ {
		s := strconv.Itoa(i)
		h := hashing([]byte(s)[0])
		if h == x {
			return s
		}
	}
	return ""
}

func hashing(b byte) string {
	h := sha256.New()
	h.Write([]byte{b})

	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
