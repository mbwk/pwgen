package random

import (
	crypto_rand "crypto/rand"
	"encoding/binary"
	"math/rand"
)

func init() {
	var b [8]byte
	_, err := crypto_rand.Read(b[:])
	if err != nil {
		panic("Could not seed rand")
	}
	rand.Seed(int64(binary.LittleEndian.Uint64(b[:])))
}

func Shuffle(buffer []byte) {
	rand.Shuffle(len(buffer), func(i, j int) {
		buffer[i], buffer[j] = buffer[j], buffer[i]
	})
}

func RandomSelection(source string) byte {
	return source[rand.Intn(len(source))]
}
