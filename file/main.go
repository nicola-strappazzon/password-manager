package file

import (
	"log"
	"os"
)

func ReadInString(in string) string {
	return string(ReadInBytes(in))
}

func ReadInBytes(in string) []byte {
	data, err := os.ReadFile(in)
	if err != nil {
		log.Fatal("read file:", err)
	}

	return data
}
