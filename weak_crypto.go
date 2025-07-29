package main

import (
	"crypto/sha256"
	"golang.org/x/crypto/pbkdf2"
)

func example(password []byte) {
	pbkdf2.Key(password, []byte("fixedSalt"), 4096, 32, sha256.New) // Noncompliant
}
