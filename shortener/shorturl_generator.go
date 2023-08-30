package shortener

import (
	"crypto/sha256"
	"fmt"
	"math/big"
	"os"

	"github.com/itchyny/base58-go"
)

// using SHA256 to has the initial inputs during the url shortening process.
func sha256Of(input string) []byte {
	algorithm := sha256.New()
	algorithm.Write([]byte(input))
	return algorithm.Sum(nil)
}

// using base58 to encode the final output
func base58Encoded(bytes []byte) string {
	encoding := base58.BitcoinEncoding
	encoded, err := encoding.Encode(bytes)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	return string(encoded)
}

func GenerateShortLink(initialLink string, usedId string) string {
	urlHashbytes := sha256Of(initialLink + usedId)
	generateNumber := new(big.Int).SetBytes(urlHashbytes).Uint64()
	finalString := base58Encoded([]byte(fmt.Sprintf("%d", generateNumber)))
	return finalString[:8]
}
