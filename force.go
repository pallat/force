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

var seeding = "310978029187439287498574938709479208749857498723895"

type seed []byte

func (ss *seed) rid(b byte) {
	s := *ss
	defer func() { *ss = s }()
	for i := range []byte(s) {
		if b == s[i] {
			if i == 0 {
				if len(s) == 0 {
					s = make([]byte, 0)
					return
				}
				s = s[1:]
				return
			}
			if x := len(s) - 1; i == x {
				s = s[:x]
				return
			}
			s = append(s[:i], s[i+1:]...)
		}
	}
}

func Validate(s string) bool {
	t := seed(seeding)
	for _, b := range []byte(s) {
		t.rid(b)
	}
	return len(t) == 0
}

func hashing(b byte) string {
	h := sha256.New()
	h.Write([]byte{b})

	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
