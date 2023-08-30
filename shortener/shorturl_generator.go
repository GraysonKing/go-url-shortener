package shortener

import (
	"crypto/sha256"
	"fmt"
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
