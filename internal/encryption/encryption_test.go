package encryption

import (
    "testing"
)

func TestGenerateIndexPair(t *testing.T) {
    config := EncryptConfig{
        UUID: "123e4567-e89b-12d3-a456-426614174000",
        cipherSize: 100,
        saltSize: 100,
    }

    cipherLimit := config.cipherSize
    saltLimit := config.saltSize

    cipherIndex, saltIndex := config.GenerateIndexPair(cipherLimit, saltLimit)

    if cipherIndex < 0 || cipherIndex >= cipherLimit {
        t.Errorf("cipherIndex out of range: got %d", cipherIndex)
    }

    if saltIndex < 0 || saltIndex >= saltLimit {
        t.Errorf("saltIndex out of range: got %d", saltIndex)
    }

    c2, s2 := config.GenerateIndexPair(cipherLimit, saltLimit)
    if cipherIndex != c2 || saltIndex != s2 {
        t.Errorf("Expected deterministic output, got (%d, %d) and (%d, %d)",
            cipherIndex, saltIndex, c2, s2)
    }
}
