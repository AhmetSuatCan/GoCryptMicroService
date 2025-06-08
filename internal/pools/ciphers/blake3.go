package ciphers

import (
    "encoding/hex"
    "github.com/zeebo/blake3"
)

//These functions uses blake3 algorithm to encrypt the data.

func EncryptBlake3V1(data, salt string) string {
    h := blake3.Sum256([]byte(data + salt))
    return hex.EncodeToString(h[:])
}

func EncryptBlake3V2(data, salt string) string {
    reversed := ReverseString(data)
    h := blake3.Sum256([]byte(reversed + salt))
    return hex.EncodeToString(h[:])
}

func EncryptBlake3V3(data, salt string) string {
    reversed := ReverseString(salt)
    h := blake3.Sum256([]byte(data + reversed))
    return hex.EncodeToString(h[:])
}

func EncryptBlake3V4(data, salt string) string {
    reversedData := ReverseString(data)
    reversedSalt := ReverseString(salt)
    h := blake3.Sum256([]byte(reversedData + reversedSalt))
    return hex.EncodeToString(h[:])
}

func EncryptBlake3V5(data, salt string) string {
    combined := data + salt
    reversed := ReverseString(combined)
    h := blake3.Sum256([]byte(reversed))
    return hex.EncodeToString(h[:])
}
