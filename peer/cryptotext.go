package main

/*
	Project: gop2p
	Author : Brandon Vessel
	Source code : github.com/brandonvessel/gop2p

MIT License

Copyright (c) 2020 Brandon Vessel

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

const pubPEM = `
-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEAxEc4k+Urb2jFln926yY/
qPvgspA3X/STUUI/1p+H15IrcdDZRKjdWMeMXYftCxUN0ARnaz11tr/WV36+cinn
fImpOk3DN7ecS93KkhmrRzp2Vluprx+SwCabv8bx+RAc35ZEk36BB0D/8nyPaSuW
o4AOOBInyx41frIEvnr5wU8efhiF9fhd1lgDGujy/mj+HKPt2RzLPagHPkBeSC91
C6dm8VjjL/pEJIFIXAAWsHkbLqE96bJkBQL4KcIemZ16JgbQWsQLhuBjLv502d6t
dbH6J3VxAsvSSKfWWVV8RbEkbulyQ0/jEoNj2wxRv99jBTGprMGu6PMi1cffAgGP
nslFx3gYnbo7nsjEeAVpoHzp5KFRIj/+dr2RyOlGamPHBxJIE3ziQkYtg6mN5qfA
MNDQbW9LlOxrwyYAcqVFsdklIQVJ7wCqipVfikqvDXEem6g1bHeTgEanVLKq9BwB
Bp/B6k1Th+72MCI2s5pdMyMXtBwf+DEMm0FS7CGnkdE/xQ217LeBOnQy1Bbr/3nD
C0yScetBIxR9ue9opM9OhVPPJiwviIWhzK3P2XrL3icEuDcArmskO9w5w0o3lzbh
60qOAGQYX3K67Rv7Xrj02EKCHs/4VsoTKGPgIhRoiQ3denaOcr9SaFsPRFmS3yXk
BnHNhdE5U4bBGs+MBJft+EsCAwEAAQ==
-----END PUBLIC KEY-----`

// textToPublicKey converts a PEM encoded public key string into a RSA public key struct
func textToPublicKey(publicKeyString string) rsa.PublicKey {
	block, _ := pem.Decode([]byte(publicKeyString))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	// the key should be in a format like "BEGIN PUBLIC KEY"
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)

	if err != nil {
		fmt.Println(err)
		fmt.Println(err)
	}

	return *(pub.(*rsa.PublicKey))
}

// textToPrivateKey converts a PEM encoded public key string into a RSA private key struct
func textToPrivateKey(privateKeyString string) rsa.PrivateKey {
	block, _ := pem.Decode([]byte(privateKeyString))
	if block == nil {
		panic("failed to parse PEM block containing the public key")
	}

	// the key should be in a format like "BEGIN RSA PRIVATE KEY"
	pvr, err := x509.ParsePKCS1PrivateKey(block.Bytes)

	if err != nil {
		fmt.Println(err)
		fmt.Println(err)
	}

	return *pvr
}

// signBytes returns the msgHashSum and signature for signing a slice of bytes
func signBytes(data *[]byte, privateKey *rsa.PrivateKey) ([]byte, []byte) {
	// Before signing, we need to hash our message
	// The hash is what we actually sign
	msgHash := sha256.New()
	_, err := msgHash.Write(*data)
	if err != nil {
		panic(err)
	}
	msgHashSum := msgHash.Sum(nil)

	// In order to generate the signature, we provide a random number generator,
	// our private key, the hashing algorithm that we used, and the hash sum
	// of our message
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, msgHashSum, nil)
	if err != nil {
		panic(err)
	}

	return msgHashSum, signature
}

// verifySignature returns a boolean (true/false) for whether or not the signature and hash sum for a piece of data match
func verifySignature(publicKey *rsa.PublicKey, msgHashSum []byte, signature []byte) bool {
	// To verify the signature, we provide the public key, the hashing algorithm
	// the hash sum of our message and the signature we generated previously
	// there is an optional "options" parameter which can omit for now
	err := rsa.VerifyPSS(publicKey, crypto.SHA256, msgHashSum, signature, nil)
	if err != nil {
		fmt.Println("could not verify signature: ", err)
		return false
	}
	// If we don't get any error from the `VerifyPSS` method, that means our
	// signature is valid
	fmt.Println("signature verified")
	return true
}

// hashSumBytes returns the hash sum for a slice of bytes
func hashSumBytes(buf []byte) []byte {
	msgHash := sha256.New()
	_, err := msgHash.Write(buf)
	if err != nil {
		panic(err)
	}
	return msgHash.Sum(nil)
}

// sliceComp compares two slices to each other using length and value
func sliceComp(buf1 []byte, buf2 []byte) bool {
	// len check
	if len(buf1) != len(buf2) {
		return false
	}

	// value check
	for i := 0; i < len(buf1); i++ {
		if buf1[i] != buf2[i] {
			return false
		}
	}

	return true
}
