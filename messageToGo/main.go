package main

import (
	"fmt"
	decrypting "gopher-meet/messageToGo/decryption"
	encrypting "gopher-meet/messageToGo/encryption"
	"gopher-meet/messageToGo/hashing"
)

func main() {
	message := "Good Morning"
	key := "noon"

	hashedKey := hashing.Hash(key)

	encryptedMessage, _ := encrypting.Encrypt(message, key)

	decryptedMessage, _ := decrypting.Decrypt(encryptedMessage, hashedKey)

	fmt.Println("Hashed Key: ", hashedKey)
	fmt.Println("Encrypted Message: ", encryptedMessage)
	fmt.Println("Decrypted Message: ", decryptedMessage)
}
