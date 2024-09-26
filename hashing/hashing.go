//hashing is permenent message conversion
// so the keyphrase to be used in encryption is going to be hashed

package hashing

func Hash(s string) []int {
	// given functionality keyValue can range btwn 9409 to 14884
	keylist := []int{}

	for index, value := range s {
		// successfully covered all indexes
		if index+1 < len(s) {
			currentNum := int(value) * int(s[index+1])
			keylist = append(keylist, currentNum)

		} else {
			currentNum := int(value) * int(s[0])
			keylist = append(keylist, currentNum)

		}

	}


	return keylist
}
