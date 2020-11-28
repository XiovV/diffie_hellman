package diffie_hellman

import (
	"math"
)

type DiffieHellman struct {
	P int
	G int
}

func (dh DiffieHellman) GetPublicKey(privateKey int) int {
	return dh.power(dh.G, privateKey)
}

func (dh DiffieHellman) GetSharedSecret(otherPublic int, localPrivate int) int {
	return dh.power(otherPublic, localPrivate)
}

func (dh DiffieHellman) power(a int, b int) int {
	pow := int(math.Pow(float64(a), float64(b)))
	return pow % dh.P
}
