package ciphers

import (
	"crypto/sha512"
	"encoding/hex"
)

func EncryptSHA512V1(data, salt string) string {
	h := sha512.New()
	h.Write([]byte(data + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptSHA512V2(data, salt string) string {
	h := sha512.New()
	h.Write([]byte(ReverseString(data) + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptSHA512V3(data, salt string) string {
	h := sha512.New()
	h.Write([]byte(data + ReverseString(salt)))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptSHA512V4(data, salt string) string {
	h := sha512.New()
	h.Write([]byte(ReverseString(data) + ReverseString(salt)))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptSHA512V5(data, salt string) string {
	combined := data + salt
	h := sha512.New()
	h.Write([]byte(ReverseString(combined)))
	return hex.EncodeToString(h.Sum(nil))
}
