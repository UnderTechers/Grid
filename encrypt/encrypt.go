package encrypt

import (
	"crypto/sha1"
	"io"
	"log"
	"os"
)

func ShaFile(filePath string) []byte { // get the SHA-1 value of hash value of a file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	h := sha1.New()

	if _, err := io.Copy(h, f); err != nil {
		log.Fatal(err)
	}

	return h.Sum(nil)
}

func ShaFile2() []byte { // get the SHA-1 value of hash value of a file

	h := sha1.New()

	io.WriteString(h, "Submit : 2022-12-06 10:44:10")
	return h.Sum(nil)
}
