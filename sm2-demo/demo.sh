#!/usr/bin/env bash
set -e

# 1. Generate SM2 key pair (private key with public key inside)
openssl ecparam -name sm2 -genkey -noout -out sm2.pem

# 2. Plaintext to encrypt
printf 'Hello, SM2 demo!\n' > plain.txt

# 3. Encrypt with the private key (public part embedded)
openssl sm2 -encrypt -inkey sm2.pem -in plain.txt -out enc.bin

# 4. Decrypt back to plaintext
openssl sm2 -decrypt -inkey sm2.pem -in enc.bin > dec.txt

# Show result
cat dec.txt
