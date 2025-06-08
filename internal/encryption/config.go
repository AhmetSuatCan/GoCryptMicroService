package encryption

import "math/rand"

type EncryptConfig struct {
    UUID string
    cipherSize int
    saltSize int
}

func (e EncryptConfig) GenerateIndexPair(cipherLimit int, saltLimit int) (int, int) {
    uuidLength := len(e.UUID)
    uuidProduct := 0

    for i, b := range []byte(e.UUID) {
        product := (i + 1) * int(b)
        uuidProduct += product
    }

    denominator := float64((uuidLength * (uuidLength + 1) / 2) + 1)
    normalizedProduct := float64(uuidProduct) / denominator

    seed := int64(normalizedProduct * 10000)
    random := rand.New(rand.NewSource(seed))

    cipherIndex := random.Intn(cipherLimit)
    saltIndex := random.Intn(saltLimit)

    return cipherIndex, saltIndex
}
