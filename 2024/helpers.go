package _024

import (
	"log"
	"os"
	"strconv"
)

func ReadInput(filename string) string {
	content, err := os.ReadFile(filename)
	if err != nil {
		log.Fatalf("error opening file: %v\n", err)
	}
	return string(content)
}

func MustParseInt(s string) int {
	return Must(strconv.Atoi(s))
}

func Must[T any](t T, err error) T {
	if err != nil {
		log.Fatalf("must function failed: %v", err)
	}
	return t
}
