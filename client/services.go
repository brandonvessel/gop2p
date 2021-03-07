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
	"net"
	"os"
	"runtime"
	"time"
)

func initP2P() {
	// initialize peerList
	initPeerList()
	// parse arguments
	initParseArguments()
	// initialize services
	go serviceInit()
}

func serviceInit() {
	bufchan := make(chan UDPData, 10)
	// start goroutines
	// pinger routine
	go servicePinger()

	// peer update routine
	go serviceUpdater()

	// listener thread
	go serviceListener(bufchan)

	// handler threads
	for i := 0; i < runtime.NumCPU()+1; i++ {
		go serviceHandler(bufchan)
		go serviceHandler(bufchan)
		go serviceHandler(bufchan)
		go serviceHandler(bufchan)
	}
}

func initParseArguments() {
	// arguments
	// 0 name of file
	var myProgramName string = os.Args[0]
	fmt.Println("Now running " + myProgramName + " on " + runtime.GOOS + "-" + runtime.GOARCH)

	// 1 ip address of current machine
	myAddress := net.UDPAddr{IP: net.ParseIP(os.Args[1]), Port: listeningPort}
	fmt.Println("My address: " + net.IP.String(myAddress.IP))

	// 2 whether or not a parent exists

	// 3 ip address of parent
	if os.Args[2] == "1" {
		// add peer to peerList
		peerListLock.Lock()
		peerList[0] = Peer{expirationTimer: expirationDefault}
		peerList[0].addr.IP = net.ParseIP(os.Args[3])
		peerList[0].addr.Port = listeningPort

		// send announce to peer
		sendAnnounce(&peerList[0].addr)

		peerListLock.Unlock()
	}
}

// servicePinger is a goroutine for keep-alive pinging every peer on an interval
func servicePinger() {
	for {
		// iterate through list sending pings
		for i := 0; i < MaxPeers; i++ {
			peerListLock.Lock()
			// if peer is not dead
			if peerList[i].expirationTimer != 0 {
				// send ping packet to peer
				sendPing(&peerList[i].addr)

				// lower expirationTimer for peer
				peerList[i].expirationTimer--

				// free peerList before sleep
				peerListLock.Unlock()

				// sleep between sending ping requests to not overload network interface
				time.Sleep(pingWait * time.Millisecond)
			} else {
				// must unlock if not used because otherwise lock would happen twice and deadlock program
				peerListLock.Unlock()
			}
		}

		// print peerlist
		for i := 0; i < MaxPeers; i++ {
			peerListLock.Lock()
			// if peer is not dead
			if peerList[i].expirationTimer != 0 {
				// lower expirationTimer for peer
				fmt.Println(peerList[i])
			}
			peerListLock.Unlock()
		}

		// sleep between rounds
		time.Sleep(pingDelay * time.Millisecond)
	}
}

// serviceUpdater is a goroutine for sending peer information through peer updates to all known peers
func serviceUpdater() {
	for {
		// iterate through list sending pings
		for i := 0; i < MaxPeers; i++ {
			peerListLock.Lock()
			// if peer is not dead
			if peerList[i].expirationTimer != 0 {
				// free peerList before sleep
				peerListLock.Unlock()
				
				// send ping packet to peer
				sendPeerUpdate(&peerList[i].addr, *(getRandomPeer()))

				// sleep between sending ping requests to not overload network interface
				time.Sleep(pingWait * time.Millisecond)
			} else {
				// must unlock if not used because otherwise lock would happen twice and deadlock program
				peerListLock.Unlock()
			}
		}

		// sleep between rounds
		time.Sleep(pingDelay * time.Millisecond)
	}
}

// serviceListener is a goroutine for listening on the UDP connection and sending the received data to the channel as fast as possible
func serviceListener(bufchan chan UDPData) {
	// setup udp listening port for messages
	ServerConn, err := net.ListenUDP("udp", &net.UDPAddr{IP: []byte{0, 0, 0, 0}, Port: listeningPort, Zone: ""})
	defer fmt.Println(err)

	// create buffer for message recieving
	buf := make([]byte, maxBufferSize)

	for {
		// read udp data into buffer
		_, addr, _ := ServerConn.ReadFromUDP(buf)

		// send data into channel for handlers
		bufchan <- UDPData{addr: *addr, buf: buf}
	}
}

// serviceHandler is a goroutine for retrieving data from the bufchan channel and processing it.
// Because there can be multiple instances of the serviceHandler, the receiving buffer can process data as fast as it can load it into the channel.
// This increases overall throughput.
func serviceHandler(bufchan chan UDPData) {
	// data holds the current data from the channel
	var data UDPData

	for {
		// read data from the channel
		data = <-bufchan

		// figure out what the message is for
		switch data.buf[0] {
		case packetCommand:
			// execute a command
			// verify command
			result, commandSlice := verifyCommand(data.buf)

			// execute command
			if result == true {
				// relay command
				relayCommand(data.buf)
				// execute command
				executeCommand(commandSlice)
			}

		case packetPing:
			// ping request
			// process
			fmt.Println("recv ping: ", data.addr)
			processPing(&data.addr)

		case packetPong:
			// pong response
			// update expirationTimer for peer
			fmt.Println("recv pong: ", data.addr)
			processPong(&data.addr)

		case packetAnnounce:
			// new bot announce
			fmt.Println("recv Announce: ", data.addr)
			processAnnounce(&data.addr)

		case packetPeerUpdate:
			// new peer update
			fmt.Println("recv Peer update: ", data.addr)
			processPeerUpdate(&data.buf)

		default:
			// put default action here
			fmt.Println("Invalid packet: ", data.addr)
		}
	}
}
