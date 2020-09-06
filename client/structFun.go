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
	"encoding/binary"
	"net"
)

// SliceMock is a mock-slice
type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

// ipToBytes converts an IP to a slice of bytes using the net.IP.To4() function
func ipToBytes(ip net.IP) []byte {
	return ip.To4()
}

// bytesToPeer converts a slice of bytes representing IP data into a Peer struct and returns it
func bytesToPeer(data []byte) Peer {
	var peeraddr net.UDPAddr
	peeraddr.IP = net.IPv4((data)[0], (data)[1], (data)[2], (data)[3])
	return Peer{addr: peeraddr}
}

// uint32ToByteArray converts a uint32 to a [4]byte slice
func uint32ToByteArray(v uint32) []byte {
	var b [4]byte
	b[0] = byte(v >> 24)
	b[1] = byte(v >> 16)
	b[2] = byte(v >> 8)
	b[3] = byte(v)
	return b[:]
}

// byteArrayUint32To converts a [4]byte slice to a uint32
func byteArrayUint32To(arr [4]byte) uint32 {
	return binary.BigEndian.Uint32(arr[0:4])
}
