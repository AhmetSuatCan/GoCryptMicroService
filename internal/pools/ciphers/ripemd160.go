package ciphers

import (
	"encoding/hex"
	"golang.org/x/crypto/ripemd160"
)

func EncryptRIPEMD160V1(data, salt string) string {
	h := ripemd160.New()
	h.Write([]byte(data + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptRIPEMD160V2(data, salt string) string {
	h := ripemd160.New()
	h.Write([]byte(ReverseString(data) + salt))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptRIPEMD160V3(data, salt string) string {
	h := ripemd160.New()
	h.Write([]byte(data + ReverseString(salt)))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptRIPEMD160V4(data, salt string) string {
	h := ripemd160.New()
	h.Write([]byte(ReverseString(data) + ReverseString(salt)))
	return hex.EncodeToString(h.Sum(nil))
}

func EncryptRIPEMD160V5(data, salt string) string {
	combined := data + salt
	h := ripemd160.New()
	h.Write([]byte(combined))
	return hex.EncodeToString(h.Sum(nil))
}
