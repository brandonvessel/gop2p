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

// standard imports
import (
	"fmt"
	"net"
)

func main() {
	// initialize p2p engine
	initP2P()

	var input string
	fmt.Scanln(&input)
}

// portGenerator returns a port based on the IP of an address
func portGenerator(addr *net.UDPAddr) int {
	return 777
}

// confirmMessage uses cryptography to ensure only signed commands from a trusted source are processed
func confirmMessage(buf []byte) {

}
