package ciphers
import (
    "encoding/hex"
    "golang.org/x/crypto/blake2s"
)

//These functions uses blake2b algorithm to encrypt the data.

func EncryptBlake2sV1(data, salt string) string {
    h := blake2s.Sum256([]byte(data + salt))
    return hex.EncodeToString(h[:])
}

func EncryptBlake2sV2(data, salt string) string {
    reversedData := ReverseString(data)
    h := blake2s.Sum256([]byte(reversedData + salt))
    return hex.EncodeToString(h[:])
}

func EncryptBlake2sV3(data, salt string) string {
    reversedSalt := ReverseString(salt)
    h := blake2s.Sum256([]byte(data + reversedSalt))
    return hex.EncodeToString(h[:])
}

func EncryptBlake2sV4(data, salt string) string {
    reversedData := ReverseString(data)
    reversedSalt := ReverseString(salt)
    h := blake2s.Sum256([]byte(reversedData + reversedSalt))
    return hex.EncodeToString(h[:])
}

func EncryptBlake2sV5(data, salt string) string {
    combined := data + salt
    reversed := ReverseString(combined)
    h := blake2s.Sum256([]byte(reversed))
    return hex.EncodeToString(h[:])
}
