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
	"fmt"
	"os"
	//"os/exec"
	"syscall"
)

var currentCommandIndex uint32 = 0

// executeCommand executes a command
func executeCommand(buf []byte) {
	// verifies and executes a command
	//executionString := string(buf[1:])
	//fmt.Println("Executing command " + string(buf[1:]))
	////exec.Command(string(buf[1:]))
	//_, err := exec.Command(executionString).Output()
	//
	//if err != nil {
	//	fmt.Printf("%s", err)
	//}
	fmt.Println(buf)
}

// commandRestart restarts the program
func commandRestart() {
	// restarts the program
	fmt.Println("Restarting")
	if err := syscall.Exec(os.Args[0], os.Args, []string{}); err != nil {
		fmt.Println(err)
	}
}

// commandExit tells the program to shutdown
func commandExit() {
	os.Exit(1)
}

// verifyCommand uses the cryptotext functions to verify the command string came from a valid source
func verifyCommand(data []byte) (bool, []byte) {
	// get hash sum
	msgHashSum := data[1:33]

	// signature
	signature := data[33:545]

	// length of command
	var lengthSlice [4]byte
	lengthSlice[0] = data[545]
	lengthSlice[1] = data[546]
	lengthSlice[2] = data[547]
	lengthSlice[3] = data[548]

	// command index (command ID. only accept command ID's longer than this)
	var commandIndex [4]byte
	commandIndex[0] = data[549]
	commandIndex[1] = data[550]
	commandIndex[2] = data[551]
	commandIndex[3] = data[552]

	// command index check
	if byteArrayUint32To(commandIndex) > currentCommandIndex {
		currentCommandIndex = byteArrayUint32To(commandIndex)
	} else {
		return false, []byte{0x0}
	}

	clength := byteArrayUint32To(lengthSlice)

	// verify command length is valid
	if clength >= 262 {
		return false, []byte{0x0}
	}
	if clength <= 0 {
		return false, []byte{0x0}
	}

	// command
	commandData := data[553 : 553+clength]
	pubfey := textToPublicKey(pubPEM)
	sigveri := verifySignature(&pubfey, msgHashSum, signature)

	if sigveri == false {
		return false, []byte{0x0}
	}

	// verify hashsum
	datatohash := data[545 : 553+clength]
	hashresult := hashSumBytes(datatohash)

	// compare hash sums
	if sliceComp(hashresult, msgHashSum) {
		return true, commandData
	}
	return false, []byte{0x0}
}
