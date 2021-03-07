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
	"bytes"
	"net"
)

var (
	// list of packet flags
	packetCommand    byte = 0x1
	packetPing       byte = 0x2
	packetPong       byte = 0x3
	packetAnnounce   byte = 0x4
	packetPeerUpdate byte = 0x5

	// maxBufferSize is the max size of the receiving buffer.
	// This value should not be changed
	maxBufferSize int = 812

	// myProgramName is the name of the program
	myProgramName string
	// myAddress is the name
	myAddress string
)

// this is just a note to easily figure out what UDPAddr contains without using an IDE
/*
type UDPAddr struct {
    IP   IP
    Port int
    Zone string // IPv6 scoped addressing zone; added in Go 1.1
}
*/

// UDPData represents a raw message
type UDPData struct {
	addr net.UDPAddr
	buf  []byte // this must be the same as maxBufferSize
}

/// Sending functions
// sendData sends a byte slice to an address
func sendData(addr *net.UDPAddr, data []byte) {
	Conn, _ := net.DialUDP("udp", nil, addr)
	defer Conn.Close()
	Conn.Write(data)
	Conn.Close()
}

// sendPing sends a ping request to a peer
func sendPing(addr *net.UDPAddr) {
	data := []byte{packetPing}
	Conn, _ := net.DialUDP("udp", nil, addr)
	defer Conn.Close()
	Conn.Write(data)
	Conn.Close()
}

// sendPong sends a pong packet to a peer
func sendPong(addr *net.UDPAddr) {
	data := []byte{packetPong}
	Conn, _ := net.DialUDP("udp", nil, addr)
	defer Conn.Close()
	Conn.Write(data)
	Conn.Close()
}

// sendAnnounce sends an announce packet to a peer
func sendAnnounce(addr *net.UDPAddr) {
	data := []byte{packetAnnounce}
	Conn, _ := net.DialUDP("udp", nil, addr)
	defer Conn.Close()
	Conn.Write(data)
	Conn.Close()
}

// sendPeerUpdate sends peer information to a peer
func sendPeerUpdate(addr *net.UDPAddr, peer Peer) {
	// use bytes to combine update byte with peer as byte slce
	data := bytes.Join([][]byte{[]byte{packetPeerUpdate}, ipToBytes(peer.addr.IP)}, []byte(""))
	Conn, _ := net.DialUDP("udp", nil, addr)
	defer Conn.Close()
	Conn.Write(data)
	Conn.Close()
}

// relayCommand relays a command to all live peers
func relayCommand(buf []byte) {
	// iterate through peerList checking for duplicate
	peerListLock.Lock()
	for i := 0; i < MaxPeers; i++ {
		// if peer not dead
		if peerList[i].expirationTimer != 0 {
			// send command
			sendData(&peerList[i].addr, buf)
		}
	}
	peerListLock.Unlock()
}

/// Processing functions
// processPing send a pong to the addr that sent the ping
func processPing(addr *net.UDPAddr) {
	// update addr.Port to match the port for that IP
	addr.Port = listeningPort

	// send pong to the addr that sent the ping
	sendPong(addr)
}

// processPong updates the expirationTimer for a peer if it exists in the list
func processPong(addr *net.UDPAddr) {
	// address of sender
	address := *addr

	// ip of sender
	compip := address.IP

	// iterate through peerList checking for match
	peerListLock.Lock()
	for i := 0; i < MaxPeers; i++ {
		// if peer not dead
		if peerList[i].expirationTimer != 0 {
			// if ip of not-dead peer equals sender's ip:
			if peerList[i].addr.IP.Equal(compip) {
				// update expirationTimer on peer and return
				peerList[i].expirationTimer = expirationDefault
				peerListLock.Unlock()
				return
			}
		}
	}
	peerListLock.Unlock()
}

// processAnnounce adds the sender as a peer in the list if the sender is not already in the list
func processAnnounce(addr *net.UDPAddr) {
	// check if peer already in list
	// address of sender
	address := *addr

	// ip of sender
	compip := address.IP

	// iterate through peerList checking for duplicate
	peerListLock.Lock()
	for i := 0; i < MaxPeers; i++ {
		// if peer not dead
		if peerList[i].expirationTimer != 0 {
			// if ip of not-dead peer equals sender's ip:
			if peerList[i].addr.IP.Equal(compip) {
				// return function without doing anything
				peerListLock.Unlock()
				return
			}
		}
	}
	peerListLock.Unlock()

	// not having returned here means no match was found
	// add peer to list
	addPeer(Peer{addr: address, expirationTimer: expirationDefault})

	// send announce back
	address.Port = listeningPort
	sendAnnounce(&address)
}

// processPeerUpdate processes a peer update by checking if the peer already exists in the list and adds it if does not exist
func processPeerUpdate(buf *[]byte) {
	// get peer info from buffer
	var peer Peer = bytesToPeer((*buf)[1:5])

	// iterate through peerList checking for duplicate
	ipcomp := peer.addr.IP
	peerListLock.Lock()
	for i := 0; i < MaxPeers; i++ {
		// if peer not dead
		if peerList[i].expirationTimer != 0 {
			// if ip of not-dead peer equals sender's ip:
			if peerList[i].addr.IP.Equal(ipcomp) {
				// return function without doing anything
				peerListLock.Unlock()
				return
			}
		}
	}
	peerListLock.Unlock()

	// not having returned here means no match was found
	// add peer to list
	addPeer(peer)
}

/// extra functions
func clearBuf(buf *[]byte) {
	for i := 0; i < len(*buf); i++ {
		(*buf)[i] = 0
	}
}
