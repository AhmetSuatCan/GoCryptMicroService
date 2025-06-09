# 🔐 goCaaS – Go Cryptography as a Service

**goCaaS** is a lightweight, high-performance Go-based encryption API designed to demonstrate secure, salted hashing using a wide range of modern and legacy hash functions.

This project is a Go-based reimplementation of my friend’s  project: [djangoCryptoMicroservice](https://github.com/erdemserhat/djangoCryptoMicroservice)

---

## 🚀 Features

- Supports multiple hash algorithms:
- `SHA-2`, `SHA-3`, `Blake2b`, `Blake2s`, `Blake3`, `MD5`, `RIPEMD160`
- Five transformation strategies for each algorithm:
- Reversed input, reversed salt, combined-reverse, etc.
- Random salt and cipher selection per request.
- Provides a minimal REST-API endpoint for encryption.



## 🛠️ Project Structure

```
goCaaS/
├── cmd/                    # App entry point
├── internal/
│   ├── encryption/         # HTTP handlers and logging
│   ├── pools/
│   │   ├── ciphers/        # Cipher function registry + metadata
│   │   └── salts/          # Salt pool definitions
│   └── models.go           # EncryptionData struct
├── go.mod
├── main.go                 # Server setup + resource logging
└── README.md
```


## 📩 Example Request

POST /encrypt HTTP/1.1
Content-Type: application/json
```
{
  "apiKey": "sample-api-key",
  "userUUID": "123e4567-e89b-12d3-a456-426614174000",
  "sensitiveData": "mySecret123"
}
```


•	TODO: Some refactoring is needed in the ciphers package.
