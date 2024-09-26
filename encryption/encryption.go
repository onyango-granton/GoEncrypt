// encryption works on plain text to ciphered text using a key
// key from hashing function

package encrypting

import (
	"goEncrypt/hashing"
	"strconv"
)

func Encrypt(message, key string) (string, []int) {
	hashedKey := hashing.Hash(key)

	encryptedMessageArr := []int{}

	encryptedMessageString := ""

	for index, value := range message {
		//length of key and message same msg=world key=hello
		if index < len(key) {
			currentValue := hashedKey[index] - int(value)
			encryptedMessageArr = append(encryptedMessageArr, currentValue)
		} else {
			// handle scenarios where message is longer than key needing iteration msg=tuesday key=hello
			// handling ay
			keyIndex := index % len(key)
			currentValue := hashedKey[keyIndex] - int(value)
			encryptedMessageArr = append(encryptedMessageArr, currentValue)
		}
		//

	}

	for _, value := range encryptedMessageArr {
		encryptedMessageString += strconv.Itoa(value)
	}

	return encryptedMessageString, hashedKey
}
