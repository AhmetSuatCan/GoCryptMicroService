package ciphers

var AllCiphers = []func(string, string) string{

	//V1 ==> hashes normal data and salt
	//V2 ==> hashes reversed data and normal salt
	//V3 ==> hashes normal data and reversed salt
	//V4 ==> hashes reversed data and reversed salt
	//V5 ==> combines data and salt, reverse it and hashes.

	//Blake2b versions
	EncryptBlake2bV1,
	EncryptBlake2bV2,
	EncryptBlake2bV3,
	EncryptBlake2bV4,
	EncryptBlake2bV5,

	//blake2s versions
	EncryptBlake2sV1,
	EncryptBlake2sV2,
	EncryptBlake2sV3,
	EncryptBlake2sV4,
	EncryptBlake2sV5,

	//blake3 versions
	EncryptBlake3V1,
	EncryptBlake3V2,
	EncryptBlake3V3,
	EncryptBlake3V4,
	EncryptBlake3V5,

	//md5 versions
	EncryptMD5V1,
	EncryptMD5V2,
	EncryptMD5V3,
	EncryptMD5V4,
	EncryptMD5V5,

	//ripemd160 versions
	EncryptRIPEMD160V1,
	EncryptRIPEMD160V2,
	EncryptRIPEMD160V3,
	EncryptRIPEMD160V4,
	EncryptRIPEMD160V5,

	//sha224 versions
	EncryptSHA224V1,
	EncryptSHA224V2,
	EncryptSHA224V3,
	EncryptSHA224V4,
	EncryptSHA224V5,

	//sha256 versions
	EncryptSHA256V1,
	EncryptSHA256V2,
	EncryptSHA256V3,
	EncryptSHA256V4,
	EncryptSHA256V5,

	//sha384 versions
	EncryptSHA384V1,
	EncryptSHA384V2,
	EncryptSHA384V3,
	EncryptSHA384V4,
	EncryptSHA384V5,

	//sha512 versions
	EncryptSHA512V1,
	EncryptSHA512V2,
	EncryptSHA512V3,
	EncryptSHA512V4,
	EncryptSHA512V5,
}

var AllCiphersMetadata = []string{
	// Blake2b
	"Blake2b with normal data & salt",
	"Blake2b with reversed data & normal salt",
	"Blake2b with normal data & reversed salt",
	"Blake2b with reversed data & reversed salt",
	"Blake2b with reversed (data + salt)",

	// Blake2s
	"Blake2s with normal data & salt",
	"Blake2s with reversed data & normal salt",
	"Blake2s with normal data & reversed salt",
	"Blake2s with reversed data & reversed salt",
	"Blake2s with reversed (data + salt)",

	// Blake3
	"Blake3 with normal data & salt",
	"Blake3 with reversed data & normal salt",
	"Blake3 with normal data & reversed salt",
	"Blake3 with reversed data & reversed salt",
	"Blake3 with reversed (data + salt)",

	// MD5
	"MD5 with normal data & salt",
	"MD5 with reversed data & normal salt",
	"MD5 with normal data & reversed salt",
	"MD5 with reversed data & reversed salt",
	"MD5 with reversed (data + salt)",

	// RIPEMD160
	"RIPEMD160 with normal data & salt",
	"RIPEMD160 with reversed data & normal salt",
	"RIPEMD160 with normal data & reversed salt",
	"RIPEMD160 with reversed data & reversed salt",
	"RIPEMD160 with reversed (data + salt)",

	// SHA224
	"SHA224 with normal data & salt",
	"SHA224 with reversed data & normal salt",
	"SHA224 with normal data & reversed salt",
	"SHA224 with reversed data & reversed salt",
	"SHA224 with reversed (data + salt)",

	// SHA256
	"SHA256 with normal data & salt",
	"SHA256 with reversed data & normal salt",
	"SHA256 with normal data & reversed salt",
	"SHA256 with reversed data & reversed salt",
	"SHA256 with reversed (data + salt)",

	// SHA384
	"SHA384 with normal data & salt",
	"SHA384 with reversed data & normal salt",
	"SHA384 with normal data & reversed salt",
	"SHA384 with reversed data & reversed salt",
	"SHA384 with reversed (data + salt)",

	// SHA512
	"SHA512 with normal data & salt",
	"SHA512 with reversed data & normal salt",
	"SHA512 with normal data & reversed salt",
	"SHA512 with reversed data & reversed salt",
	"SHA512 with reversed (data + salt)",
}

// Utility Function for reversing.
func ReverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
