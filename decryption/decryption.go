// decryption uses hashedKey and ciphred text to obtain a plain text

package decrypting

import (
	"log"
	"strconv"
)

func Decrypt(cipheredText string, hashedKey []int) (string, []int) {
	decipheredText := ""

	for cipheredText != "" { //iterate through whole string
		for i := 0; i < len(hashedKey); i++ {
			if cipheredText == "" {
				break
			}
			//same length hashedKey and message msg=world hashedKey=hello
			hashedKeyValueString := strconv.Itoa(hashedKey[i])
			possibleMessageChunkStr := cipheredText[0:len(hashedKeyValueString)]

			possibleMessageChunkInt, err := strconv.Atoi(possibleMessageChunkStr)
			if err != nil {
				log.Fatalf(err.Error())
			}

			cipheredText = string([]byte(cipheredText)[len(hashedKeyValueString):])

			possibleMessageInt := hashedKey[i] - possibleMessageChunkInt

			decipheredText += string(rune(possibleMessageInt))

		}
	}

	return decipheredText, hashedKey
}
