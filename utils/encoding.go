package utils

import (
	"math/big"
	"strings"
)

const base62Chars = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

func EncodeBase62(input string) string {
	bytes := []byte(input)

	bigInt := new(big.Int).SetBytes(bytes)

	var result strings.Builder
	base := big.NewInt(62)
	zero := big.NewInt(0)
	mod := new(big.Int)

	for bigInt.Cmp(zero) > 0 {
		bigInt.DivMod(bigInt, base, mod)
		result.WriteByte(base62Chars[mod.Int64()])
	}

	return reverse(result.String())
}

func DecodeBase62(encoded string) string {
	bigInt := new(big.Int)
	for _, char := range encoded {
		bigInt.Mul(bigInt, big.NewInt(62))
		bigInt.Add(bigInt, big.NewInt(int64(strings.IndexByte(base62Chars, byte(char)))))
	}

	bytes := bigInt.Bytes()

	return string(bytes)
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
