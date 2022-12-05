package safe

import (
	"crypto/rand"
	"encoding/base64"
	"math/big"
	"strings"
)

func Reverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func Unreverse(str string) (result string) {
	for _, v := range str {
		result = string(v) + result
	}
	return
}

func random(min, max int64) int64 {
	bg := big.NewInt(max - min)

	n, err := rand.Int(rand.Reader, bg)
	if err != nil {
		panic(err)
	}
	return n.Int64() + min
}

func GenerateSym(length int) string {
	letterRunes := []rune("-|/#$&")
	mixedRunes := []rune("-|/#$&")
	b := make([]rune, length)
	b[0] = letterRunes[random(0, int64(len(mixedRunes)-1))]
	for i := range b {
		if i != 0 {
			b[i] = mixedRunes[random(0, int64(len(mixedRunes)-1))]
		}
	}

	return string(b)
}

func Encode(str string) string {

	base := base64.StdEncoding.EncodeToString([]byte(str))

	base = strings.ReplaceAll(base, "=", "")

	rev := Reverse(base[:len(base)/2+5])

	var res string

	for _, sym := range rev {
		res += string(sym) + GenerateSym(1)
	}

	return res

}

func Decode(str string) string {

	d := Unreverse(str)
	un := strings.ReplaceAll(d, "-", "")
	un = strings.ReplaceAll(un, "|", "")
	un = strings.ReplaceAll(un, "/", "")
	un = strings.ReplaceAll(un, "#", "")
	un = strings.ReplaceAll(un, "$", "")
	un = strings.ReplaceAll(un, "&", "")

	return un
}
