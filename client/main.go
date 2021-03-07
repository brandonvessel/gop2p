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
	"flag"
	"fmt"
	"net"
	"os"
)

// listening port for client
var listeningPort int

func main() {
	// arguments
	// listening port number integer
	flag.IntVar(&listeningPort, "p", 0, "Listening port")

	// IP string of peer
	var parentIP string

	flag.StringVar(&parentIP, "b", "", "IP of peer node")

	// parse flags
	flag.Parse()

	// check if peer IP exists
	if parentIP == "" {
		// print message
		fmt.Println("Peer IP must be specified")
		
		// print usage
		flag.Usage()
		os.Exit(1)
	}

	// check if port exists
	if listeningPort == 0 {
		// print message
		fmt.Println("Port must be specified")
		
		// print usage
		flag.Usage()
		os.Exit(1)
	}

	runClient(net.UDPAddr{IP: net.ParseIP(parentIP), Port: listeningPort})
}

// confirmMessage uses cryptography to ensure only signed commands from a trusted source are processed
func confirmMessage(buf []byte) {

}
