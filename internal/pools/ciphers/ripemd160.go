package ciphers
import (
    "golang.org/x/crypto/ripemd160"
    "encoding/hex"
)

//These functions uses ripemd160 algorithm to encrypt the data.

func EncryptRIPEMD160V1(data, salt string) string {
    h := ripemd160.New()
    h.Write([]byte(data + salt))
    return hex.EncodeToString(h.Sum(nil))
}

func EncryptRIPEMD160V2(data, salt string) string {
    reversed := ReverseString(data)
    h := ripemd160.New()
    h.Write([]byte(reversed + salt))
    return hex.EncodeToString(h.Sum(nil))
}

func EncryptRIPEMD160V3(data, salt string) string {
    reversed := ReverseString(salt)
    h := ripemd160.New()
    h.Write([]byte(data + reversed))
    return hex.EncodeToString(h.Sum(nil))
}

func EncryptRIPEMD160V4(data, salt string) string {
    reversedData := ReverseString(data)
    reversedSalt := ReverseString(salt)
    h := ripemd160.New()
    h.Write([]byte(reversedData + reversedSalt))
    return hex.EncodeToString(h.Sum(nil))
}

func EncryptRIPEMD160V5(data, salt string) string {
    combined := data + salt
    h := ripemd160.New()
    h.Write([]byte(combined))
    return hex.EncodeToString(h.Sum(nil))
}
