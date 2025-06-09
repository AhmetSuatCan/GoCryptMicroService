package ciphers

import (
	"crypto/md5"
	"encoding/hex"
)

func EncryptMD5V1(data, salt string) string {
	sum := md5.Sum([]byte(data + salt))
	return hex.EncodeToString(sum[:])
}

func EncryptMD5V2(data, salt string) string {
	reversed := ReverseString(data)
	sum := md5.Sum([]byte(reversed + salt))
	return hex.EncodeToString(sum[:])
}

func EncryptMD5V3(data, salt string) string {
	reversed := ReverseString(salt)
	sum := md5.Sum([]byte(data + reversed))
	return hex.EncodeToString(sum[:])
}

func EncryptMD5V4(data, salt string) string {
	reversedData := ReverseString(data)
	reversedSalt := ReverseString(salt)
	sum := md5.Sum([]byte(reversedData + reversedSalt))
	return hex.EncodeToString(sum[:])
}

func EncryptMD5V5(data, salt string) string {
	combined := data + salt
	sum := md5.Sum([]byte(combined))
	return hex.EncodeToString(sum[:])
}
