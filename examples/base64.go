package examples

import (
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"math/rand"
)

func intToBase64String(value uint64) string {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, value)
	return base64.RawURLEncoding.EncodeToString(b)
}

func DemoBase64() {
	for i := 0; i < 10; i++ {
		value := rand.Uint64()
		fmt.Println(value, intToBase64String(value))
	}
}
