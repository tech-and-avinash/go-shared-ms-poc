package booking

import (
	"crypto/rand"
	"encoding/hex"
)

func generateID() string {
	bytes := make([]byte, 4)
	if _, err := rand.Read(bytes); err != nil {
		return "00000000"
	}
	return hex.EncodeToString(bytes)
}
