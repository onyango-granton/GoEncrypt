# GoEncrypt

GoEncrypt is a simple encryption project that demonstrates basic text encryption and decryption using a key-based hashing mechanism.

## Features

- Key hashing
- Text encryption using the hashed key
- Text decryption using the hashed key

## Installation

To use this project, make sure you have Go installed on your system. Then, clone the repository:

```
git clone https://github.com/onyango-granton/GoEncrypt.git
cd GoEncrypt
```

## Usage

[Provide instructions on how to use the encryption and decryption functions. For example:]

```go
key := "noon"
plaintext := "Good Morning"

Hashed Key:  [12210 12321 12210 12100]
Encrypted Message:  121391221012099120001217812244120991198612100122161210011997
Decrypted Message:  Good Morning
```

## How It Works

1. The provided key is hashed to create a secure encryption key.
2. The hashed key is used to encrypt the plain text, producing ciphertext.
3. To decrypt, the same hashed key is used on the ciphertext to recover the original plain text.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Author

Granton Onyango

## Contact

[Provide contact information or links to your professional profiles if desired]

## Acknowledgments

[Optional: Add any acknowledgments or credits here]