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
	"net"
	"sync"
	"math/rand"
	"time"
)

// Peer represents a peer
type Peer struct {
	addr            net.UDPAddr
	expirationTimer uint8 // peer is dead at 0
}

// peerList is a slice of Peers
var peerList []Peer

// peerListLock is a mutex for ensuring no race conditions among go threads for accessing the peerList
var peerListLock sync.Mutex

// addPeer adds a peer to the peerList safely
func addPeer(peer Peer) {
	// iterate to find empty slot
	peerListLock.Lock()
	for i := 0; i < MaxPeers; i++ {
		// if peer IS dead
		if peerList[i].expirationTimer == 0 {
			// ensure proper port by generating again
			peer.addr.Port = listeningPort

			// ensure expiration timer is reset
			peer.expirationTimer = expirationDefault

			// put new peer in that slot and return
			peerList[i] = peer
			peerListLock.Unlock()
			return
		}
	}
	peerListLock.Unlock()
}

// getRandomPeer returns a random peer
func getRandomPeer() *Peer {
	// iterate to count number of live peers
	livepeers := 0
	peerListLock.Lock()
	for i := 0; i < MaxPeers; i++ {
		// if peer is not dead
		if peerList[i].expirationTimer != 0 {
			// increment livepeers
			livepeers++
		}
	}
	peerListLock.Unlock()

	if livepeers == 0 {
		return &Peer{expirationTimer:0}
	}

	// generate random number
	rand.Seed(time.Now().UnixNano())
	// rnum is the distance from the start of the list through all the live peers.
	// it will give us a "kind of" random peer selection
	rnum := rand.Intn(MaxPeers)+10
	peerListLock.Lock()
	for {
		for i := 0; i < MaxPeers; i++ {
			// if peer is not dead
			if peerList[i].expirationTimer != 0 {
				// decrement rnum
				rnum--

				// if last peer
				if rnum == 0 {
					// return current peer
					peerListLock.Unlock()
					return &peerList[i]
				}
			}
		}
	}
}

// initPeerList initializes the peerList by reserving memory and ensuring the list is populated with dead peers
func initPeerList() {
	// peerList
	peerList = make([]Peer, MaxPeers)

	// populate peerList with dead peers
	for i := 0; i < MaxPeers; i++ {
		peerList[i] = Peer{expirationTimer: 0}
	}
}
