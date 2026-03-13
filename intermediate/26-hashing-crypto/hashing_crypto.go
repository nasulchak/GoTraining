package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
)

func main() {
	password := "password123"

	// hash256 := sha256.Sum256([]byte(password))
	// hash512 := sha512.Sum512([]byte(password))

	// fmt.Printf("SHA-256: %x\n", hash256)
	// fmt.Printf("SHA-512: %x\n", hash512)
	salt, err := generateSalt()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Original Salt:", salt)
	fmt.Printf("Hex Salt: %x\n", salt)
	// Hash the password with salt
	signUpHash := hashPassword(password, salt)

	// Store the salt and passwod in database, just printing as of now
	saltStr := base64.StdEncoding.EncodeToString(salt)
	// fmt.Println("Salt:", saltStr)
	// fmt.Println("Hash:", signUpHash)

	// Verify
	// Retrieve the saltStr and decode it
	decodeSalt, err := base64.StdEncoding.DecodeString(saltStr)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	loginHash := hashPassword("password124", decodeSalt)

	// Compare the stored signUpHash with loginHash
	if signUpHash == loginHash {
		fmt.Println("Password is correct. You are logged in.")
	} else {
		fmt.Println("Login failed. Please check user credential.")
	}
}

// Function to generate salt
func generateSalt() ([]byte, error) {
	salt := make([]byte, 16)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return nil, err
	}

	return salt, nil
}

// Function to hash password
func hashPassword(password string, salt []byte) string {
	saltedPassword := append(salt, []byte(password)...)
	hash := sha256.Sum256(saltedPassword)
	return base64.StdEncoding.EncodeToString(hash[:])
}
