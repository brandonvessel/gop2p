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
	"bufio"
	"bytes"
	"fmt"
	"net"
	"os"
)

func runClient(addr net.UDPAddr) {
	var userinput string
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("Enter command:")

		// Taking input from user
		userinput, _ = reader.ReadString('\n')

		// create slice from string
		dataSlice := []byte(userinput[:len(userinput)-1])

		// command length
		commandlength := uint32ToByteArray(uint32(len(dataSlice)))

		// command index
		nextCommandIndex := uint32ToByteArray(currentCommandIndex + 1)

		// more data slices
		dataSlice2 := bytes.Join([][]byte{commandlength, nextCommandIndex}, []byte(""))
		dataSlice3 := bytes.Join([][]byte{dataSlice2, dataSlice}, []byte(""))

		// sign bytes
		rss := textToPrivateKey(prvPEM)
		msgHashSum, signature := signBytes(&dataSlice3, &rss)

		// combine data together
		datatosend := bytes.Join([][]byte{[]byte{packetCommand}, msgHashSum}, []byte(""))
		datatosend2 := bytes.Join([][]byte{datatosend, signature}, []byte(""))
		datatosend3 := bytes.Join([][]byte{datatosend2, dataSlice3}, []byte(""))

		sendData(&addr, datatosend3)
		currentCommandIndex++
	}
}
