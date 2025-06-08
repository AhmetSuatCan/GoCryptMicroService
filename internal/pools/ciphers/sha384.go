package ciphers

import (
	"crypto/sha512"
	"encoding/hex"
)


func EncryptSHA384V1(data, salt string) string {
	h := sha512.New384()
	h.Write([]byte(data + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptSHA384V2(data, salt string) string {
	h := sha512.New384()
	h.Write([]byte(ReverseString(data) + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptSHA384V3(data, salt string) string {
	h := sha512.New384()
	h.Write([]byte(data + ReverseString(salt)))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptSHA384V4(data, salt string) string {
	h := sha512.New384()
	h.Write([]byte(ReverseString(data) + ReverseString(salt)))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptSHA384V5(data, salt string) string {
	combined := data + salt
	h := sha512.New384()
	h.Write([]byte(ReverseString(combined)))
	return hex.EncodeToString(h.Sum(nil))
}
