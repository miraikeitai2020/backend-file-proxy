package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"time"
)

const (
	KEY_DIGIT = 32
)

var (
	secret string
)

func init() {
	base := sha256.Sum256([]byte(time.Now().String()))
	secret = hex.EncodeToString(base[:])
}

func main() {
	public, err := makePublicKey()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Secret Key:\t%s\n", secret)
	fmt.Printf("Access Key:\t%s\n", public)
}

func makePublicKey() (string, error) {
	b := make([]byte, KEY_DIGIT)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	var result string
	for _, v := range b {
		result += string(secret[int(v)%len(secret)])
	}
	return result, nil
}
