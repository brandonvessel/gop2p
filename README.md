# P2PGo

gop2p is a Golang library for enabling dead-simple UDP-based peer-to-peer communication between software clients. The objective is to be able to remotely control a variety of nodes by issuing commands that are RSA-signed. The peers maintain the public key to the private key that only the client has.

It's my first major go project, so it may not be pretty, but it does run really fast.

## Installation

The client folder contains all the code necessary for running a client that issues commands.

The peer folder contains all the code necessary for running a node.

After generating public and private RSA keys, put them in the cryptotext.go files for the client and the peer.

## Usage

After putting the keys in the correct location, run the client with:
```
./client <IP ADDRESS OF PEER HERE>
```
Run the initial peer like this:
```
./peer <SELF IP> 0
```
And run any new peers like this:
```
./peer <SELF IP> 1 <ANY PEER IP THAT IS CURRENTLY RUNNING>
```

## Peer Discovery
Peers will automatically share information about peers currently in the network, so the bootstrap peer does not have to be the same for every subsequent peer. The peers are themselves each bootstrap peers.

## Commands
The peers do not currently do anything with the commands sent to them. They only print out the raw byte slice of the data. I was only interested in the p2p communication implemented in Golang, so there are no default commands. However, implementing some commands for the peers to do requires only modifying the "executeCommand" function in the commands.go file. The buffer contains only the command array, so do with it what you want. I have some code commented out that would execute the command as a shell (but this may have dire consequences for security should you type a malicious command. You have been warned).

## Contributing
Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.

## License
Feel free to include the code in any projects you have. I only ask that you attribute this project and link back to the github page. You also don't have to use the project in its entirety. You can use any snippets you want. Just attribute it through a comment somewhere in your source code and do not claim to be the original author.

[MIT](https://choosealicense.com/licenses/mit/)
