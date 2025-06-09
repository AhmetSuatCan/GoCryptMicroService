package ciphers

import (
	"encoding/hex"
	"golang.org/x/crypto/blake2b"
)

func EncryptBlake2bV1(data, salt string) string {
	h := blake2b.Sum512([]byte(data + salt))
	return hex.EncodeToString(h[:])
}

func EncryptBlake2bV2(data, salt string) string {
	h := blake2b.Sum512([]byte(ReverseString(data) + salt))
	return hex.EncodeToString(h[:])
}

func EncryptBlake2bV3(data, salt string) string {
	h := blake2b.Sum512([]byte(data + ReverseString(salt)))
	return hex.EncodeToString(h[:])
}

func EncryptBlake2bV4(data, salt string) string {
	h := blake2b.Sum512([]byte(ReverseString(data) + ReverseString(salt)))
	return hex.EncodeToString(h[:])
}

func EncryptBlake2bV5(data, salt string) string {
	combined := data + salt
	h := blake2b.Sum512([]byte(ReverseString(combined)))
	return hex.EncodeToString(h[:])
}
