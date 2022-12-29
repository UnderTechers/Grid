package sha1_encode

import (
	"crypto/sha1"
	"encoding/hex"
	"io"
	"log"
	"os"
)

func ShaFile(filePath string) string { // get the SHA-1 value of hash value of a file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	h := sha1.New()

	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return hex.EncodeToString(h.Sum(nil))
}

func ShaText(content string) string { // get the SHA-1 value of hash value of a file

	h := sha1.New()

	io.WriteString(h, content)
	return hex.EncodeToString(h.Sum(nil))
}
