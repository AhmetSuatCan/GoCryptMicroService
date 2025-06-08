package ciphers

import (
	"crypto/sha256"
	"encoding/hex"
)

func EncryptSHA224V1(data, salt string) string {
	h := sha256.New224()
	h.Write([]byte(data + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptSHA224V2(data, salt string) string {
	h := sha256.New224()
	h.Write([]byte(ReverseString(data) + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptSHA224V3(data, salt string) string {
	h := sha256.New224()
	h.Write([]byte(data + ReverseString(salt)))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptSHA224V4(data, salt string) string {
	h := sha256.New224()
	h.Write([]byte(ReverseString(data) + ReverseString(salt)))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptSHA224V5(data, salt string) string {
	combined := data + salt
	h := sha256.New224()
	h.Write([]byte(ReverseString(combined)))
	return hex.EncodeToString(h.Sum(nil))
}
