package main

/*
	Project: p2pgo
	Author : Brandon Vessel
	Source code : github.com/brandonvessel/p2pgo

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

const prvPEM = `
-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----`

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
