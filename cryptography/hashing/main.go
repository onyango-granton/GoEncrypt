package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
)

func main() {
	data := "Hello world"

	fmt.Printf("MD5 %x  \n", md5.Sum([]byte(data)))
	fmt.Printf("SHA1 %x  \n", sha1.Sum([]byte(data)))
	fmt.Printf("SHA256 %x  \n", sha256.Sum256([]byte(data)))
	fmt.Printf("SHA512 %x  \n", sha512.Sum512([]byte(data)))
}
